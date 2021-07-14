package utils

import (
	"time"
)

// 全局变量
var MyApp *App

type App struct {
	IsLogin              bool      `json:"is_login"`                 // 用户登录状态
	QrcodeLatest         bool      `json:"qrcode_latest"`            // 二维码过期状态
	LastPlayInfoUrlTime  time.Time `json:"last_play_info_url_time"` // 最后一次获取playinfo接口地址时间
	PlayInfo PlayInfoData `json:"play_info"` // 直播间数据



	LastLiveListUrlsTime time.Time `json:"last_live_list_urls_time"` // 最后一次更新直播间列表数据地址时间
	LastLiveDataTime     time.Time `json:"last_live_data_time"`      // 最后一次拉取数据时间
	LastSaveLiveDataTime time.Time `json:"last_save_live_data_time"` // 最后一次写入数据间隔
	LastSaveEXLiveDataTime time.Time `json:"last_save_ex_live_data_time"` // 最后一次写入数据间隔
}

type Data struct {
	Title string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	PayCnt string `json:"pay_cnt"`
	PayUcnt string `json:"pay_ucnt"`
	IncrFansCnt string `json:"incr_fans_cnt"`
	OnlineUserUcnt string `json:"online_user_ucnt"`
	Gmv string `json:"gmv"`
	Exposure string `json:"exposure"`
	Click string `json:"click"`
	YinLiu string `json:"yin_liu"`
	FYinLiu string `json:"f_yin_liu"`
	SSSD string `json:"sssd"`

	SUV string `json:"suv"`
	OZHL string `json:"ozhl"`
	CJRSZHL string `json:"cjrszhl"`
	ZFL string `json:"zfl"`
	GWCDJL string `json:"gwcdjl"`
	KDJ string `json:"kdj"`
	CJFSZB string `json:"cjfszb"`
	RJKBSC string `json:"rjkbsc"`

	// live v2
	ZBJBGRS string `json:"zbjbgrs"`
	ZBHMZHL string `json:"zbhmzhl"` // LJGKRS/ZBJBGRS
	// data_trend
	LKZBJRS string `json:"lkzbjrs"`
	SSZXRS string `json:"sszxrs"`
	JRZBJRS string `json:"jrzbjrs"`


	DDZHLB bool `json:"ddzhlb"`
	ZFLB bool `json:"zflb"`
	GWCDJLB bool `json:"gwcdjlb"`
	ZBHMZHLB bool `json:"zbhmzhlb"`

	ZBJPLS string `json:"zbjpls"`
	ZBPLCSZZL string `json:"zbplcszzl"`

	StreamUrl string `json:"stream_url"`
}

func (data *Data) ToWriteStrings(t time.Time, uv *UV) []string {
	return []string{
		TimeFormat("Y-m-d H:i:s", t),
		data.PayCnt,
		data.PayUcnt,
		data.IncrFansCnt,
		data.OnlineUserUcnt,
		data.Gmv,
		data.Exposure,
		data.Click,
		data.YinLiu,
		data.FYinLiu,
		data.SSSD,
		KeepFloat64ToString(uv.UV, 2),
		data.SUV,
		data.OZHL,
		data.CJRSZHL,
		data.ZFL,
		data.GWCDJL,
		data.KDJ,
		data.CJFSZB,
		data.RJKBSC,
		data.ZBJBGRS,
		data.ZBHMZHL,
		data.LKZBJRS,
		data.SSZXRS,
		data.JRZBJRS,
		data.ZBJPLS,
		data.ZBPLCSZZL,
		KeepFloat64ToString(uv.YDDZHL, 2) + "%",
		KeepFloat64ToString(uv.YZFL, 2) + "%",
		KeepFloat64ToString(uv.YGWCDJL, 2) + "%",
		KeepFloat64ToString(uv.YZBHHZHL, 2) + "%",
	}
}