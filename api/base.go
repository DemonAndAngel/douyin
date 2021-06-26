package api

import (
	"douyin/utils"
	"github.com/spf13/viper"
	"io"
	"net/http"
)


func NewRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url,body)
	// 添加头
	req.Header.Add("User-Agent", viper.GetString("System.UserAgent"))
	// 添加cookie
	cookies := utils.LoadCookies()
	for _, c := range cookies {
		req.AddCookie(&http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			//Expires:    time.Time{},
			//RawExpires: "",
			//MaxAge:     0,
			Secure:     c.Secure,
			HttpOnly:   c.HTTPOnly,
			//SameSite:   0,
			//Raw:        "",
			//Unparsed:   nil,
		})
	}
	return req
}