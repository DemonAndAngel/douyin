package api

import (
	"douyin/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

var baseInfoM *sync.RWMutex
var productDetailM *sync.RWMutex
var overviewM *sync.RWMutex
var dataTrendM *sync.RWMutex
var dashboardM *sync.RWMutex

func init() {
	baseInfoM = new(sync.RWMutex)
	productDetailM = new(sync.RWMutex)
	overviewM = new(sync.RWMutex)
	dataTrendM = new(sync.RWMutex)
	dashboardM = new(sync.RWMutex)
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

type ScreenRoomOverviewResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data ScreenRoomOverviewRespData `json:"data"`
	Extra struct {
		LogID string `json:"log_id"`
		Now int64 `json:"now"`
	} `json:"extra"`
}
type ScreenRoomOverviewRespData struct {
	FlowStats struct {
		AvgOnlineUv int `json:"avg_online_uv"`
		AvgWatchDuration int `json:"avg_watch_duration"`
		FansAvgWatchDuration int `json:"fans_avg_watch_duration"`
		MaxOnlineUv int `json:"max_online_uv"`
		RtOnlineUv int `json:"rt_online_uv"`
		WatchPv int `json:"watch_pv"`
		WatchUv int `json:"watch_uv"`
	} `json:"flow_stats"`
	InteractionStats struct {
		CommentNum int `json:"comment_num"`
		IncrFansNum int `json:"incr_fans_num"`
		LikeNum int `json:"like_num"`
		ShareNum int `json:"share_num"`
		Last1CommentNum int `json:"last1_comment_num"` // 自有字段 上一次直播评论数
		Last2CommentNum int `json:"last2_comment_num"` // 自有字段 上一次直播评论数
	} `json:"interaction_stats"`
	LiveStatus int `json:"live_status"`
	OrderStats struct {
		FansPayInLiveGmvRatio float64 `json:"fans_pay_in_live_gmv_ratio"`
		FansPayInLiveNumRatio float64 `json:"fans_pay_in_live_num_ratio"`
		PayInLiveGmv int `json:"pay_in_live_gmv"`
		PayInLiveNum int `json:"pay_in_live_num"`
		PayInLiveNumRatio float64 `json:"pay_in_live_num_ratio"`
	} `json:"order_stats"`
	ProductStats struct {
		ClickUv int `json:"click_uv"`
		FansClickUvRatio int `json:"fans_click_uv_ratio"`
		FansPayInLiveUvRatio float64 `json:"fans_pay_in_live_uv_ratio"`
		FansShowUvRatio int `json:"fans_show_uv_ratio"`
		PayInLiveUv int `json:"pay_in_live_uv"`
		ShowUv int `json:"show_uv"`
	} `json:"product_stats"`
	UpdateTs int `json:"update_ts"`
}
func ScreenRoomOverview(url string) (result ScreenRoomOverviewResp) {
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

type ScreenRoomDataTrendTPResp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data ScreenRoomDataTrendTPRespData `json:"data"`
}
type ScreenRoomDataTrendTPRespData struct {
	TrendPopularity struct {
		Unit struct {
			LeaveUcnt string `json:"leave_ucnt"`
			OnlineUserCnt string `json:"online_user_cnt"`
			WatchUcnt string `json:"watch_ucnt"`
		} `json:"unit"`
		Value []struct {
			X string `json:"x"`
			Y struct {
				LeaveUcnt int `json:"leave_ucnt"`
				OnlineUserCnt int `json:"online_user_cnt"`
				WatchUcnt int `json:"watch_ucnt"`
			} `json:"y"`
		} `json:"value"`
	} `json:"trend_popularity"`
}
func ScreenRoomDataTrendTP(url string) (result ScreenRoomDataTrendTPResp) {
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


type LiveRoomDashboardV2Resp struct {
	St int `json:"st"`
	Msg string `json:"msg"`
	Data LiveRoomDashboardV2RespData `json:"data"`
}
type LiveRoomDashboardV2RespData struct {
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
func ScreenLiveRoomDashboardV2(url string) (result LiveRoomDashboardV2Resp) {
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
func ScreenSaveLiveRoomDashboardV2(roomId string, data LiveRoomDashboardV2RespData) (err error) {
	dataTrendM.Lock()
	defer dataTrendM.Unlock()
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
	if err = ioutil.WriteFile(path + "/live_room_dashboard_v2.tmp", b, 0755); err != nil {
		return
	}
	return
}

func ScreenSaveRoomDataTrendTP(roomId string, data ScreenRoomDataTrendTPRespData) (err error) {
	dataTrendM.Lock()
	defer dataTrendM.Unlock()
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
	if err = ioutil.WriteFile(path + "/screen_data_data_trend-tp.tmp", b, 0755); err != nil {
		return
	}
	return
}

func ScreenSaveRoomOverview(roomId string, data ScreenRoomOverviewRespData, n time.Time, s1 time.Time, s2 time.Time) (err error) {
	overviewM.Lock()
	defer overviewM.Unlock()
	// 判断是否需要更新数据
	// 取出上一次数据
	o, err := ScreenLoadRoomOverview(roomId)
	if err != nil {
		return
	}
	if s1.IsZero() || n.Sub(s1).Seconds() >= float64(utils.MyConfig.Interval.SaveS) {
		data.InteractionStats.Last1CommentNum = o.InteractionStats.CommentNum
	}
	if s2.IsZero() || n.Sub(s2).Seconds() >= float64(utils.MyConfig.Interval.SaveSEX) {
		data.InteractionStats.Last1CommentNum = o.InteractionStats.CommentNum
	}
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
	if err = ioutil.WriteFile(path + "/screen_room_overview.tmp", b, 0755); err != nil {
		return
	}
	return
}


func ScreenSaveBaseInfo(roomId string, data ScreenBaseInfoRespData) (err error) {
	baseInfoM.Lock()
	defer baseInfoM.Unlock()
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
	productDetailM.Lock()
	defer productDetailM.Unlock()
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
	_ = json.Unmarshal(b, &data)
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
	_ = json.Unmarshal(b, &data)
	return
}
func ScreenLoadRoomOverview(roomId string) (data ScreenRoomOverviewRespData, err error) {
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path + "/screen_room_overview.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(path + "/screen_room_overview.tmp")
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &data)
	return
}
func ScreenLoadRoomDataTrendTP(roomId string) (data ScreenRoomDataTrendTPRespData, err error) {
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path + "/screen_data_data_trend-tp.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(path + "/screen_data_data_trend-tp.tmp")
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &data)
	return
}
func ScreenLoadLiveRoomDashboardV2Resp(roomId string) (data LiveRoomDashboardV2RespData, err error) {
	path := fmt.Sprintf(utils.RoomsDataPath, roomId)
	if _, _err := os.Stat(path + "/live_room_dashboard_v2.tmp"); os.IsNotExist(_err) {
		return
	}
	b, err := ioutil.ReadFile(path + "/live_room_dashboard_v2.tmp")
	if err != nil {
		return
	}
	// 反序列化
	_ = json.Unmarshal(b, &data)
	return
}
