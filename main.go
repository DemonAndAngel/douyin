package main

import (
	"context"
	"douyin/api"
	http "douyin/server"
	"douyin/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	utils.MyApp = new(utils.App)
	// 创建临时文件夹
	_, err := os.Stat(utils.FolderPath + "/tmp")
	if err != nil && os.IsNotExist(err) {
		err = nil
		err = os.MkdirAll(utils.FolderPath+"/tmp", os.ModePerm)
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
		SetConfig()
		fmt.Println("Config file changed:", e.Name)
	})
	SetConfig()
	// 清除 cookies 外所有数据
	//os.RemoveAll(utils.FolderPath + "/tmp/rooms")
	//os.Remove(utils.FolderPath + "/tmp/updated_at.tmp")
	//os.Remove(utils.FolderPath + "/tmp/qrcode.png")
	//os.Remove(utils.FolderPath + "/tmp/room_url_info.tmp")
}
func SetConfig() {
	utils.MyConfig = &utils.Config{
		Interval: utils.ConfigInterval{
			GrabS:         viper.GetInt64("Interval.GrabS"),
			CheckLoginS:   viper.GetInt64("Interval.CheckLoginS"),
			QrcodeExpireS: viper.GetInt64("Interval.QrcodeExpireS"),
			UrlS:          viper.GetInt64("Interval.UrlS"),
		},
		Server: utils.ConfigServer{
			Port: viper.GetInt("Server.Port"),
		},
		System: utils.ConfigSystem{
			UserAgent: viper.GetString("System.UserAgent"),
		},
	}
}

func main() {
	// 运行协程
	go func() {
		// 定时检查用户登录状态
		for {
			err := checkLogin()
			if err != nil {
				fmt.Println("check login error:" + err.Error())
			}
			if utils.MyApp.IsLogin {
				time.Sleep(time.Duration(utils.MyConfig.Interval.CheckLoginS) * time.Second)
			}else{
				time.Sleep(10 * time.Second)
			}
		}
	}()
	go func() {
		// 监听并获取最新二维码
		for {
			if utils.MyApp.FirstLogin && !utils.MyApp.IsLogin && !utils.MyApp.QrcodeLatest && !utils.MyApp.CheckLogin {
				//_ = os.Remove(utils.CookiesPath)
				_ = os.Remove(utils.QrcodePath)
				err := getQrcode()
				if err != nil {
					fmt.Println("get qrcode error:" + err.Error())
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		// 更新直播间拉取地址
		for {
			if utils.MyApp.FirstLogin && utils.MyApp.IsLogin {
				err := getLiveUrl()
				if err != nil {
					fmt.Println("get live url error:" + err.Error())
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		// 定时获取直播间列表
		for {
			if utils.MyApp.FirstLogin && utils.MyApp.IsLogin {
				err := getLiveList()
				if err != nil {
					fmt.Println("get live list error:" + err.Error())
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		// 定时获取直播间数据地址
		for {
			if utils.MyApp.FirstLogin && utils.MyApp.IsLogin {
				err := getLiveDataUrls()
				if err != nil {
					fmt.Println("get live data urls error:" + err.Error())
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		// 定时获取数据
		for {
			if utils.MyApp.FirstLogin && utils.MyApp.IsLogin {
				err := getData()
				if err != nil {
					fmt.Println("get data error:" + err.Error())
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	http.Run()
}

func getData() (err error) {
	// 检测时间是否满足
	now := time.Now()
	if !utils.MyApp.LastLiveDataTime.IsZero() && now.Sub(utils.MyApp.LastLiveDataTime).Seconds() < float64(utils.MyConfig.Interval.GrabS) {
		return
	}
	list, err := utils.GetRoomUrlInfo()
	if err != nil {
		return
	}
	if len(list.Rooms) <= 0 {
		return
	}
	for _, room := range list.Rooms {
		if room.BaseInfoUrl == "" || room.ProductDetailUrl == "" || room.LiveDetailUrl == "" {
			return
		}
		if room.BaseInfoUrl != "" {
			bResp := api.ScreenBaseInfo(room.BaseInfoUrl)
			if bResp.St != 0 {
				err = errors.New(bResp.Msg)
				return
			} else {
				// 存数据
				err = api.ScreenSaveBaseInfo(room.RoomId, bResp.Data)
				if err != nil {
					return
				}
			}
		}
		if room.ProductDetailUrl != "" {
			pResp := api.ScreenProductDetail(room.ProductDetailUrl)
			if pResp.St != 0 {
				err = errors.New(pResp.Msg)
				return
			} else {
				// 存数据
				err = api.ScreenSaveProductDetail(room.RoomId, pResp.Data)
				if err != nil {
					return
				}
			}
		}
		if room.LiveDetailUrl != "" {
			dResp := api.ScreenRoomBoardV2(room.LiveDetailUrl)
			if dResp.St != 0 {
				err = errors.New(dResp.Msg)
				return
			} else {
				// 存数据
				err = api.ScreenSaveRoomBoardV2(room.RoomId, dResp.Data)
				if err != nil {
					err = errors.New(dResp.Msg)
					return
				}
			}
		}
	}
	utils.MyApp.LastLiveDataTime = now
	// 检测是否需要写入数据
	if utils.MyApp.LastSaveLiveDataTime.IsZero() || now.Sub(utils.MyApp.LastSaveLiveDataTime).Seconds() >= float64(utils.MyConfig.Interval.SaveS) {
		for _, room := range list.Rooms {
			uv, _err := utils.GetUV()
			if _err != nil {
				err = _err
				return
			}
			c, _err := getLiveCsv(room)
			if _err != nil {
				err = _err
				return
			}
			// 拿数据
			b, _err := api.ScreenLoadBaseInfo(room.RoomId)
			if _err != nil {
				err = _err
				return
			}
			bo, _err := api.ScreenLoadRoomBoardV2(room.RoomId)
			if _err != nil {
				err = _err
				return
			}
			exposure := 0
			click := 0
			for _, boo := range bo.ProductData {
				if boo.IndexDisplay == "商品曝光人数" {
					exposure = boo.Value.Value
				} else if boo.IndexDisplay == "商品点击人数" {
					click = boo.Value.Value
				}
			}
			// 写
			err = c.Write([]string{
				utils.TimeFormat("Y-m-d H:i:s", now),
				strconv.Itoa(b.PayCnt.Value),
				strconv.Itoa(b.PayUcnt.Value),
				strconv.Itoa(b.IncrFansCnt.Value),
				strconv.Itoa(b.OnlineUserUcnt.Value),
				utils.KeepFloat64ToString(float64(b.Gmv)/100, 2),
				strconv.Itoa(exposure),
				strconv.Itoa(click),
				"",
				"",
				utils.KeepFloat64ToString((float64(b.Gmv)/100)-float64(b.OnlineUserUcnt.Value)*uv, 2),
				utils.KeepFloat64ToString(uv, 2),
				utils.KeepFloat64ToString((float64(b.Gmv)/100)/float64(b.OnlineUserUcnt.Value), 2),
				utils.KeepFloat64ToString((float64(b.PayCnt.Value)/float64(b.OnlineUserUcnt.Value))*100, 2) + "%",
				utils.KeepFloat64ToString((float64(b.PayUcnt.Value)/float64(b.OnlineUserUcnt.Value))*100, 2) + "%",
				utils.KeepFloat64ToString((float64(b.IncrFansCnt.Value)/float64(b.OnlineUserUcnt.Value))*100, 2) + "%",
				utils.KeepFloat64ToString((float64(click)/float64(exposure))*100, 2) + "%",
				utils.KeepFloat64ToString(float64(b.Gmv) / 100 /float64(b.PayCnt.Value), 2),
				utils.KeepFloat64ToString(b.PayFansRatio.Value*100, 2) + "%",
				strconv.Itoa(b.AvgWatchDuration.Value) + "秒",
			})
			if err != nil {
				return
			}
		}
		utils.MyApp.LastSaveLiveDataTime = now
	}
	return
}

func getLiveDataUrls() (err error) {
	// 检测时间是否满足
	now := time.Now()
	if !utils.MyApp.LastLiveListUrlsTime.IsZero() && now.Sub(utils.MyApp.LastLiveListUrlsTime).Seconds() < float64(utils.MyConfig.Interval.UrlS) {
		return
	}
	info, err := utils.GetRoomUrlInfo()
	if err != nil {
		return
	}
	if len(info.Rooms) > 0 {
		for _, room := range info.Rooms {
			// 拼接浏览器地址
			winUrl := `https://compass.jinritemai.com/screen/list/shop/main` + fmt.Sprintf("?live_room_id=%s&live_app_id=%d&source=shop_real_time",
				room.RoomId, room.AppId)
			baseUrl := ""
			proUrl := ""
			detailUrl := ""
			ctx, cancel, _ := genChromeCtx()
			// 添加监听
			chromedp.ListenTarget(ctx, func(ev interface{}) {
				switch ev := ev.(type) {
				case *network.EventRequestWillBeSent:
					req := ev.Request
					if strings.Index(req.URL, "screen/base_info") != -1 {
						baseUrl = req.URL
						// 解析url并存储数据
						err := utils.SaveRoomsDataUrlInfo(utils.RoomsDataUrlInfo{
							RoomId:           room.RoomId,
							BaseInfoUrl:      req.URL,
							ProductDetailUrl: "",
							LiveDetailUrl:    "",
						})
						if err != nil {
							fmt.Println(err)
						}
					} else if strings.Index(req.URL, "screen/product_detail") != -1 {
						proUrl = req.URL
						// 解析url并存储数据
						err := utils.SaveRoomsDataUrlInfo(utils.RoomsDataUrlInfo{
							RoomId:           room.RoomId,
							BaseInfoUrl:      "",
							ProductDetailUrl: req.URL,
							LiveDetailUrl:    "",
						})
						if err != nil {
							fmt.Println(err)
						}
					} else if strings.Index(req.URL, "business_api/shop/live_room/room_board_v2") != -1 {
						detailUrl = req.URL
						// 解析url并存储数据
						err := utils.SaveRoomsDataUrlInfo(utils.RoomsDataUrlInfo{
							RoomId:           room.RoomId,
							BaseInfoUrl:      "",
							ProductDetailUrl: "",
							LiveDetailUrl:    req.URL,
						})
						if err != nil {
							fmt.Println(err)
						}
					}
					break
				}
				// other needed network Event
			})
			err = chromedp.Run(ctx, &chromedp.Tasks{
				chromedp.Navigate(winUrl),
				waitUrl(&baseUrl, 10), // 等待url获取
				waitUrl(&proUrl, 10),  // 等待url获取
			})
			if err != nil {
				cancel()
				return
			}
			err = chromedp.Run(ctx, &chromedp.Tasks{
				chromedp.Navigate(`https://compass.jinritemai.com/shop/live-detail` + fmt.Sprintf("?live_room_id=%s", room.RoomId)),
				waitUrl(&detailUrl, 10), // 等待url获取
			})
			cancel()
			if err != nil {
				return
			}
		}
		utils.MyApp.LastLiveListUrlsTime = now
	}
	return
}

func getLiveList() (err error) {
	// 检测时间是否满足
	now := time.Now()
	info, err := utils.GetRoomUrlInfo()
	if err != nil {
		return
	}
	if len(info.Rooms) > 0 {
		if !utils.MyApp.LastLiveListTime.IsZero() && now.Sub(utils.MyApp.LastLiveListTime).Seconds() < float64(utils.MyConfig.Interval.UrlS) {
			return
		}
	}else{
		if !utils.MyApp.LastLiveListTime.IsZero() && now.Sub(utils.MyApp.LastLiveListTime).Seconds() < 5 {
			return
		}
	}
	if info.LiveUrl != "" {
		// 更新数据
		lResp := api.ListQuickview(info.LiveUrl)
		if lResp.St != 0 {
			err = errors.New(lResp.Msg)
			return
		}
		for _, l := range lResp.Data.DataResult {
			err = utils.SaveRoomsDataUrlInfo(utils.RoomsDataUrlInfo{
				RoomId:    l.LiveRoomID,
				AppId:     l.LiveAppID,
				StartTime: utils.CreateTimeFormat("Y/m/d H:i:s", l.LiveRoomStartTime),
				Nickname:  l.AnchorNickname,
			})
			if err != nil {
				return
			}
		}
		utils.MyApp.LastLiveListTime = now
	}
	return
}

func getLiveUrl() (err error) {
	// 检测时间是否满足
	now := time.Now()
	if !utils.MyApp.LastLiveUrlTime.IsZero() && now.Sub(utils.MyApp.LastLiveUrlTime).Seconds() < float64(utils.MyConfig.Interval.UrlS) {
		return
	}
	url := ""
	// 打开浏览器
	ctx, cancel, _ := genChromeCtx()
	defer func() {
		_ = chromedp.Cancel(ctx)
		cancel()
	}()
	// 添加监听
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			if strings.Index(ev.Request.URL, "/business_api/shop/core_data/live_quickview") != -1 {
				url = ev.Request.URL
				err = utils.SaveRoomLiveUrl(ev.Request.URL)
				if err == nil {
					utils.MyApp.LastLiveUrlTime = time.Now()
				}
			}
			break
		}
		// other needed network Event
	})
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate("https://compass.jinritemai.com/shop/real-time"),
		waitUrl(&url, 5),
	})
	return
}

func waitUrl(url *string, waitS int) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		now := time.Now()
		for {
			end := time.Now()
			if end.Sub(now).Seconds() > float64(waitS) || *url != "" {
				// 五秒直接超时
				break
			}
			time.Sleep(1 * time.Second)
		}
		return
	}
}

// 获取二维码
func getQrcode() (err error) {
	ctx, cancel, _ := genChromeCtx()
	defer func() {
		_ = chromedp.Cancel(ctx)
		cancel()
	}()
	click1 := `#my-node > main > div > div.loginWrapper--2RTfW > div.rolePannel--25WFD > div.pannelRest--21kS3 > div > div.roleCard---v7UW.seller--1k9ot > svg > rect:nth-child(1)`
	click2 := `#my-node > main > div > div.loginWrapper--2RTfW > div > div:nth-child(2) > div > div > div > div.index_oauthLogin__1vWnv > div.index_oauthLoginBody__378Xb > div:nth-child(1)`
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate("https://compass.jinritemai.com/login"),
		chromedp.WaitVisible(click1),
		chromedp.Click(click1),
		chromedp.WaitVisible(click2),
		chromedp.Click(click2),
		waitLogin(), // 等待登录
	})
	return
}
func waitLogin() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		defer func(){
			utils.MyApp.QrcodeLatest = false
		}()
		// 1. 用于存储图片的字节切片
		html := ""
		// 2. 截图
		// 注意这里需要注明直接使用ID选择器来获取元素（chromedp.ByID）
		if err = chromedp.OuterHTML(`#root > div > div.content-container > div.auth-container > div.auth-qr-container > div.qr-container > img`,
			&html).Do(ctx); err != nil {
			return
		}
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			return
		}
		str, ok := dom.Find("img").Attr("src")
		if !ok {
			err = errors.New("二维码获取失败")
			return
		}
		str = strings.Replace(str, `data:image/png;base64,`, ``, 1)
		qB, _ := base64.StdEncoding.DecodeString(str)
		file, err := os.OpenFile(utils.QrcodePath, os.O_CREATE|os.O_RDWR, 0775)
		if err != nil {
			return err
		}
		_, err = file.Write(qB)
		if err != nil {
			return
		}
		utils.MyApp.QrcodeLatest = true
		// 检测二维码是否过期
		b := false
		now := time.Now()
		for {
			end := time.Now()
			if err := chromedp.OuterHTML("html", &html).Do(ctx); err == nil {
				if dom, err := goquery.NewDocumentFromReader(strings.NewReader(html)); err == nil {
					if dom.Find("#root > div > div.headerWrapper--7HYa6 > div > div.headerTools--TX8PU > div > div").Length() > 0 {
						// 成功登录
						b = true
						break
					}
				}
			}
			if end.Sub(now).Seconds() >= float64(utils.MyConfig.Interval.QrcodeExpireS) {
				break
			}
			time.Sleep(1 * time.Second)
		}
		// 保存cookies
		if b {
			if err = utils.SaveCookies(ctx); err != nil {
				return
			}
		}
		return
	}
}

// 检测登录
func checkLogin() error {
	utils.MyApp.CheckLogin = true
	url := ""
	ctx, cancel, _ := genChromeCtx()
	defer func() {
		utils.MyApp.FirstLogin = true
		utils.MyApp.CheckLogin = false
		_ = chromedp.Cancel(ctx)
		cancel()
	}()
	// 添加监听
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			if strings.Index(ev.Request.URL, "/business_api/shop/core_data/live") != -1 {
				url = ev.Request.URL
			}
			break
		}
		// other needed network Event
	})

	err := chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate("https://compass.jinritemai.com"),
		waitUrl(&url, 10), // 等待10s
	})
	// 更新状态
	if url != "" {
		utils.MyApp.IsLogin = true
	}else{
		utils.MyApp.IsLogin = false
	}
	fmt.Println("%v", utils.MyApp)
	return err
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
			chromedp.UserAgent(utils.MyConfig.System.UserAgent),
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

// 获取直播专用csv对象
func getLiveCsv(room utils.RoomsDataUrlInfo) (*utils.Csv, error) {
	now := time.Now()
	return utils.NewCsv(utils.FolderPath+"/数据/"+room.Nickname+"_"+utils.TimeFormat("Ymd", room.StartTime),
		"数据", now, []string{
			"抓取时间",
			"成交件数", "成交人数", "新增粉丝数", "累计观看人数", "GMV", "商品曝光人数", "商品点击人数", "引流品金额（低于10块）", "非引流品金额",
			"实时刷单金额", "预期UV价值", "实时uv价值", "订单转化率", "成交人数转化率", "转粉率", "购物车点击率", "客单价", "成交粉丝占比",
			"人均看播时长",
		})
}
