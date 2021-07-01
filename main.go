package main

import (
	"context"
	"douyin/api"
	http "douyin/server"
	"douyin/utils"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/fsnotify/fsnotify"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
	"log"
	"math"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	playInfo, _ := utils.GetPlayInfoData()
	playInfo.BaseInfoUrl = ""
	playInfo.DataTrendUrl = ""
	playInfo.LiveRoomDashboardV2Url = ""
	playInfo.LiveDetailUrl = ""
	playInfo.ProductDetailUrl = ""
	utils.MyApp = &utils.App{
		PlayInfo: playInfo,
	}
	//os.RemoveAll(utils.FolderPath + "/tmp")
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
}
func SetConfig() {
	utils.MyConfig = &utils.Config{
		Interval: utils.ConfigInterval{
			GrabS:         viper.GetInt64("Interval.GrabS"),
			CheckLoginS:   viper.GetInt64("Interval.CheckLoginS"),
			QrcodeExpireS: viper.GetInt64("Interval.QrcodeExpireS"),
			UrlS:          viper.GetInt64("Interval.UrlS"),
			SaveS:         viper.GetInt64("Interval.SaveS"),
			SaveSEX: viper.GetInt64("Interval.SaveSEX"),
		},
		Server: utils.ConfigServer{
			Port: viper.GetInt("Server.Port"),
		},
	}
}
var p *ants.PoolWithFunc
type Write struct {
	Now time.Time
	App utils.App
	Data utils.Data
	UV utils.UV
	S int64
}
func main() {
	p, _ = ants.NewPoolWithFunc(3, writeCsv)
	defer p.Release()
	// 运行协程
	go func() {
		// 定时检查用户登录状态
		for {
			err := checkLogin()
			if err != nil {
				fmt.Println("check login error:" + err.Error())
			}
			if !utils.MyApp.IsLogin {
				fmt.Println("未登录 获取二维码")
				//_ = os.Remove(utils.CookiesPath)
				_ = os.Remove(utils.QrcodePath)
				err := getQrcode()
				if err != nil {
					fmt.Println("get qrcode error:" + err.Error())
				}
				// 立马检测登录
			} else {
				time.Sleep(time.Duration(utils.MyConfig.Interval.CheckLoginS) * time.Second)
			}
		}
	}()
	go func() {
		// 更新是否开播拉取地址
		for {
			if utils.MyApp.IsLogin {
				err := getLivePlayInfoUrl()
				if err != nil {
					fmt.Println("get live play info url error:" + err.Error())
				}
			}
		}

	}()
	go func() {
		// 监听是否开播 一秒一次
		for {
			if utils.MyApp.IsLogin && utils.MyApp.PlayInfo.Url != "" {
				err := getLivePlayInfo()
				if err != nil {
					fmt.Println("get live play info error:" + err.Error())
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		// 定时获取直播间数据地址
		for {
			if utils.MyApp.IsLogin && utils.MyApp.PlayInfo.RoomID != "" {
				err := getLiveDataUrls()
				if err != nil {
					fmt.Println("get live data urls error:" + err.Error())
				}
			}
		}
	}()
	go func() {
		// 定时获取数据
		for {
			if utils.MyApp.IsLogin && utils.MyApp.PlayInfo.RoomID != "" {
				err := getData(false)
				if err != nil {
					fmt.Println("get data error:" + err.Error())
				}
			}
		}
	}()
	http.Run()
}

func getLivePlayInfoUrl() (err error) {
	// 检测时间是否满足
	now := time.Now()
	if !utils.MyApp.LastPlayInfoUrlTime.IsZero() && now.Sub(utils.MyApp.LastPlayInfoUrlTime).Seconds() < float64(utils.MyConfig.Interval.UrlS) {
		return
	}
	playUrl := ""
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
			if strings.Index(ev.Request.URL, "api/livepc/playinfo") != -1 {
				playUrl = ev.Request.URL
				playInfo, err := utils.SavePlayInfoData(utils.PlayInfoData{
					Url:                        playUrl,
				}, "URL")
				if err == nil {
					utils.MyApp.PlayInfo = playInfo
					utils.MyApp.LastPlayInfoUrlTime = time.Now()
				}
			}
			break
		}
		// other needed network Event
	})
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate("https://buyin.jinritemai.com/dashboard/live/control"),
		waitUrl(&playUrl, 5),
	})
	return
}

func getLivePlayInfo() (err error) {
	// 检测是否正在直播是否满足
	result := api.LivePlayInfo(utils.MyApp.PlayInfo.Url)
	fmt.Println("play info result", result)
	if result.St == 0 && result.Code == 0 {
		// 更新当前直播数据
		playInfo, _err := utils.SavePlayInfoData(utils.PlayInfoData{
			NickName:                   result.Data.NickName,
			UserAvatar:                 result.Data.UserAvatar,
			StreamURL:                  result.Data.StreamURL,
			UserApp:                    result.Data.UserApp,
			RoomID:                     result.Data.RoomID,
			QrcodeSchemaURL:            result.Data.QrcodeSchemaURL,
			HasReleasedFissionActivity: result.Data.HasReleasedFissionActivity,
		}, "PLAY_INFO")
		if _err != nil {
			err = _err
			return
		}
		utils.MyApp.PlayInfo = playInfo
		// 重置地址
		if result.Data.RoomID != "" && utils.MyApp.PlayInfo.RoomID != result.Data.RoomID {
			// 需要重置地址
			fmt.Println("需要重置地址")
			_err = getLiveDataUrls()
			if _err != nil {
				err = _err
				return
			}
			// 重新获取一遍数据
			_err = getData(true)
			if _err != nil {
				err = _err
				return
			}
		}
	}
	return
}


func getData(q bool) (err error) {
	// 检测时间是否满足
	now := time.Now()
	if !q && !utils.MyApp.LastLiveDataTime.IsZero() && now.Sub(utils.MyApp.LastLiveDataTime).Seconds() < float64(utils.MyConfig.Interval.GrabS) {
		return
	}
	// 获取直播间
	app := utils.MyApp
	info := app.PlayInfo
	if info.BaseInfoUrl == "" || info.ProductDetailUrl == "" ||
		info.LiveDetailUrl == "" ||
		info.LiveRoomDashboardV2Url == "" ||
		info.DataTrendUrl == "" {
		return
	}
	bResp := api.ScreenBaseInfo(info.BaseInfoUrl)
	fmt.Println("bResp", bResp)
	if bResp.St == 0 {
		// 存数据
		_ = api.ScreenSaveBaseInfo(info.RoomID, bResp.Data)
	}
	pResp := api.ScreenProductDetail(info.ProductDetailUrl)
	fmt.Println("pResp", pResp)
	if pResp.St == 0 {
		// 存数据
		_ = api.ScreenSaveProductDetail(info.RoomID, pResp.Data)
	}
	dTResp := api.ScreenRoomDataTrend(info.DataTrendUrl, "trend_popularity")
	fmt.Println("dTResp", dTResp)
	if dTResp.St == 0 {
		// 存数据
		_ = api.ScreenSaveRoomDataTrend(info.RoomID, dTResp.Data, "trend_popularity")
	}
	lResp := api.ScreenLiveRoomDashboardV2(info.LiveRoomDashboardV2Url)
	fmt.Println("lResp", lResp)
	if lResp.St == 0 {
		// 存数据
		_ = api.ScreenSaveLiveRoomDashboardV2(info.RoomID, lResp.Data)
	}
	oResp := api.ScreenRoomOverview(info.LiveDetailUrl)
	fmt.Println("oResp", oResp)
	if oResp.St == 0 {
		// 存数据
		_ = api.ScreenSaveRoomOverview(info.RoomID, oResp.Data)
	}
	utils.MyApp.LastLiveDataTime = now
	uv, _ := utils.GetUV()
	// 检测是否需要写入数据
	if !q && utils.MyApp.LastSaveLiveDataTime.IsZero() || now.Sub(utils.MyApp.LastSaveLiveDataTime).Seconds() >= float64(utils.MyConfig.Interval.SaveS) {
		err = runWrite(*app, now, utils.MyConfig.Interval.SaveS, uv, bResp.Data, oResp.Data, lResp.Data, dTResp.Data)
		if err != nil {
			return
		}
		utils.MyApp.LastSaveLiveDataTime = now
	}
	if utils.MyConfig.Interval.SaveSEX != 0 {
		if !q && utils.MyApp.LastSaveEXLiveDataTime.IsZero() || now.Sub(utils.MyApp.LastSaveEXLiveDataTime).Seconds() >= float64(utils.MyConfig.Interval.SaveSEX) {
			err = runWrite(*app, now, utils.MyConfig.Interval.SaveSEX, uv, bResp.Data, oResp.Data, lResp.Data, dTResp.Data)
			if err != nil {
				return
			}
			utils.MyApp.LastSaveEXLiveDataTime = now
		}
	}
	return
}

func getLiveDataUrls() (err error) {
	// 检测时间是否满足
	now := time.Now()
	if !utils.MyApp.LastLiveListUrlsTime.IsZero() && now.Sub(utils.MyApp.LastLiveListUrlsTime).Seconds() < float64(utils.MyConfig.Interval.UrlS) {
		return
	}
	info := utils.MyApp.PlayInfo
	// 拼接浏览器地址
	//data := make(map[string]interface{})
	//data["live_room_id"] = info.RoomID
	//data["live_app_id"] = strconv.Itoa(info.UserApp)
	//data["source"] = "baiying_live_data"
	//pB, _ := json.Marshal(data)
	//params := string(pB)
	//winUrl := `https://compass.jinritemai.com/business_api/home/buyin_redirect/15101` + fmt.Sprintf("?params=%s", url.QueryEscape(params))
	winUrl := `https://compass.jinritemai.com/screen/list/talent/main` + fmt.Sprintf("?source=%s&live_app_id=%d&live_room_id=%s", "baiying_live_data", info.UserApp, info.RoomID)
	baseUrl := ""
	proUrl := ""
	detailUrl := ""
	dataTrendUrl := ""
	liveRoomDashboardV2Url := ""
	ctx, cancel, _ := genChromeCtx()
	defer func() {
		_ = chromedp.Cancel(ctx)
		cancel()
	}()
	// 添加监听
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			req := ev.Request
			if strings.Index(req.URL, "business_api/author/screen/base_info") != -1 {
				baseUrl = req.URL
				// 解析url并存储数据
				pageInfo, err := utils.SavePlayInfoData(utils.PlayInfoData{
					BaseInfoUrl:                req.URL,
					ProductDetailUrl:           "",
					LiveDetailUrl:              "",
					DataTrendUrl: "",
					LiveRoomDashboardV2Url: "",
				}, "ROOM_DATA_URL")
				if err != nil {
					fmt.Println(err)
				}
				utils.MyApp.PlayInfo = pageInfo
			} else if strings.Index(req.URL, "business_api/author/screen/product_detail") != -1 {
				proUrl = req.URL
				// 解析url并存储数据
				pageInfo, err := utils.SavePlayInfoData(utils.PlayInfoData{
					BaseInfoUrl:                "",
					ProductDetailUrl:           req.URL,
					LiveDetailUrl:              "",
					DataTrendUrl: "",
					LiveRoomDashboardV2Url: "",
				}, "ROOM_DATA_URL")
				if err != nil {
					fmt.Println(err)
				}
				utils.MyApp.PlayInfo = pageInfo
			} else if strings.Index(req.URL, "api/livepc/data/room/overview") != -1 {
				detailUrl = req.URL
				// 解析url并存储数据
				pageInfo, err := utils.SavePlayInfoData(utils.PlayInfoData{
					BaseInfoUrl:                "",
					ProductDetailUrl:           "",
					LiveDetailUrl:              req.URL,
					DataTrendUrl: "",
					LiveRoomDashboardV2Url: "",
				}, "ROOM_DATA_URL")
				if err != nil {
					fmt.Println(err)
				}
				utils.MyApp.PlayInfo = pageInfo
			} else if strings.Index(req.URL, "business_api/author/screen/data_trend") != -1 {
				// 去掉 index_selected 参数
				u, _ := url.Parse(req.URL)
				m, _ := url.ParseQuery(u.RawQuery)
				if _, ok := m["index_selected"]; ok {
					delete(m, "index_selected")
				}
				u.RawQuery = m.Encode()
				dataTrendUrl = u.String()
				// 解析url并存储数据
				pageInfo, err := utils.SavePlayInfoData(utils.PlayInfoData{
					BaseInfoUrl:                "",
					ProductDetailUrl:           "",
					LiveDetailUrl:              "",
					DataTrendUrl: req.URL,
					LiveRoomDashboardV2Url: "",
				}, "ROOM_DATA_URL")
				if err != nil {
					fmt.Println(err)
				}
				utils.MyApp.PlayInfo = pageInfo
			} else if strings.Index(req.URL, "business_api/author/live_detail/live_room/dashboard_v2") != -1 {
				liveRoomDashboardV2Url = req.URL
				// 解析url并存储数据
				pageInfo, err := utils.SavePlayInfoData(utils.PlayInfoData{
					BaseInfoUrl:                "",
					ProductDetailUrl:           "",
					LiveDetailUrl:              "",
					DataTrendUrl: "",
					LiveRoomDashboardV2Url: req.URL,
				}, "ROOM_DATA_URL")
				if err != nil {
					fmt.Println(err)
				}
				utils.MyApp.PlayInfo = pageInfo
			}
			break
		}
		// other needed network Event
	})
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate(`https://buyin.jinritemai.com/dashboard/livedata/detail` + fmt.Sprintf("?room_id=%s", info.RoomID)),
		waitUrl(&detailUrl, 10), // 等待url获取
	})
	if err != nil {
		return
	}
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate(winUrl),
		waitUrl(&baseUrl, 10), // 等待url获取
		waitUrl(&proUrl, 10),  // 等待url获取
		waitUrl(&dataTrendUrl, 10),  // 等待url获取
	})
	// 拿电商罗盘-整体看板-流量数据
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("https://compass.jinritemai.com/talent/live-statement?live_room_id=%s", info.RoomID)),
		waitUrl(&liveRoomDashboardV2Url, 10), // 等待url获取
	})
	fmt.Println("baseUrl", baseUrl)
	fmt.Println("proUrl", proUrl)
	fmt.Println("dataTrendUrl", dataTrendUrl)
	fmt.Println("detailUrl", detailUrl)
	fmt.Println("liveRoomDashboardV2Url", liveRoomDashboardV2Url)
	if err != nil {
		return
	}
	utils.MyApp.LastLiveListUrlsTime = now
	return
}

//func updateCookies() chromedp.ActionFunc {
//	return func(ctx context.Context) (err error) {
//		if err = utils.SaveCookies(ctx); err != nil {
//			return
//		}
//		return
//	}
//}

func waitUrl(url *string, waitS int) chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		now := time.Now()
		for {
			end := time.Now()
			if end.Sub(now).Seconds() > float64(waitS) || *url != "" {
				break
			}
			time.Sleep(10 * time.Millisecond)
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
	click1 := `#app > div > div.left-content > div.login-platform > div > ul > li:nth-child(1)`
	err = chromedp.Run(ctx, &chromedp.Tasks{
		chromedp.Navigate("https://buyin.jinritemai.com/mpa/account/login?log_out=1&type=24"),
		chromedp.WaitVisible(click1),
		chromedp.Click(click1),
		waitLogin(), // 等待登录
	})
	return
}
func waitLogin() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		defer func() {
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
			// 判断是否登录成功
			if err := chromedp.OuterHTML("html", &html).Do(ctx); err == nil {
				if dom, err := goquery.NewDocumentFromReader(strings.NewReader(html)); err == nil {
					if dom.Find("#portal > section > header > div > div.header-btns > div > div > div.btn-item-role-exchange-name > span").Length() > 0 {
						// 成功登录
						b = true
						break
					}
				}
			}
			if end.Sub(now).Seconds() >= float64(utils.MyConfig.Interval.QrcodeExpireS) {
				break
			}
			//// 保存cookies
			//if err = utils.SaveCookies(ctx); err != nil {
			//	return
			//}
			//// 用新的cookies去请求
			//err = checkLogin()
			//if err != nil {
			//	return
			//}
			//if utils.MyApp.IsLogin || end.Sub(now).Seconds() >= float64(utils.MyConfig.Interval.QrcodeExpireS) {
			//	break
			//}
			time.Sleep(100 * time.Millisecond)
		}
		if b {
			talentUrl := ""
			chromedp.ListenTarget(ctx, func(ev interface{}) {
				switch ev := ev.(type) {
				case *network.EventRequestWillBeSent:
					req := ev.Request
					if strings.Index(req.URL, "talent") != -1 {
						talentUrl = req.URL
					}
					break
				}
			})
			if err = chromedp.Run(ctx, &chromedp.Tasks{
				chromedp.Navigate("https://creator.douyin.com/business_api/home/buyin_redirect/10100"),
				waitUrl(&talentUrl, 5),
			}); err != nil {
				return
			}
			if talentUrl == "" {
				return errors.New("电商罗盘登录失败")
			}
			// 保存cookies
			if err = utils.SaveCookies(ctx); err != nil {
				return
			}
		}
		return
	}
}
func waitLoginCreator() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {

		return
	}
}

// 检测登录
func checkLogin() error {
	// 判断是否存在cookie
	if !utils.HasCookies() {
		utils.MyApp.IsLogin = false
		return nil
	}
	result := api.GetUser()
	b, _ := json.Marshal(result)
	fmt.Println("result", string(b))
	if result.St != 0 || result.Code != 0 || result.Data.UserRole == 0 {
		utils.MyApp.IsLogin = false
	}else{
		utils.MyApp.IsLogin = true
	}
	return nil
}

func genChromeCtx() (context.Context, context.CancelFunc, error) {
	// 打开浏览器
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		// 以默认配置的数组为基础，覆写headless参数
		// 当然也可以根据自己的需要进行修改，这个flag是浏览器的设置
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.Flag("blink-settings", "imagesEnabled=false"),
			chromedp.UserAgent(utils.GenUserAgent().Value),
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
func getLiveCsv(info utils.PlayInfoData, h string) (*utils.Csv, error) {
	now := time.Now()
	path := utils.FolderPath+"/数据/"+info.NickName + "/" + utils.TimeFormat("Y年m月d日H时i分", info.StartTime)
	return utils.NewCsv(path,
		"间隔" + h + "秒的数据", now, []string{
			"抓取时间",
			"成交件数", "成交人数", "新增粉丝数", "累计观看人数", "GMV", "商品曝光人数", "商品点击人数", "引流品金额（低于10块）", "非引流品金额",
			"实时刷单金额", "预期UV价值", "实时uv价值", "订单转化率", "成交人数转化率", "转粉率", "购物车点击率", "客单价", "成交粉丝占比",
			"人均看播时长",
			"直播间曝光人数", "直播画面转化率", "离开直播间人数", "实时在线人数", "进入直播间人数",
			"预期订单转化率", "预期转粉率", "预期购物车点击率", "预期直播画面转化率",
		})
}
func writeCsv(t interface{}) {
	w := t.(Write)
	info := w.App.PlayInfo
	data := w.Data
	uv := w.UV
	c, err := getLiveCsv(info, strconv.FormatInt(w.S, 10))
	if err != nil {
		return
	}
	// 写
	err = c.Write(data.ToWriteStrings(w.Now, &uv))
	return
}
func runWrite(app utils.App, t time.Time, s int64, uv utils.UV, b api.ScreenBaseInfoRespData, o api.ScreenRoomOverviewRespData,
	ld api.LiveRoomDashboardV2RespData, dt api.ScreenRoomDataTrendRespData) (err error) {
	zbjbgrs := 0
	for _, l := range ld.PopularityData {
		if l.IndexDisplay == "直播间曝光人数" {
			zbjbgrs = l.Value.Value
		}
	}
	// 取出最後一條
	leaveUcnt := 0
	onlineUserCnt := 0
	watchUcnt := 0
	if len(dt.TrendPopularity.Value) > 0 {
		y := dt.TrendPopularity.Value[len(dt.TrendPopularity.Value)-1].Y
		leaveUcnt = y.LeaveUcnt
		onlineUserCnt = y.OnlineUserCnt
		watchUcnt = y.WatchUcnt
	}
	ozhl := (float64(b.PayCnt.Value)/float64(b.OnlineUserUcnt.Value))*100
	if math.IsInf(ozhl, 0) {
		ozhl = 100
	}
	zfl := (float64(b.IncrFansCnt.Value)/float64(b.OnlineUserUcnt.Value))*100
	if math.IsInf(zfl, 0) {
		zfl = 100
	}
	gwcdjl := (float64(o.ProductStats.ClickUv)/float64(o.ProductStats.ShowUv))*100
	if math.IsInf(gwcdjl, 0) {
		gwcdjl = 100
	}
	zbhmzhl := (float64(b.OnlineUserUcnt.Value)/float64(zbjbgrs))*100
	if math.IsInf(zbhmzhl, 0) {
		zbhmzhl = 100
	}
	sssd := (float64(b.Gmv)/100)-float64(b.OnlineUserUcnt.Value)
	if math.IsInf(sssd, 0) {
		sssd = 0
	}
	suv := (float64(b.Gmv)/100)/float64(b.OnlineUserUcnt.Value)
	if math.IsInf(suv, 0) {
		suv = 0
	}
	cjrszhl := (float64(b.PayUcnt.Value)/float64(b.OnlineUserUcnt.Value))*100
	if math.IsInf(cjrszhl, 0) {
		cjrszhl = 100
	}
	kdj := float64(b.Gmv)/100/float64(b.PayCnt.Value)
	if math.IsInf(kdj, 0) {
		kdj = 0
	}
	err = p.Invoke(Write{
		Now:  t,
		App:  app,
		Data: utils.Data{
			Title: 			b.Title,
			PayCnt:         strconv.Itoa(b.PayCnt.Value),
			PayUcnt:        strconv.Itoa(b.PayUcnt.Value),
			IncrFansCnt:    strconv.Itoa(b.IncrFansCnt.Value),
			OnlineUserUcnt: strconv.Itoa(b.OnlineUserUcnt.Value),
			Gmv:            utils.KeepFloat64ToString(float64(b.Gmv)/100, 2),
			Exposure:       strconv.Itoa(o.ProductStats.ShowUv),
			Click:          strconv.Itoa(o.ProductStats.ClickUv),
			YinLiu:         "",
			FYinLiu:        "",
			SSSD:           utils.KeepFloat64ToString(sssd*uv.UV, 2),
			SUV:            utils.KeepFloat64ToString(suv, 2),
			OZHL:           utils.KeepFloat64ToString(ozhl, 2) + "%",
			CJRSZHL:        utils.KeepFloat64ToString(cjrszhl, 2) + "%",
			ZFL:            utils.KeepFloat64ToString(zfl, 2) + "%",
			GWCDJL:         utils.KeepFloat64ToString(gwcdjl, 2) + "%",
			KDJ:            utils.KeepFloat64ToString(kdj, 2),
			CJFSZB:         utils.KeepFloat64ToString(b.PayFansRatio.Value*100, 2) + "%",
			RJKBSC:         strconv.Itoa(b.AvgWatchDuration.Value) + "秒",
			ZBJBGRS: strconv.Itoa(zbjbgrs),
			ZBHMZHL: utils.KeepFloat64ToString(zbhmzhl, 2) + "%",
			LKZBJRS: strconv.Itoa(leaveUcnt),
			SSZXRS: strconv.Itoa(onlineUserCnt),
			JRZBJRS: strconv.Itoa(watchUcnt),
			DDZHLB: ozhl < uv.YDDZHL,
			ZFLB: zfl < uv.YZFL,
			GWCDJLB: gwcdjl < uv.YGWCDJL,
			ZBHMZHLB: zbhmzhl < uv.YZBHHZHL,
		},
		UV:   uv,
		S:    s,
	})
	if err != nil {
		return err
	}
	return
}
