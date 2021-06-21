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
		UserType int `json:"user_type"`
		UserIsLogin bool `json:"user_is_login"`
		LoginFrom int `json:"login_from"`
	} `json:"data"`
}

func UserTrack() (result UserTrackResp) {
	client := &http.Client{}
	req := NewRequest("GET", USER_TRACK,nil)
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