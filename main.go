package main

import (
	"context"
	"douyin/api"
	"douyin/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/http/httptest"
	urlPkg "net/url"
	"os"
	"strconv"
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
	// 清除 cookies 外所有数据
	os.RemoveAll(utils.FolderPath + "/tmp/rooms")
	//os.Remove(utils.FolderPath + "/tmp/uv.tmp")
	os.Remove(utils.FolderPath + "/tmp/updated_at.tmp")
	os.Remove(utils.FolderPath + "/tmp/qrcode.png")
	os.Remove(utils.FolderPath + "/tmp/room_url_info.tmp")
}

var isLogin = false
var qrcodeLatest = false
var liveQuickviewUrl = ""

func main() {
	go func() {
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
		// 获取登录二维码
		r.GET("/api/get/qrcode", func(c *gin.Context) {
			// 直接读取图片数据返回
			// 判断二维码是否最新
			if qrcodeLatest {
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg": "success",
					"qrcodeLatest": qrcodeLatest,
					"isLogin": isLogin,
				})
			}else if isLogin {
				// 等待
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg": "已登录",
					"qrcodeLatest": qrcodeLatest,
					"isLogin": isLogin,
				})
			}else{
				// 等待
				c.JSON(http.StatusOK, gin.H{
					"code": 8888,
					"msg": "请等待二维码抓取",
					"qrcodeLatest": qrcodeLatest,
					"isLogin": isLogin,
				})
			}
		})

		// 检查登录状态
		r.GET("/api/login/status", func(c *gin.Context) {
			if !isLogin {
				c.JSON(http.StatusOK, gin.H{
					"code": 4003,
					"msg": "请登录",
				})
			}else{
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg": "success",
				})
			}
		})
		r.POST("/api/set/uv", func(c *gin.Context) {
			// 设置预期uv值
			str := c.PostForm("uv")
			f, _ := strconv.ParseFloat(str, 64)
			_ = utils.SetUV(f)
			c.JSON(http.StatusOK, gin.H{})
		})
		// 获取数据
		r.GET("/api/get/data", func(c *gin.Context) {
			if !isLogin {
				c.JSON(http.StatusOK, gin.H{
					"code": 4003,
					"msg": "请登录",
				})
				return
			}
			// 遍历room数据
			rooms, _ := utils.GetRoomInfoUrls()
			var list []map[string]interface{}
			for _, r := range rooms {
				baseInfo, _ := api.ScreenLoadBaseInfo(r.RoomId)
				productDetail, _ := api.ScreenLoadProductDetail(r.RoomId)
				list = append(list, map[string]interface{}{
					"base_info": baseInfo,
					"product_detail": productDetail,
				})

			}
			f, _ := utils.GetUV()
			updatedAt, _ := utils.GetUpdatedAt()
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "success",
				"list": list,
				"updated_at": utils.TimeFormat("Y-m-d H:i:s", updatedAt),
				"uv": f,
			})
		})
		err := r.Run(":" + viper.GetString("Server.Port")) // listen and serve on 0.0.0.0:8080
		if err != nil {
			panic(err)
		}

	}()


	// 打开浏览器
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
	ctx, cancel := chromedp.NewContext(
		ctx,
		// 设置日志方法
		chromedp.WithLogf(log.Printf),
	)
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			req := ev.Request
			if strings.Index(req.URL, "live_quickview") != -1 {
				liveQuickviewUrl = req.URL
			}
			break
		}
		// other needed network Event
	})
	if err := chromedp.Run(ctx, cookies()); err != nil {
		// 重启
		fmt.Println(err)
	}
	if !isLogin {
		if err := chromedp.Run(ctx, login()); err != nil {
			// 重启
			fmt.Println(err)
		}
	}
	// 开始爬数据
	if err := chromedp.Run(ctx, getData()); err != nil {
		// 重启
		fmt.Println(err)
	}
	cancel()
	// 获取所有直播间
	resp := api.ListQuickview(liveQuickviewUrl)
	if resp.St != 0 {
		fmt.Println(resp.Msg)
	}else{
		baseUrl := `https://compass.jinritemai.com/screen/list/shop/main`
		list := resp.Data.DataResult
		for _, l := range list {
			url := baseUrl + fmt.Sprintf("?live_room_id=%s&live_app_id=%d&source=shop_real_time", l.LiveRoomID, l.LiveAppID)
			// 打开浏览器
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
			ctx, _ = chromedp.NewContext(
				ctx,
				// 设置日志方法
				chromedp.WithLogf(log.Printf),
			)
			// 加载cookies
			if err := chromedp.Run(ctx, goToPagePre()); err != nil {
				// 重启
				fmt.Println(err)
			}
			// 添加监听
			chromedp.ListenTarget(ctx, func(ev interface{}) {
				switch ev := ev.(type) {
				case *network.EventRequestWillBeSent:
					req := ev.Request
					if strings.Index(req.URL, "screen/base_info") != -1 {
						// 解析url并存储数据
						u, _ := urlPkg.Parse(req.URL)
						m, _ := urlPkg.ParseQuery(u.RawQuery)
						err := utils.SaveRoomInfoUrl(utils.RoomUrlInfo{RoomId: m["live_room_id"][0], BaseInfoUrl: req.URL})
						if err != nil {
							fmt.Println(err)
						}
					} else if strings.Index(req.URL, "screen/product_detail") != -1 {
						// 解析url并存储数据
						u, _ := urlPkg.Parse(req.URL)
						m, _ := urlPkg.ParseQuery(u.RawQuery)
						err := utils.SaveRoomInfoUrl(utils.RoomUrlInfo{RoomId: m["live_room_id"][0], ProductDetailUrl: req.URL})
						if err != nil {
							fmt.Println(err)
						}
					}
					break
				}
				// other needed network Event
			})
			// 定时更新url
			go func() {
				for {
					if err := chromedp.Run(ctx, goToPage(url)); err != nil {
						// 重启
						fmt.Println(err)
					}
					time.Sleep(time.Duration(viper.GetInt64("Interval.UrlS")) * time.Second)
				}
			}()
			// 定时获取数据
			for {
				AGAIN:
				rooms, _ := utils.GetRoomInfoUrls()
				if len(rooms) <= 0 {
					continue
				}
				for _, r := range rooms {
					if r.BaseInfoUrl == "" || r.ProductDetailUrl == "" {
						goto AGAIN
					}
					// 保存数据
					bResp := api.ScreenBaseInfo(r.BaseInfoUrl)
					if bResp.St	== 0 {
						err := api.ScreenSaveBaseInfo(r.RoomId, bResp.Data)
						if err != nil {
							fmt.Println(err)
						}
					}
					pResp := api.ScreenProductDetail(r.ProductDetailUrl)
					if pResp.St	== 0 {
						err := api.ScreenSaveProductDetail(r.RoomId, pResp.Data)
						if err != nil {
							fmt.Println(err)
						}
					}
					// 保存更新时间
					err := utils.SaveUpdatedAt(time.Now())
					if err != nil {
						fmt.Println(err)
					}
				}
				time.Sleep(time.Duration(viper.GetInt64("Interval.GrabS")) * time.Second)
			}
		}
	}
}
func goToPagePre() chromedp.Tasks {
	return chromedp.Tasks{
		loadCookies(),
	}
}
func goToPage(url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
	}
}


func cookies() chromedp.Tasks {
	return chromedp.Tasks{
		loadCookies(),
		chromedp.Navigate("https://compass.jinritemai.com"),
		chromedp.Sleep(5 * time.Second),
		checkLogin(),
	}
}
func login() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click("#my-node > main > div > div.loginWrapper--2RTfW > div.rolePannel--25WFD > div.pannelRest--21kS3 > div > div.roleCard---v7UW.seller--1k9ot > svg > rect:nth-child(1)"),
		chromedp.Click("#my-node > main > div > div.loginWrapper--2RTfW > div > div:nth-child(2) > div > div > div > div.index_oauthLogin__1vWnv > div.index_oauthLoginBody__378Xb > div:nth-child(1)"),
		getCode(),
		waitLogin(),
	}
}

func getData() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click("#root > div > div.compassWrapper--3aIha > div.containerWrapper--kR6gq > div > div:nth-child(2) > div:nth-child(2) > div > div > div > div > div > div > div > div > div.title--eKQSF > a"),
		chromedp.Click("#root > div > div.compassWrapper--3aIha > div.containerWrapper--kR6gq > div > div.ecom-spin-nested-loading > div > div > div.cardContainer--MZhpq > div > div.info--1VBeX > div.link--915CX.active--38eYL > div:nth-child(2)"),
		run(),
	}
}
func run() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		return
	}
}

func waitLogin() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		if err = chromedp.WaitVisible(`#root > div > div.headerWrapper--7HYa6 > div > div.headerTools--TX8PU > div > div`).Do(ctx); err != nil {
			return
		}
		// 保存cookies
		if err = utils.SaveCookies(ctx); err != nil {
			return
		}
		return
	}
}

/**
 * 注册新tab标签的监听服务
 */
func addNewTabListener(ctx context.Context) <-chan target.ID {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	return chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		return info.URL != ""
	})
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

// 加载Cookies
func loadCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 设置cookies
		if len(utils.LoadCookies()) > 0 {
			return network.SetCookies(utils.LoadCookies()).Do(ctx)
		}
		return
	}
}

// 检测是否需要登录
func checkLogin() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 检测是否需要登录
		html := ""
		err = chromedp.InnerHTML(`html`, &html).Do(ctx)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			return
		}
		sel := dom.Find(`#root > div > div.headerWrapper--7HYa6 > div > div.headerTools--TX8PU > div > div`)
		if sel.Length() <= 0 {
			// 需要登录
			isLogin = false
			return
		}
		isLogin = true
		return
	}
}

