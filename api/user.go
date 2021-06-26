package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetUserResp struct {
	Code int `json:"code"`
	Data struct {
		AccountAvatar string `json:"account_avatar"`
		AgreePro int `json:"agree_pro"`
		AnchorCouponMenuShow int `json:"anchor_coupon_menu_show"`
		BuyinAccountID string `json:"buyin_account_id"`
		CenterMenuShow int `json:"center_menu_show"`
		CheckStatus int `json:"check_status"`
		ChildStatus int `json:"child_status"`
		CompassFirstLevelMenuShow int `json:"compass_first_level_menu_show"`
		CompassSecondLevelMenuShow int `json:"compass_second_level_menu_show"`
		ContactNotSet int `json:"contact_not_set"`
		DarenPlazaPopup bool `json:"daren_plaza_popup"`
		DarenPlazaStatus int `json:"daren_plaza_status"`
		DoudianShopID int `json:"doudian_shop_id"`
		DrAuth struct {
			Num1128 struct {
				AuthorityShop int `json:"authority_shop"`
				AuthorityItem int `json:"authority_item"`
				AuthorityLive int `json:"authority_live"`
			} `json:"1128"`
		} `json:"dr_auth"`
		HasBindStar int `json:"has_bind_star"`
		OriginUID string `json:"origin_uid"`
		PlazaStatus int `json:"plaza_status"`
		Qianchuan int `json:"qianchuan"`
		SelectionPlaza int `json:"selection_plaza"`
		ShopID string `json:"shop_id"`
		ShopName string `json:"shop_name"`
		ShopType int `json:"shop_type"`
		ShopTypeChild int `json:"shop_type_child"`
		Shops []struct {
			AgreeProtocol int `json:"agree_protocol"`
			CheckStatus int `json:"check_status"`
			ShopID string `json:"shop_id"`
			ShopName string `json:"shop_name"`
			ShopType int `json:"shop_type"`
			ShopTypeChild int `json:"shop_type_child"`
			Status int `json:"status"`
			UserName string `json:"user_name"`
		} `json:"shops"`
		Status int `json:"status"`
		UserApp int `json:"user_app"`
		UserID string `json:"user_id"`
		UserIdentityType int `json:"user_identity_type"`
		UserName string `json:"user_name"`
		UserRole int `json:"user_role"`
	} `json:"data"`
	LogID string `json:"log_id"`
	Msg string `json:"msg"`
	St int `json:"st"`
}

func GetUser() (result GetUserResp) {
	client := &http.Client{}
	req := NewRequest("GET", "https://buyin.jinritemai.com/index/getUser",nil)
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
func UserBasicInfo(url string) (result UserBasicInfoResp) {
	client := &http.Client{}
	req := NewRequest("GET", url,nil)
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