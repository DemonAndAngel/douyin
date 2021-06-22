package utils

import (
	"encoding/json"
	"fmt"
	"github.com/kardianos/osext"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var FolderPath = ""

func init() {
	var err error
	FolderPath, err = osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}
	FolderPath = "."
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
	if err = ioutil.WriteFile(FolderPath + "/tmp/uv.tmp", []byte(strconv.FormatFloat(f, 'E', -1, 64)), 0755); err != nil {
		return
	}
	return
}
func GetUV() (f float64, err error) {
	if _, _err := os.Stat(FolderPath + "/tmp/uv.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(FolderPath + "/tmp/uv.tmp")
	if err != nil {
		return
	}
	// 反序列化
	str := string(b)
	f, _ = strconv.ParseFloat(str, 64)
	return
}

type RoomUrlInfo struct {
	RoomId string `json:"room_id"`
	BaseInfoUrl string `json:"base_info_url"`
	ProductDetailUrl string `json:"product_detail_url"`
}
func SaveRoomInfoUrl(room RoomUrlInfo) (err error) {
	// 先取出来
	rooms, err := GetRoomInfoUrls()
	if err != nil {
		return
	}
	// 再存
	bo := false
	for k, r := range rooms {
		if r.RoomId == room.RoomId {
			if room.BaseInfoUrl != "" {
				r.BaseInfoUrl = room.BaseInfoUrl
			}
			if room.ProductDetailUrl != "" {
				r.ProductDetailUrl = room.ProductDetailUrl
			}
			rooms[k] = r
			bo = true
			break
		}
	}
	if !bo {
		rooms = append(rooms, room)
	}
	b, _ := json.Marshal(rooms)
	if err = ioutil.WriteFile(FolderPath + "/tmp/room_url_info.tmp", b, 0755); err != nil {
		return
	}
	return
}
func GetRoomInfoUrls() (rooms []RoomUrlInfo, err error) {
	if _, _err := os.Stat(FolderPath + "/tmp/room_url_info.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(FolderPath + "/tmp/room_url_info.tmp")
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &rooms)
	return
}
func SaveUpdatedAt(t time.Time) (err error) {
	if err = ioutil.WriteFile(FolderPath + "/tmp/updated_at.tmp", []byte(TimeFormat("Y-m-d H:i:s", t)), 0755); err != nil {
		return
	}
	return
}
func GetUpdatedAt() (t time.Time, err error) {
	if _, _err := os.Stat(FolderPath + "/tmp/updated_at.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(FolderPath + "/tmp/updated_at.tmp")
	if err != nil {
		return
	}
	// 反序列化
	t = CreateTimeFormat("Y-m-d H:i:s", string(b))
	return
}