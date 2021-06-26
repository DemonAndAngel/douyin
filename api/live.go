package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type LivePlayInfoResp struct {
	Code int `json:"code"`
	Data LivePlayInfoRespData `json:"data"`
	Extra struct {
		LogID string `json:"log_id"`
		Now int64 `json:"now"`
	} `json:"extra"`
	Msg string `json:"msg"`
	St int `json:"st"`
}
type LivePlayInfoRespData struct {
	NickName string `json:"nick_name"`
	UserAvatar string `json:"user_avatar"`
	StreamURL string `json:"stream_url"`
	UserApp int `json:"user_app"`
	RoomID string `json:"room_id"`
	QrcodeSchemaURL string `json:"qrcode_schema_url"`
	HasReleasedFissionActivity bool `json:"has_released_fission_activity"`
}
func LivePlayInfo(url string) (result LivePlayInfoResp) {
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

type ListQuickviewResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data struct {
		LiveGmv int `json:"live_gmv"`
		LiveCnt int `json:"live_cnt"`
		ProductCount int `json:"product_count"`
		LivePayComboNum int `json:"live_pay_combo_num"`
		UpdateTime int `json:"update_time"`
		HasMore bool `json:"has_more"`
		DataResult []ListQuickviewResult `json:"data_result"`
	} `json:"data"`
}
type ListQuickviewResult struct {
	AnchorID string `json:"anchor_id"`
	AnchorNickname string `json:"anchor_nickname"`
	AnchorImageURI string `json:"anchor_image_uri"`
	AnchorAwemeID string `json:"anchor_aweme_id"`
	FollowersCount int `json:"followers_count"`
	LiveRoomID string `json:"live_room_id"`
	LiveRoomTitle string `json:"live_room_title"`
	LiveRoomStartTime string `json:"live_room_start_time"`
	IsLiving bool `json:"is_living"`
	ProductCount int `json:"product_count"`
	PayGmv int `json:"pay_gmv"`
	CurrentOnlineUserCnt int `json:"current_online_user_cnt"`
	MaxOnlineUserCnt int `json:"max_online_user_cnt"`
	IsBindShop bool `json:"is_bind_shop"`
	AType int `json:"a_type"`
	IncrFansCnt int `json:"incr_fans_cnt"`
	WatchUcnt int `json:"watch_ucnt"`
	ShowBigScreen bool `json:"show_big_screen"`
	LiveAppID int `json:"live_app_id"`
}
func ListQuickview(url string) (result ListQuickviewResp) {
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

type LiveAnalysisResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data LiveAnalysisRespData `json:"data"`
}
type LiveAnalysisRespData struct {
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
		Sorted int `json:"sorted,omitempty"`
	} `json:"data_head"`
	DataResult []struct {
		ACU string `json:"a_c_u"`
		Author struct {
			Image string `json:"image"`
			NickName string `json:"nick_name"`
			OnLive bool `json:"on_live"`
			AwemeID string `json:"aweme_id"`
			AType int `json:"a_type"`
		} `json:"author"`
		LiveAppID int `json:"live_app_id"`
		LiveBeginTime struct {
			BeginTime int `json:"begin_time"`
			LiveDuration int `json:"live_duration"`
		} `json:"live_begin_time"`
		LiveRoomID string `json:"live_room_id"`
		PayGmv string `json:"pay_gmv"`
		PayProductCnt string `json:"pay_product_cnt"`
		ProductCount string `json:"product_count"`
		ShowBigScreen bool `json:"show_big_screen"`
		ShowDetail bool `json:"show_detail"`
	} `json:"data_result"`
	PageResult struct {
		PageSize int `json:"page_size"`
		PageNo int `json:"page_no"`
		Total int `json:"total"`
	} `json:"page_result"`
}

func LiveAnalysis(url string) (result ListQuickviewResp) {
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