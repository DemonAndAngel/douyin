package utils

import (
	"encoding/csv"
	"os"
	"time"
)

type Csv struct {
	File *os.File
	W *csv.Writer
	Path string
	Name string
	CreateTime time.Time
}

func NewCsv(filepath string, filename string, createTime time.Time, title []string) (c *Csv, err error) {
	if createTime.IsZero() {
		createTime = time.Now()
	}
	filename = filename + "_" + TimeFormat("Y-m-d", createTime) + ".csv"
	_, err = os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		err = nil
		_ = os.MkdirAll(filepath, os.ModePerm)
	}
	if err != nil {
		return
	}
	file, err := os.OpenFile(filepath + "/" + filename, os.O_APPEND, os.ModePerm)
	b := false
	if err != nil && os.IsNotExist(err) {
		err = nil
		file, err = os.Create(filepath + "/" + filename)
		if err != nil {
			return
		}
		// 写入UTF-8 BOM，防止中文乱码
		_, _ = file.WriteString("\xEF\xBB\xBF")
		b = true
	}
	if err != nil {
		return
	}
	w := csv.NewWriter(file)
	if b {
		_ = w.Write(title)
	}
	c = &Csv{
		File: file,
		W:    w,
		Path: filepath,
		Name: filename,
		CreateTime: createTime,
	}
	return
}

func (c *Csv) Close() (err error) {
	err = c.File.Close()
	return
}