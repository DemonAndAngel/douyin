package http

import (
	"douyin/api"
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Run() {
	// 开启gin服务
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 加载html文件
	r.LoadHTMLGlob(utils.TemplatesPath + "/*.html")
	r.Static("/js", utils.TemplatesPath + "/js")
	r.Static("/css", utils.TemplatesPath + "/css")
	r.Static("/img", utils.TemplatesPath + "/img")
	r.Static("/tmp", utils.FolderPath + "/tmp")
	r.StaticFile("/favicon.ico", utils.TemplatesPath + "/favicon.ico")
	// 初始化路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// 获取登录二维码
	r.GET("/api/get/qrcode", func(c *gin.Context) {
		// 直接读取图片数据返回
		// 判断二维码是否最新
		if utils.MyApp.QrcodeLatest {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "success",
				"qrcodeLatest": utils.MyApp.QrcodeLatest,
				"isLogin": utils.MyApp.IsLogin,
			})
		}else if utils.MyApp.IsLogin {
			// 等待
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "已登录",
				"qrcodeLatest": utils.MyApp.QrcodeLatest,
				"isLogin": utils.MyApp.IsLogin,
			})
		}else{
			// 等待
			c.JSON(http.StatusOK, gin.H{
				"code": 8888,
				"msg": "请等待二维码抓取",
				"qrcodeLatest": utils.MyApp.QrcodeLatest,
				"isLogin": utils.MyApp.IsLogin,
			})
		}
	})

	// 检查登录状态
	r.GET("/api/login/status", func(c *gin.Context) {
		if !utils.MyApp.IsLogin {
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
		uv := utils.UV{}
		if c.ShouldBind(&uv) == nil {
			_ = utils.SetUV(uv)
		}
		c.JSON(http.StatusOK, gin.H{
			"uv": uv,
		})
	})
	r.GET("/api/get/uv", func(c *gin.Context) {
		// 设置预期uv值
		uv, _ := utils.GetUV()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": gin.H{
				"uv": uv,
			},
		})
	})
	// 获取数据
	r.GET("/api/get/last/data", func(c *gin.Context) {
		if !utils.MyApp.IsLogin {
			c.JSON(http.StatusOK, gin.H{
				"code": 4003,
				"msg": "请登录",
			})
			return
		}
		// 遍历room数据
		info := utils.MyApp.PlayInfo
		// 先取出最后一场直播
		uv, _ := utils.GetUV()
		if info.RoomID == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "success",
				"data": nil,
			})
			return
		}else{
			b, _ := api.ScreenLoadBaseInfo(info.RoomID)
			o, _ := api.ScreenLoadRoomOverview(info.RoomID)
			d, _ := api.ScreenLoadRoomDataTrend(info.RoomID, "trend_popularity")
			lv2, _ := api.ScreenLoadLiveRoomDashboardV2Resp(info.RoomID)
			if b.Title == "" {
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg": "success",
					"data": nil,
				})
				return
			}else{
				zbjbgrs := 0
				for _, l := range lv2.PopularityData {
					if l.IndexDisplay == "直播间曝光人数" {
						zbjbgrs = l.Value.Value
					}
				}
				// 取出最後一條
				leaveUcnt := 0
				onlineUserCnt := 0
				watchUcnt := 0
				if len(d.TrendPopularity.Value) > 0 {
					y := d.TrendPopularity.Value[len(d.TrendPopularity.Value)-1].Y
					leaveUcnt = y.LeaveUcnt
					onlineUserCnt = y.OnlineUserCnt
					watchUcnt = y.WatchUcnt
				}
				ozhl := (float64(b.PayCnt.Value)/float64(b.OnlineUserUcnt.Value))*100
				zfl := (float64(b.IncrFansCnt.Value)/float64(b.OnlineUserUcnt.Value))*100
				gwcdjl := (float64(o.ProductStats.ClickUv)/float64(o.ProductStats.ShowUv))*100
				zbhmzhl := (float64(b.OnlineUserUcnt.Value)/float64(zbjbgrs))*100
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg": "success",
					"data": utils.Data{
						Title: 			b.Title,
						UpdatedAt:      utils.TimeFormat("Y-m-d H:i:s", utils.MyApp.LastLiveDataTime),
						PayCnt:         strconv.Itoa(b.PayCnt.Value),
						PayUcnt:        strconv.Itoa(b.PayUcnt.Value),
						IncrFansCnt:    strconv.Itoa(b.IncrFansCnt.Value),
						OnlineUserUcnt: strconv.Itoa(b.OnlineUserUcnt.Value),
						Gmv:            utils.KeepFloat64ToString(float64(b.Gmv)/100, 2),
						Exposure:       strconv.Itoa(o.ProductStats.ShowUv),
						Click:          strconv.Itoa(o.ProductStats.ClickUv),
						YinLiu:         "",
						FYinLiu:        "",
						SSSD:           utils.KeepFloat64ToString((float64(b.Gmv)/100)-float64(b.OnlineUserUcnt.Value)*uv.UV, 2),
						UV:             utils.KeepFloat64ToString(uv.UV, 2),
						SUV:            utils.KeepFloat64ToString((float64(b.Gmv)/100)/float64(b.OnlineUserUcnt.Value), 2),
						OZHL:           utils.KeepFloat64ToString(ozhl, 2) + "%",
						CJRSZHL:        utils.KeepFloat64ToString((float64(b.PayUcnt.Value)/float64(b.OnlineUserUcnt.Value))*100, 2) + "%",
						ZFL:            utils.KeepFloat64ToString(zfl, 2) + "%",
						GWCDJL:         utils.KeepFloat64ToString(gwcdjl, 2) + "%",
						KDJ:            utils.KeepFloat64ToString(float64(b.Gmv)/100/float64(b.PayCnt.Value), 2),
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

						StreamUrl: info.StreamURL,
					},
				})
				return
			}
		}
	})
	err := r.Run(fmt.Sprintf(":%d", utils.MyConfig.Server.Port)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}