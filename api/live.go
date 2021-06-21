package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
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

