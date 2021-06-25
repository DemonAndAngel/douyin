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
var RoomUrlInfoPath = ""
var RoomsDataPath = ""
var UVPath = ""
var TemplatesPath = ""

var mRoomUrlInfo *sync.RWMutex

func init() {
	mRoomUrlInfo = new(sync.RWMutex)
	var err error
	FolderPath, err = osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}
	//FolderPath = "."
	QrcodePath = FolderPath + "/tmp/qrcode.png"
	CookiesPath = FolderPath + "/cookies.tmp"
	RoomUrlInfoPath = FolderPath + "/tmp/room_url_info.tmp"
	RoomsDataPath = FolderPath + "/tmp/rooms/%s"
	UVPath = FolderPath + "/uv.tmp"
	TemplatesPath = FolderPath + "/templates"
}

// KeepFloat64 保留几位小数
func KeepFloat64(val float64, num int) float64 {
	r, _ := strconv.ParseFloat(KeepFloat64ToString(val, num), 64)
	return r
}
func KeepFloat64ToString(val float64, num int) string {
	return fmt.Sprintf(fmt.Sprintf("%%.%df", num), val)
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

func SetUV(f float64) (err error) {
	// 3. 存储到临时文件
	if err = ioutil.WriteFile(UVPath, []byte(strconv.FormatFloat(f, 'E', -1, 64)), 0755); err != nil {
		return
	}
	return
}
func GetUV() (f float64, err error) {
	if _, _err := os.Stat(UVPath); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(UVPath)
	if err != nil {
		return
	}
	// 反序列化
	str := string(b)
	f, _ = strconv.ParseFloat(str, 64)
	return
}

type RoomUrlInfo struct {
	LiveUrl string `json:"live_url"`
	Rooms []RoomsDataUrlInfo `json:"rooms_data_url_info"`
}

type RoomsDataUrlInfo struct {
	Nickname string `json:"nickname"` // 昵称
	StartTime time.Time `json:"start_time"` // 开播时间
	RoomId string `json:"room_id"`
	AppId int `json:"app_id"`
	BaseInfoUrl string `json:"base_info_url"`
	ProductDetailUrl string `json:"product_detail_url"`
	LiveDetailUrl string `json:"live_detail_url"`
}
func SaveRoomLiveUrl(url string) (err error) {
	mRoomUrlInfo.Lock()
	defer mRoomUrlInfo.Unlock()
	// 先取出来
	info, err := GetRoomUrlInfo()
	if err != nil {
		return
	}
	// 再存
	info.LiveUrl = url
	b, _ := json.Marshal(info)
	if err = ioutil.WriteFile(RoomUrlInfoPath, b, 0755); err != nil {
		return
	}
	return
}
func SaveRoomsDataUrlInfo(room RoomsDataUrlInfo) (err error) {
	mRoomUrlInfo.Lock()
	defer mRoomUrlInfo.Unlock()
	// 先取出来
	info, err := GetRoomUrlInfo()
	if err != nil {
		return
	}
	// 再存
	bo := false
	for k, r := range info.Rooms {
		if r.RoomId == room.RoomId {
			if room.BaseInfoUrl != "" {
				r.BaseInfoUrl = room.BaseInfoUrl
			}
			if room.ProductDetailUrl != "" {
				r.ProductDetailUrl = room.ProductDetailUrl
			}
			if room.LiveDetailUrl != "" {
				r.LiveDetailUrl = room.LiveDetailUrl
			}
			if room.Nickname != "" {
				r.Nickname = room.Nickname
			}
			if !room.StartTime.IsZero() {
				r.StartTime = room.StartTime
			}
			info.Rooms[k] = r
			bo = true
			break
		}
	}
	if !bo {
		info.Rooms = append(info.Rooms, room)
	}
	b, _ := json.Marshal(info)
	if err = ioutil.WriteFile(RoomUrlInfoPath, b, 0755); err != nil {
		return
	}
	return
}
func GetRoomUrlInfo() (info RoomUrlInfo, err error) {
	if _, _err := os.Stat(RoomUrlInfoPath); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(RoomUrlInfoPath)
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &info)
	return
}