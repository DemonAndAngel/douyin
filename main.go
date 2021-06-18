package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	goQrcode "github.com/skip2/go-qrcode"
	"github.com/asmcos/requests"
	"time"
)

type QrcodeResp struct {
	Data struct {
		AppName string `json:"app_name"`
		Qrcode string `json:"qrcode"`
		QrcodeIndexURL string `json:"qrcode_index_url"`
		Token string `json:"token"`
		WebName string `json:"web_name"`
		Status string `json:"status"`
	} `json:"data"`
	Description string `json:"description"`
	ErrorCode int `json:"error_code"`
	Message string `json:"message"`
}

func printQRCode(code []byte) (err error) {
	// 1. 因为我们的字节流是图像，所以我们需要先解码字节流
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		return
	}

	// 2. 然后使用gozxing库解码图片获取二进制位图
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return
	}

	// 3. 用二进制位图解码获取gozxing的二维码对象
	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return
	}
	// 4. 用结果来获取go-qrcode对象（注意这里我用了库的别名）
	qr, err := goQrcode.New(res.String(), goQrcode.Low)
	if err != nil {
		return
	}
	// 5. 输出到标准输出流
	fmt.Println(qr.ToSmallString(false))
	return
}

func main() {
	_, err := requests.Get("https://creator.douyin.com/")
	if err != nil {
		panic(err)
	}
	resp, err := requests.Get("https://sso.douyin.com/get_qrcode/?next=https:%2F%2Fcreator.douyin.com%2Fcreator-micro%2Fhome&aid=2906&service=https:%2F%2Fcreator.douyin.com&is_vcd=1")
	if err != nil {
		panic(err)
	}
	body := resp.Content()
	ret := QrcodeResp{}
	json.Unmarshal(body, &ret)
	if ret.ErrorCode != 0 {
		panic("获取二维码失败:" + ret.Message)
	}
	// 获取二维码
	qB, err := base64.StdEncoding.DecodeString(ret.Data.Qrcode)
	if err != nil {
		panic(err)
	}
	// 打印二维码
	err = printQRCode(qB)
	if err != nil {
		panic(err)
	}
	for {
		// 监听返回
		resp, err = requests.Get(`https://sso.douyin.com/check_qrconnect/?next=https:%2F%2Fcreator.douyin.com%2Fcreator-micro%2Fhome&service=https:%2F%2Fcreator.douyin.com%2F%3Flogintype%3Duser%26loginapp%3Ddouyin%26jump%3Dhttps:%2F%2Fcreator.douyin.com%2Fcreator-micro%2Fhome&correct_service=https:%2F%2Fcreator.douyin.com%2F%3Flogintype%3Duser%26loginapp%3Ddouyin%26jump%3Dhttps:%2F%2Fcreator.douyin.com%2Fcreator-micro%2Fhome&aid=2906&is_vcd=1&token=` + ret.Data.Token)
		if err != nil {
			panic(err)
		}
		body = resp.Content()
		fmt.Println(string(body))
		json.Unmarshal(body, &ret)
		if ret.ErrorCode != 0 {
			panic("获取二维码扫描结果失败:" + ret.Message)
		}
		if ret.Data.Status == "5" {
			// 过期
			panic("二维码过期")
		} else if ret.Data.Status == "2" {
			// 这里有可能是验证请求的
		} else if ret.Data.Status == "3" {
			fmt.Println("登录成功!")
			break
		}
		time.Sleep(1 * time.Second)
	}

}