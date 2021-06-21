package main

import (
	"context"
	"douyin/api"
	"douyin/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	// 创建临时文件夹
	_, err := os.Stat(utils.FolderPath + "/tmp")
	if err != nil && os.IsNotExist(err) {
		err = nil
		err = os.MkdirAll(utils.FolderPath + "/tmp", os.ModePerm)
	}
	if err != nil {
		panic(err)
	}
	time.Local = time.FixedZone("CST", 3600*8)
	// 加载配置
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(utils.FolderPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

// 二维码获取状态
var qrcodeLatest = false
var waitScan = false // 等待扫描

// 检测登录状态
func getLoginStatus() (b bool, msg string, list []api.ListQuickviewResult) {
	b = true
	msg = "success"
	uResp := api.UserTrack()
	if uResp.St != 0 || !uResp.Data.UserIsLogin {
		b = false
		msg = "请登录"
		return
	} else {
		lResp := api.ListQuickview()
		if lResp.St != 0 {
			b = false
			msg = "获取商户数据失败;请重新登录;原因:" + lResp.Msg
			return
		}else{
			list = lResp.Data.DataResult
			return
		}
	}
}

func main() {

	go func(){
		// 循环检测登录状态
		// 监听登录状态
		for {
			//return
			// 未登录并且未获取二维码 则获取最新二维码
			b, _, _ := getLoginStatus()
			// 未获取最新状态并且也没有等待扫描 则获取最新二维码
			if !b && !waitScan && !qrcodeLatest {
				waitScan = true
				ctx, _ := chromedp.NewExecAllocator(
					context.Background(),
					// 以默认配置的数组为基础，覆写headless参数
					// 当然也可以根据自己的需要进行修改，这个flag是浏览器的设置
					append(
						chromedp.DefaultExecAllocatorOptions[:],
						chromedp.Flag("headless", false),
						chromedp.Flag("blink-settings", "imagesEnabled=false"),
						chromedp.UserAgent(viper.GetString("System.UserAgent")),
					)...,
				)
				ctx, _ = context.WithTimeout(ctx, time.Minute * 1)
				ctx, _ = chromedp.NewContext(
					ctx,
					// 设置日志方法
					chromedp.WithLogf(log.Printf),
				)
				// 设置超时关闭
				//go func() {
				//	time.Sleep(time.Minute * 2)
				//	fmt.Println("二维码过期")
				//	waitScan = false
				//	qrcodeLatest = false
				//	chromedp.Cancel(ctx)
				//	cancel()
				//}()
				// 执行我们自定义的任务 - myTasks函数在第4步
				if err := chromedp.Run(ctx, myTasks()); err != nil {
					fmt.Println(err.Error())
				}
				waitScan = false
				qrcodeLatest = false
			}
			time.Sleep(time.Second * time.Duration(viper.GetInt64("Interval.CheckLoginS")))
		}
	}()
	// 开启gin服务
	r := gin.Default()
	// 加载html文件
	r.LoadHTMLGlob(utils.FolderPath + "/templates/html/*")
	r.Static("/js", utils.FolderPath + "/templates/js")
	r.Static("/css", utils.FolderPath + "/templates/css")
	r.Static("/tmp/img", utils.FolderPath + "/tmp")
	// 初始化路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/api/get/config", func(c *gin.Context) {
		grabS := viper.GetInt64("Interval.GrabS")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"config": gin.H{
				"grabS": grabS,
			},
		})
	})
	// 获取登录二维码
	r.GET("/api/get/qrcode", func(c *gin.Context) {
		// 直接读取图片数据返回
		// 判断二维码是否最新
		if qrcodeLatest {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "success",
			})
		}else{
			// 等待
			c.JSON(http.StatusOK, gin.H{
				"code": 8888,
				"msg": "请等待二维码抓取",
			})
		}
	})

	// 检查登录状态
	r.GET("/api/login/status", func(c *gin.Context) {
		b, msg, _ := getLoginStatus()
		if !b {
			c.JSON(http.StatusOK, gin.H{
				"code": 4003,
				"msg": msg,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "success",
			})
		}
	})
	// 获取数据
	r.GET("/api/get/data", func(c *gin.Context) {
		b, msg, list := getLoginStatus()
		if !b {
			c.JSON(http.StatusOK, gin.H{
				"code": 4003,
				"msg": msg,
			})
			return
		}
		var respList []ListData
		for _, l := range list {
			var ll []api.ScreenProductDetailRespData
			result := api.ScreenProductDetail(l.LiveRoomID, l.LiveAppID, "bind_time", false)
			if result.St == 0 {
				ll = append(ll, result.Data)
			}
			respList = append(respList, ListData{
				Mch:  l,
				List: ll,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"list": respList,
			"updated_at": utils.TimeFormat("Y-m-d H:i:s", time.Now()),
		})
	})
	err := r.Run(":" + viper.GetString("Server.Port")) // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}

type ListData struct {
	Mch api.ListQuickviewResult `json:"mch"`
	List []api.ScreenProductDetailRespData `json:"list"`
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
		// 保存二维码数据
		//file, err := os.OpenFile(utils.FolderPath + "/tmp/qrcode.tmp", os.O_CREATE | os.O_RDWR, 0775)
		//if err != nil {
		//	return err
		//}
		//defer file.Close()
		//_, err = file.Write([]byte(str))
		//if err != nil {
		//	return
		//}
		str = strings.Replace(str, `data:image/png;base64,`, ``, 1)
		qB, _ := base64.StdEncoding.DecodeString(str)
		file, err := os.OpenFile(utils.FolderPath + "/tmp/qrcode.png", os.O_CREATE | os.O_RDWR, 0775)
		if err != nil {
			return err
		}
		_, err = file.Write(qB)
		if err != nil {
			return
		}
		qrcodeLatest = true
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
	c, err := getCsv(c, utils.FolderPath + "/数据/" + nickName + "/商品", "商品数据", []string{
		"抓取时间", "名称", "价格", "商品点击率", "成交订单数", "成交金额", "成交转化率",
	})
	return c, err
}