package main

import (
	"bytes"
	"context"
	"douyin/api"
	"douyin/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/fsnotify/fsnotify"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	goQrcode "github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"image"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	// 创建临时文件夹
	_, err := os.Stat("./tmp")
	if err != nil && os.IsNotExist(err) {
		err = nil
		err = os.MkdirAll("./tmp", os.ModePerm)
	}
	if err != nil {
		panic(err)
	}
	time.Local = time.FixedZone("CST", 3600*8)
	// 加载配置
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

var running = false
var inLogin = false // 是否正在登录

func main() {
	// 监听登录状态
	for {
		run()
		time.Sleep(time.Second * time.Duration(viper.GetInt64("Interval.CheckLoginS")))
	}
}

func run() {
	uResp := api.UserTrack()
	if uResp.St != 0 || !uResp.Data.UserIsLogin {
		// 未登录
		login()
	} else if !running {
		running = true
		// 获取数据
		lResp := api.ListQuickview()
		if lResp.St != 0 {
			fmt.Println("获取商户数据失败;请重新登录,失败信息:" + lResp.Msg)
			// 获取失败 重新登录
			login()
		}else{
			if len(lResp.Data.DataResult) <= 0 {
				running = false
				// 未找到直播间
				fmt.Println("未找到直播间;请重新登录")
			}else{
				for _, r := range lResp.Data.DataResult {
					go timer(r)
				}
			}
		}
	}
}

func login() {
	if !inLogin {
		inLogin = true
		running = false // 关闭所有爬虫
		// 打开浏览器
		// chromdp依赖context上限传递参数
		ctx, _ := chromedp.NewExecAllocator(
			context.Background(),
			// 以默认配置的数组为基础，覆写headless参数
			// 当然也可以根据自己的需要进行修改，这个flag是浏览器的设置
			append(
				chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", true),
				chromedp.Flag("blink-settings", "imagesEnabled=false"),
				chromedp.UserAgent(viper.GetString("System.UserAgent")),
			)...,
		)
		// 创建新的chromedp上下文对象，超时时间的设置不分先后
		// 注意第二个返回的参数是cancel()，只是我省略了
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		ctx, _ = chromedp.NewContext(
			ctx,
			// 设置日志方法
			chromedp.WithLogf(log.Printf),
		)
		defer cancel()
		// 设置超时关闭
		go func() {
			time.Sleep(time.Minute * 2)
			fmt.Println("二维码过期")
			chromedp.Cancel(ctx)
			inLogin = false
		}()
		// 执行我们自定义的任务 - myTasks函数在第4步
		if err := chromedp.Run(ctx, myTasks()); err != nil {
			fmt.Println(err.Error())
		}
		inLogin = false
	}
}

// 自定义任务
func myTasks() chromedp.Tasks {
	return chromedp.Tasks{
		// 1. 打开登陆界面
		chromedp.Navigate("https://compass.jinritemai.com/login"),
		chromedp.Click("#my-node > main > div > div.loginWrapper--2RTfW > div.rolePannel--25WFD > div.pannelRest--21kS3 > div > div.roleCard---v7UW.seller--1k9ot > svg > rect:nth-child(1)"),
		chromedp.Click("#my-node > main > div > div.loginWrapper--2RTfW > div > div:nth-child(2) > div > div > div > div.index_oauthLogin__1vWnv > div.index_oauthLoginBody__378Xb > div:nth-child(1)"),
		getCode(),
		WaitLogin(),
	}
}


func getCode() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 1. 用于存储图片的字节切片
		html := ""
		// 2. 截图
		// 注意这里需要注明直接使用ID选择器来获取元素（chromedp.ByID）
		if err = chromedp.OuterHTML(`#root > div > div.content-container > div.auth-container > div.auth-qr-container > div.qr-container > img`,
			&html).Do(ctx); err != nil {
			return
		}
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil{
			return
		}
		str, ok := dom.Find("img").Attr("src")
		if  !ok {
			err = errors.New("二维码获取失败")
			return
		}
		str = strings.Replace(str, `data:image/png;base64,`, ``, 1)
		qB, err := base64.StdEncoding.DecodeString(str)
		// 打印二维码
		err = printQRCode(qB)
		if err != nil {
			return
		}
		return
	}
}

func WaitLogin() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		sel := `#root > div > div.headerWrapper--7HYa6 > div > div.headerTools--TX8PU > div > div`
		if err = chromedp.WaitVisible(sel).Do(ctx); err != nil {
			return
		}
		// 保存cookies
		if err = utils.SaveCookies(ctx); err != nil {
			return
		}
		return
	}
}

func printQRCode(code []byte) (err error) {
	// 1. 因为我们的字节流是图像，所以我们需要先解码字节流
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		return
	}

	// 2. 然后使用gozxing库解码图片获取二进制位图
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return
	}

	// 3. 用二进制位图解码获取gozxing的二维码对象
	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return
	}
	// 4. 用结果来获取go-qrcode对象（注意这里我用了库的别名）
	qr, err := goQrcode.New(res.String(), goQrcode.Low)
	if err != nil {
		return
	}
	// 5. 输出到标准输出流
	fmt.Println(qr.ToString(false))
	return
}


func timer(r api.ListQuickviewResult) {
	productC, _ := getProductCsv(nil, r.AnchorNickname)
	for {
		if !running {
			// 停止
			return
		}
		//sResp := api.ScreenBaseInfo(r.LiveRoomID, r.LiveAppID)
		//if sResp.St != 0 {
		//	return
		//}
		//data := sResp.Data
		//fmt.Println("=========================================================")
		//fmt.Println("获取时间:" + utils.TimeFormat("Y-m-d H:i:s", time.Now()))
		//fmt.Println(fmt.Sprintf("直播间:%s", data.Title))
		//fmt.Println(fmt.Sprintf("累计成交金额(元):%d万", data.Gmv/100))
		//fmt.Println(fmt.Sprintf("成交件数:%d    成交人数:%d    成交转化率:%s%%    千次观看成交金额:%s    成交粉丝占比:%s%%",
		//	data.PayCnt.Value, data.PayUcnt.Value, utils.KeepFloat64ToString(data.ProductClickToPayRate.Value * 100, 2),
		//	utils.KeepFloat64ToString(data.Gpm.Value, 2), utils.KeepFloat64ToString(data.PayFansRatio.Value * 100, 2)))
		//fmt.Println("=========================================================")
		n := time.Now()
		pResp := api.ScreenProductDetail(r.LiveRoomID, r.LiveAppID, "bind_time", false)
		if pResp.St != 0 {
			return
		}
		data := pResp.Data.DataResult
		// 写csv
		productC, _ = getProductCsv(productC, r.AnchorNickname)
		var columns [][]string
		for _, d := range data {
			//"名称", "价格", "商品点击率", "成交订单数", "成交金额", "成交转化率",
			columns = append(columns, []string{
				utils.TimeFormat("Y-m-d H:i:s", n), d.Title, d.CurrMinPrice, d.ProductClickInLiveRate, d.PayInLiveOrderCnt, d.PayInLiveOrderProductGmv, d.ProductClickToPayRate,
			})
		}
		fmt.Println("-----------------------"+utils.TimeFormat("Y-m-d H:i:s", n)+"---------------------------")
		fmt.Println(columns)
		err := productC.W.WriteAll(columns)
		if err != nil {
			fmt.Println("写入异常", err)
		}
		fmt.Println("---------------------------------------------------------------------")
		time.Sleep(time.Second * time.Duration(viper.GetInt64("Interval.GrabS")))
	}
}
// 获取对应csv对象
func getCsv(oldC *utils.Csv, path, name string, title []string) (*utils.Csv, error) {
	// 取出当前时间
	n := time.Now()
	if oldC == nil {
		return utils.NewCsv(path, name, n, title)
	}
	// 判断csv是否过期
	tmp := utils.TimeFormat("Y-m-d", oldC.CreateTime)
	s := utils.CreateTimeFormat("Y-m-d", tmp)
	tmp = utils.TimeFormat("Y-m-d", n)
	e := utils.CreateTimeFormat("Y-m-d", tmp)
	// 判断
	if s.Before(e) {
		// 需要新建
		oldC.Close()
		return utils.NewCsv(path, name, n, title)
	}
	return oldC, nil
}

// 获取商品写入csv
func getProductCsv(c *utils.Csv, nickName string) (*utils.Csv, error) {
	c, err := getCsv(c, "./数据/" + nickName + "/商品", "商品数据", []string{
		"抓取时间", "名称", "价格", "商品点击率", "成交订单数", "成交金额", "成交转化率",
	})
	return c, err
}