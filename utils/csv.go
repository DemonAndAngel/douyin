package utils

import (
	"encoding/csv"
	"os"
	"sync"
	"time"
)
var MyCsvMap map[string]*Csv

type Csv struct {
	File *os.File
	W *csv.Writer
	Path string
	Name string
	CreateTime time.Time
	M *sync.RWMutex
}

func init() {
	MyCsvMap = make(map[string]*Csv)
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
	path := filepath + "/" + filename
	if cc, ok := MyCsvMap[path]; ok {
		return cc, nil
	}
	file, err := os.OpenFile(path, os.O_APPEND, os.ModePerm)
	b := false
	if err != nil && os.IsNotExist(err) {
		err = nil
		file, err = os.Create(path)
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
		err = w.Write(title)
		w.Flush()
		if err != nil {
			return
		}
	}
	c = &Csv{
		File: file,
		W:    w,
		Path: filepath,
		Name: filename,
		CreateTime: createTime,
		M: new(sync.RWMutex),
	}
	MyCsvMap[path] = c
	return
}

func (c *Csv) Write(data []string) (err error) {
	c.M.Lock()
	defer c.M.Unlock()
	err = c.W.Write(data)
	if err != nil {
		return
	}
	c.W.Flush()
	return
}
func (c *Csv) WriteAll(datas [][]string) (err error) {
	c.M.Lock()
	defer c.M.Unlock()
	err = c.W.WriteAll(datas)
	if err != nil {
		return
	}
	c.W.Flush()
	return
}

func (c *Csv) Close() (err error) {
	err = c.File.Close()
	return
}