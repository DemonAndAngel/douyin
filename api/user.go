package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserBasicInfoResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data struct {
		Name string `json:"name"`
		Image string `json:"image"`
		DefaultRole int `json:"default_role"`
		ActiveRole int `json:"active_role"`
		RoleAuthor bool `json:"role_author"`
		RoleInstitution bool `json:"role_institution"`
		RoleShop bool `json:"role_shop"`
		IsShopChild bool `json:"is_shop_child"`
	} `json:"data"`
}

type UserTrackResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data struct {
		UUID string `json:"uuid"`
		UserID int64 `json:"user_id"`
		UserName string `json:"user_name"`
		UserType int `json:"user_type"`
		UserIsLogin bool `json:"user_is_login"`
		ToutiaoID string `json:"toutiao_id"`
		AccountGroupID string `json:"account_group_id"`
		UserAppID int `json:"user_app_id"`
		RoleType int `json:"role_type"`
		MinorID string `json:"minor_id"`
		LoginFrom int `json:"login_from"`
		ShopID string `json:"shop_id"`
		IsSelf int `json:"is_self"`
		IsBoss int `json:"is_boss"`
	} `json:"data"`
}

func UserTrack(url string) (result UserTrackResp) {
	client := &http.Client{}
	req := NewRequest("GET", url,nil)
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
func UserBasicInfo() (result UserBasicInfoResp) {
	client := &http.Client{}
	req := NewRequest("GET", USER_BASIC_INFO,nil)
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