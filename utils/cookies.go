package utils

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/network"
	"io/ioutil"
	"net/http"
	"os"
)

type Cookie struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Domain string `json:"domain"`
	Path string `json:"path"`
	Expires int `json:"expires"`
	Size int `json:"size"`
	HTTPOnly bool `json:"httpOnly"`
	Secure bool `json:"secure"`
	Session bool `json:"session"`
	SameSite string `json:"sameSite"`
	Priority string `json:"priority"`
	SameParty bool `json:"sameParty"`
	SourceScheme string `json:"sourceScheme"`
	SourcePort int `json:"sourcePort"`
}

// LoadCookies 加载Cookies
func LoadCookies() (cookies []*http.Cookie) {
	// 如果cookies临时文件不存在则直接跳过
	if _, _err := os.Stat("./tmp/cookies.tmp"); os.IsNotExist(_err) {
		return
	}
	// 如果存在则读取cookies的数据
	cookiesData, err := ioutil.ReadFile("./tmp/cookies.tmp")
	if err != nil {
		return
	}
	// 反序列化
	cs := make(map[string][]Cookie)
	json.Unmarshal(cookiesData, &cs)
	cc, _ := cs["cookies"]
	for _, c := range cc {
		cookies = append(cookies, &http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			//Expires:    c.Expires,
			//RawExpires: ,
			//MaxAge:     ,
			Secure:     c.Secure,
			HttpOnly:   c.HTTPOnly,
			//SameSite:   ,
			//Raw:       ,
			//Unparsed:   nil,
		})
	}
	return
}
//func LoadCookies() chromedp.ActionFunc {
//	return func(ctx context.Context) (err error) {
//		// 如果cookies临时文件不存在则直接跳过
//		if _, _err := os.Stat("./tmp/cookies.tmp"); os.IsNotExist(_err) {
//			return
//		}
//		// 如果存在则读取cookies的数据
//		cookiesData, err := ioutil.ReadFile("cookies.tmp")
//		if err != nil {
//			return
//		}
//
//		// 反序列化
//		cookiesParams := network.SetCookiesParams{}
//		if err = cookiesParams.UnmarshalJSON(cookiesData); err != nil {
//			return
//		}
//		// 设置cookies
//		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
//	}
//}

// SaveCookies 保存Cookies
func SaveCookies(ctx context.Context) (err error) {
	// cookies的获取对应是在devTools的network面板中
	// 1. 获取cookies
	cookies, err := network.GetAllCookies().Do(ctx)
	if err != nil {
		return
	}
	// 2. 序列化
	cookiesData, err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
	if err != nil {
		return
	}
	// 3. 存储到临时文件
	if err = ioutil.WriteFile("./tmp/cookies.tmp", cookiesData, 0755); err != nil {
		return
	}
	return
}