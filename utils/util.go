package utils

import (
	"encoding/json"
	"fmt"
	"github.com/kardianos/osext"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var FolderPath = ""
var QrcodePath = ""
var CookiesPath = ""
var PlayInfoPath = ""
var RoomsDataPath = ""
var UVPath = ""
var TemplatesPath = ""

var mPlay *sync.RWMutex
var mUv *sync.RWMutex

func init() {
	mPlay = new(sync.RWMutex)
	mUv = new(sync.RWMutex)
	var err error
	FolderPath, err = osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}
	//FolderPath = "."
	QrcodePath = FolderPath + "/tmp/qrcode.png"
	CookiesPath = FolderPath + "/cookies.tmp"
	UVPath = FolderPath + "/uv.tmp"
	TemplatesPath = FolderPath + "/templates"
	PlayInfoPath = FolderPath + "/tmp/play_info.tmp"
	RoomsDataPath = FolderPath + "/tmp/rooms/%s"
	//RoomUrlInfoPath = FolderPath + "/tmp/room_url_info.tmp"
}

type UserAgentInfo struct {
	Value string `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	Index int `json:"index"`
}
var UserAgent *UserAgentInfo
var US = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 Edg/91.0.864.59",
}
func GenUserAgent() *UserAgentInfo {
	now := time.Now()
	index := 0
	if UserAgent != nil {
		if now.Sub(UserAgent.CreatedAt).Hours() < 3 {
			return UserAgent
		}
		index = UserAgent.Index + 1
		if index >= len(US) {
			index = 0
		}
	}
	UserAgent = &UserAgentInfo{
		Value: US[index],
		CreatedAt: now,
		Index: index,
	}
	return UserAgent
}

// KeepFloat64 保留几位小数
func KeepFloat64(val float64, num int) float64 {
	r, _ := strconv.ParseFloat(KeepFloat64ToString(val, num), 64)
	return r
}
func KeepFloat64ToString(val float64, num int) string {
	return fmt.Sprintf(fmt.Sprintf("%%.%df", num), val)
}
func KeepStringToFloat64(val string) float64 {
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

// TimeFormat 时间转日期
func TimeFormat(format string, t time.Time) string {
	if format != time.RFC3339 {
		format = strings.Replace(format, "Y", "2006", -1)
		format = strings.Replace(format, "m", "01", -1)
		format = strings.Replace(format, "d", "02", -1)
		format = strings.Replace(format, "H", "15", -1)
		format = strings.Replace(format, "i", "04", -1)
		format = strings.Replace(format, "s", "05", -1)
	}
	if t.IsZero() {
		return ""
	}
	return t.Format(format)
}

// CreateTimeFormat 日期转时间
func CreateTimeFormat(format string, now string) time.Time {
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	if format == time.RFC3339 {
		t, _ := time.Parse(time.RFC3339, now)
		return t
	} else {
		format = strings.Replace(format, "Y", "2006", -1)
		format = strings.Replace(format, "m", "01", -1)
		format = strings.Replace(format, "d", "02", -1)
		format = strings.Replace(format, "H", "15", -1)
		format = strings.Replace(format, "i", "04", -1)
		format = strings.Replace(format, "s", "05", -1)
		t, _ := time.Parse(format, now)
		return t
	}
}

type UV struct {
	UV float64 `json:"uv"`
	YDDZHL float64 `json:"yddzhl"`
	YZFL float64 `json:"yzfl"`
	YGWCDJL float64 `json:"ygwcdjl"`
	YZBHHZHL float64 `json:"yzbhhzhl"`
}

func SetUV(data UV) (err error) {
	mUv.Lock()
	defer mUv.Unlock()
	// 2. 序列化
	b, err := json.Marshal(data)
	// 3. 存储到临时文件
	if err = ioutil.WriteFile(UVPath, b, 0755); err != nil {
		return
	}
	return
}
func GetUV() (data UV, err error) {
	if _, _err := os.Stat(UVPath); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(UVPath)
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &data)
	return
}

type PlayInfoData struct {
	Url string `json:"url"`
	NickName string `json:"nick_name"`
	UserAvatar string `json:"user_avatar"`
	StreamURL string `json:"stream_url"`
	UserApp int `json:"user_app"`
	RoomID string `json:"room_id"` // 可以用这个判断是否开播
	QrcodeSchemaURL string `json:"qrcode_schema_url"`
	HasReleasedFissionActivity bool `json:"has_released_fission_activity"`
	StartTime time.Time `json:"start_time"` // 开播时间
	BaseInfoUrl string `json:"base_info_url"`
	ProductDetailUrl string `json:"product_detail_url"`
	LiveDetailUrl string `json:"live_detail_url"`
	DataTrendUrl string `json:"data_trend_url"`
	LiveRoomDashboardV2Url string `json:"live_room_dashboard_v_2_url"`
}
func SavePlayInfoData(newInfo PlayInfoData, t string) (info PlayInfoData, err error) {
	mPlay.Lock()
	defer mPlay.Unlock()
	// 先取出来
	info, err = GetPlayInfoData()
	if err != nil {
		return
	}
	switch t {
	case "URL":
		info.Url = newInfo.Url
		break
	case "PLAY_INFO":
		info.NickName = newInfo.NickName
		info.UserAvatar = newInfo.UserAvatar
		info.UserApp = newInfo.UserApp
		info.HasReleasedFissionActivity = newInfo.HasReleasedFissionActivity
		if newInfo.RoomID == "" {
			info.StreamURL = ""
		}else{
			info.QrcodeSchemaURL = newInfo.QrcodeSchemaURL
			info.StreamURL = newInfo.StreamURL
			if info.RoomID != newInfo.RoomID {
				// 重置roomid
				info.RoomID = newInfo.RoomID
				info.StartTime = time.Now()
			}
		}
		break
	case "ROOM_DATA_URL":
		if newInfo.BaseInfoUrl != "" {
			info.BaseInfoUrl = newInfo.BaseInfoUrl
		}
		if newInfo.ProductDetailUrl != "" {
			info.ProductDetailUrl = newInfo.ProductDetailUrl
		}
		if newInfo.LiveDetailUrl != "" {
			info.LiveDetailUrl = newInfo.LiveDetailUrl
		}
		if newInfo.DataTrendUrl != "" {
			info.DataTrendUrl = newInfo.DataTrendUrl
		}
		if newInfo.LiveRoomDashboardV2Url != "" {
			info.LiveRoomDashboardV2Url = newInfo.LiveRoomDashboardV2Url
		}
		break
	}
	b, _ := json.Marshal(info)
	if err = ioutil.WriteFile(PlayInfoPath, b, 0755); err != nil {
		return
	}
	return
}
func GetPlayInfoData() (info PlayInfoData, err error) {
	if _, _err := os.Stat(PlayInfoPath); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(PlayInfoPath)
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &info)
	return
}