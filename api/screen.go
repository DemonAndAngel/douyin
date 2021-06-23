package api

import (
	"douyin/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var m *sync.RWMutex

func init() {
	m = new(sync.RWMutex)
}

type ScreenBaseInfoResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data ScreenBaseInfoRespData `json:"data"`
}
type ScreenBaseInfoRespData struct  {
	Title string `json:"title"`
	CoverImgURI string `json:"cover_img_uri"`
	AppPlatform string `json:"app_platform"`
	StartTime string `json:"start_time"`
	LiveDuration string `json:"live_duration"`
	PayCnt struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"pay_cnt"`
	PayUcnt struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"pay_ucnt"`
	ProductClickToPayRate struct {
		Value float64 `json:"value"`
		Unit string `json:"unit"`
	} `json:"product_click_to_pay_rate"`
	Gpm struct {
		Value float64 `json:"value"`
		Unit string `json:"unit"`
	} `json:"gpm"`
	PayFansRatio struct {
		Value float64 `json:"value"`
		Unit string `json:"unit"`
	} `json:"pay_fans_ratio"`
	OnlineUserCnt struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"online_user_cnt"`
	OnlineUserUcnt struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"online_user_ucnt"`
	FansClubUcnt struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"fans_club_ucnt"`
	IncrFansCnt struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"incr_fans_cnt"`
	AvgWatchDuration struct {
		Value int `json:"value"`
		Unit string `json:"unit"`
	} `json:"avg_watch_duration"`
	LiveStatus int `json:"live_status"`
	EndTime int `json:"end_time"`
	Gmv int `json:"gmv"`
	AppID int `json:"app_id"`
	LiveAppID int `json:"live_app_id"`
	IsTouxi bool `json:"is_touxi"`
}

func ScreenBaseInfo(url string) (result ScreenBaseInfoResp) {
	client := &http.Client{}
	req := NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	result.St = 500
	json.Unmarshal(body, &result)
	return
}

type ScreenProductDetailResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data ScreenProductDetailRespData `json:"data"`
}
type ScreenProductDetailRespData struct {
	IndexGroups []struct {
		GroupDisplay string `json:"group_display"`
		GroupName string `json:"group_name"`
		List []struct {
			IndexDisplay string `json:"index_display"`
			IndexName string `json:"index_name"`
		} `json:"list"`
	} `json:"index_groups"`
	IndexSelected []string `json:"index_selected"`
	DataHead []struct {
		IndexDisplay string `json:"index_display"`
		IndexName string `json:"index_name"`
	} `json:"data_head"`
	DataResult []struct {
		Product struct {
			Link string `json:"link"`
			Promotion bool `json:"promotion"`
			ImageURI string `json:"image_uri"`
			ID string `json:"id"`
		} `json:"product"`
		Title string `json:"title"`
		CurrMinPrice string `json:"curr_min_price"`
		ProductClickInLiveRate string `json:"product_click_in_live_rate"`
		PayInLiveOrderCnt string `json:"pay_in_live_order_cnt"`
		PayInLiveOrderProductGmv string `json:"pay_in_live_order_product_gmv"`
		ProductClickToPayRate string `json:"product_click_to_pay_rate"`
	} `json:"data_result"`
}

func ScreenProductDetail(url string) (result ScreenProductDetailResp) {
	client := &http.Client{}
	req := NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	result.St = 500
	json.Unmarshal(body, &result)
	return
}


type ScreenRoomBoardV2Resp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data ScreenRoomBoardV2RespData `json:"data"`
}
type ScreenRoomBoardV2RespData struct {
	PopularityData []struct {
		IndexDisplay string `json:"index_display"`
		Value struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"value"`
		ChangeValue struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"change_value"`
	} `json:"popularity_data"`
	ProductData []struct {
		IndexDisplay string `json:"index_display"`
		Value struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"value"`
		ChangeValue struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"change_value"`
		IndexTip string `json:"index_tip,omitempty"`
	} `json:"product_data"`
	InteractiveData []struct {
		IndexDisplay string `json:"index_display"`
		Value struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"value"`
		ChangeValue struct {
			Value float64 `json:"value"`
			Unit string `json:"unit"`
		} `json:"change_value"`
		IndexTip string `json:"index_tip,omitempty"`
	} `json:"interactive_data"`
	TradeData []struct {
		IndexDisplay string `json:"index_display"`
		Value struct {
			Value float64 `json:"value"`
			Unit string `json:"unit"`
		} `json:"value"`
		ChangeValue struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"change_value"`
		IndexTip string `json:"index_tip,omitempty"`
	} `json:"trade_data"`
	AfterSaleData []struct {
		IndexDisplay string `json:"index_display"`
		Value struct {
			Value int `json:"value"`
			Unit string `json:"unit"`
		} `json:"value"`
		ChangeValue struct {
			Value float64 `json:"value"`
			Unit string `json:"unit"`
		} `json:"change_value"`
		IndexTip string `json:"index_tip"`
	} `json:"after_sale_data"`
}

func ScreenRoomBoardV2(url string) (result ScreenRoomBoardV2Resp) {
	client := &http.Client{}
	req := NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	result.St = 500
	json.Unmarshal(body, &result)
	return
}

func ScreenSaveRoomBoardV2(roomId string, data ScreenRoomBoardV2RespData) (err error) {
	m.Lock()
	defer m.Unlock()
	// 2. 序列化
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	// 3. 存储到临时文件
	// 判断文件夹是否存在
	if _, _err := os.Stat(path); _err != nil && os.IsNotExist(_err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	if err = ioutil.WriteFile(path + "/screen_room_board_v2.tmp", b, 0755); err != nil {
		return
	}
	return
}


func ScreenSaveBaseInfo(roomId string, data ScreenBaseInfoRespData) (err error) {
	m.Lock()
	defer m.Unlock()
	// 2. 序列化
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	// 3. 存储到临时文件
	// 判断文件夹是否存在
	if _, _err := os.Stat(path); _err != nil && os.IsNotExist(_err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	if err = ioutil.WriteFile(path + "/screen_base_info.tmp", b, 0755); err != nil {
		return
	}
	return
}

func ScreenSaveProductDetail(roomId string, data ScreenProductDetailRespData) (err error) {
	m.Lock()
	defer m.Unlock()
	// 2. 序列化
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	// 3. 存储到临时文件
	// 判断文件夹是否存在
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path); _err != nil && os.IsNotExist(_err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	if err = ioutil.WriteFile(path + "/screen_product_detail.tmp", b, 0755); err != nil {
		return
	}
	return
}

func ScreenLoadBaseInfo(roomId string) (data ScreenBaseInfoRespData, err error) {
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path + "/screen_base_info.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(path + "/screen_base_info.tmp")
	if err != nil {
		return
	}
	// 反序列化
	err = json.Unmarshal(b, &data)
	return
}
func ScreenLoadProductDetail(roomId string) (data ScreenProductDetailRespData, err error) {
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path + "/screen_product_detail.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(path + "/screen_product_detail.tmp")
	if err != nil {
		return
	}
	// 反序列化
	err = json.Unmarshal(b, &data)
	return
}
func ScreenLoadRoomBoardV2(roomId string) (data ScreenRoomBoardV2RespData, err error) {
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path + "/screen_room_board_v2.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(path + "/screen_room_board_v2.tmp")
	if err != nil {
		return
	}
	// 反序列化
	err = json.Unmarshal(b, &data)
	return
}