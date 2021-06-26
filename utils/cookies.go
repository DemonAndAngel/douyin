package utils

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
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

func ChromedpLoadCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 设置cookies
		if len(LoadCookies()) > 0 {
			return network.SetCookies(LoadCookies()).Do(ctx)
		}
		return
	}
}

func HasCookies() bool {
	// 如果cookies临时文件不存在则直接跳过
	_, err := os.Stat(CookiesPath)
	if os.IsNotExist(err) {
		return false
	}else{
		return true
	}
}

func LoadCookies() (cookies []*network.CookieParam) {
	// 如果cookies临时文件不存在则直接跳过
	if _, _err := os.Stat(CookiesPath); os.IsNotExist(_err) {
		return
	}
	// 如果存在则读取cookies的数据
	b, err := ioutil.ReadFile(CookiesPath)
	if err != nil {
		return
	}
	// 反序列化
	err = json.Unmarshal(b, &cookies)
	return
}

// SaveCookies 保存Cookies
func SaveCookies(ctx context.Context) (err error) {
	// cookies的获取对应是在devTools的network面板中
	// 1. 获取cookies
	tmpCookies, err := network.GetAllCookies().Do(ctx)
	if err != nil {
		return
	}
	var cookies []*network.CookieParam
	for _, tmp := range tmpCookies {
		cookies = append(cookies, &network.CookieParam{
			Name:         tmp.Name,
			Value:        tmp.Value,
			URL:          "",
			Domain:       tmp.Domain,
			Path:         tmp.Path,
			Secure:       tmp.Secure,
			HTTPOnly:     tmp.HTTPOnly,
			SameSite:     tmp.SameSite,
			Priority:     tmp.Priority,
			SameParty:    tmp.SameParty,
			SourceScheme: tmp.SourceScheme,
			SourcePort:   tmp.SourcePort,
			//Expires:
		})
	}

	// 2. 序列化
	b, err := json.Marshal(cookies)
	if err != nil {
		return
	}
	// 3. 存储到临时文件
	if err = ioutil.WriteFile(CookiesPath, b, 0755); err != nil {
		return
	}
	return
}