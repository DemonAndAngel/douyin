package utils

import (
	"fmt"
	"github.com/kardianos/osext"
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
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if format == time.RFC3339 {
		t, _ := time.ParseInLocation(time.RFC3339, now, loc)
		return t
	} else {
		format = strings.Replace(format, "Y", "2006", -1)
		format = strings.Replace(format, "m", "01", -1)
		format = strings.Replace(format, "d", "02", -1)
		format = strings.Replace(format, "H", "15", -1)
		format = strings.Replace(format, "i", "04", -1)
		format = strings.Replace(format, "s", "05", -1)
		t, _ := time.ParseInLocation(format, now, loc)
		return t
	}
}