package utils

import "time"

// 全局变量
var MyApp *App

type App struct {
	IsLogin              bool      `json:"is_login"`                 // 用户登录状态
	QrcodeLatest         bool      `json:"qrcode_latest"`            // 二维码过期状态
	LastLiveUrlTime      time.Time `json:"last_live_url_time"`       // 最后一次拉取直播间地址时间
	LastLiveListTime     time.Time `json:"last_live_list_time"`      // 最后一次更新直播间列表时间
	LastLiveListUrlsTime time.Time `json:"last_live_list_urls_time"` // 最后一次更新直播间列表数据地址时间
	LastLiveDataTime     time.Time `json:"last_live_data_time"`      // 最后一次拉取数据时间
	LastSaveLiveDataTime time.Time `json:"last_save_live_data_time"` // 最后一次写入数据间隔
	LastSaveEXLiveDataTime time.Time `json:"last_save_ex_live_data_time"` // 最后一次写入数据间隔
}
