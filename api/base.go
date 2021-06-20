package api

import (
	"douyin/utils"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

const (
	BASE_URL = `https://compass.jinritemai.com`
	USER_TRACK = BASE_URL + `/business_api/home/track`
	USER_BASIC_INFO = BASE_URL + `/business_api/home/user_basic_info`
	Live_Quickview = BASE_URL + `/business_api/shop/core_data/live_quickview?date_type=1`

	// 实时大屏
	SCREEN_BASE_URL = `https://compass.jinritemai.com`
	SCREEN_BASE_INFO = SCREEN_BASE_URL + `/business_api/shop/screen/base_info`
	SCREEN_PRODUCT_DETAIL = SCREEN_BASE_URL + `/business_api/shop/screen/product_detail`

)

func NewRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url,body)
	// 添加头
	req.Header.Add("User-Agent", viper.GetString("System.UserAgent"))
	// 添加cookie
	cookies := utils.LoadCookies()
	for _, c := range cookies {
		req.AddCookie(c)
	}
	return req
}