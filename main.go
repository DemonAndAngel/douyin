package main

import (
	"douyin/utils"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"context"
	"os"
	"strings"
	"time"
)

// 全局变量
var app *App
type App struct {
	IsLogin bool `json:"is_login"`
}
func init() {
	app = new(App)
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
	//os.RemoveAll(utils.FolderPath + "/tmp/rooms")
	//os.Remove(utils.FolderPath + "/tmp/updated_at.tmp")
	//os.Remove(utils.FolderPath + "/tmp/qrcode.png")
	//os.Remove(utils.FolderPath + "/tmp/room_url_info.tmp")
}

func main() {
	go checkLogin()
	time.Sleep(100 * time.Second)
}

func genChromeCtx() (context.Context, context.CancelFunc, error) {
	// 打开浏览器
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		// 以默认配置的数组为基础，覆写headless参数
		// 当然也可以根据自己的需要进行修改，这个flag是浏览器的设置
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			//chromedp.Flag("blink-settings", "imagesEnabled=false"),
			chromedp.UserAgent(viper.GetString("System.UserAgent")),
		)...,
	)
	ctx, _ = chromedp.NewContext(
		ctx,
		// 设置日志方法
		chromedp.WithLogf(log.Printf),
	)
	// 加载cookie
	err := chromedp.Run(ctx, &chromedp.Tasks{
		utils.ChromedpLoadCookies(),
	})
	return ctx, cancel, err
}

// 检测登录
func checkLogin() error {
	isLogin := false
	ctx, cancel, _ := genChromeCtx()
	defer func() {
		_ = chromedp.Cancel(ctx)
		cancel()
	}()
	// 添加监听
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			if strings.Index(ev.Request.URL, "/business_api/shop/core_data/live") != -1 {
				isLogin = true
			}
			break
		}
		// other needed network Event
	})

	err := chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate("https://compass.jinritemai.com"),
		chromedp.Sleep(time.Duration(viper.GetInt64("Interval.CheckLoginS")) * time.Second), // 等待10s
	})
	// 更新状态
	app.IsLogin = isLogin
	fmt.Println("%v", app)
	return err
}
