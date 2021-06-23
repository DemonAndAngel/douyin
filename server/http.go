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
	r.LoadHTMLGlob(utils.TemplatesPath + "/html/*")
	r.Static("/js", utils.TemplatesPath + "/js")
	r.Static("/css", utils.TemplatesPath + "/css")
	r.Static("/tmp", utils.FolderPath + "/tmp")
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
		str := c.PostForm("uv")
		f, _ := strconv.ParseFloat(str, 64)
		_ = utils.SetUV(f)
		c.JSON(http.StatusOK, gin.H{})
	})
	// 获取数据
	r.GET("/api/get/data", func(c *gin.Context) {
		if !utils.MyApp.IsLogin {
			c.JSON(http.StatusOK, gin.H{
				"code": 4003,
				"msg": "请登录",
			})
			return
		}
		// 遍历room数据
		info, _ := utils.GetRoomUrlInfo()
		var list []map[string]interface{}
		for _, r := range info.Rooms {
			baseInfo, _ := api.ScreenLoadBaseInfo(r.RoomId)
			productDetail, _ := api.ScreenLoadProductDetail(r.RoomId)
			list = append(list, map[string]interface{}{
				"base_info": baseInfo,
				"product_detail": productDetail,
			})

		}
		f, _ := utils.GetUV()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"list": list,
			"updated_at": utils.TimeFormat("Y-m-d H:i:s", utils.MyApp.LastLiveDataTime),
			"uv": f,
		})
	})
	err := r.Run(fmt.Sprintf(":%d", utils.MyConfig.Server.Port)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}