package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Import struct {
	Domain string `json:"domain"`
	ExpirationDate float64 `json:"expirationDate,omitempty"`
	HostOnly bool `json:"hostOnly"`
	HTTPOnly bool `json:"httpOnly"`
	Name string `json:"name"`
	Path string `json:"path"`
	SameSite string `json:"sameSite"`
	Secure bool `json:"secure"`
	Session bool `json:"session"`
	StoreID string `json:"storeId"`
	Value string `json:"value"`
	ID int `json:"id"`
}
type Export struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Domain string `json:"domain"`
	Path string `json:"path"`
	Priority string `json:"priority"`
	Secure bool `json:"secure,omitempty"`
	HTTPOnly bool `json:"httpOnly,omitempty"`
	SameSite string `json:"sameSite,omitempty"`
}

func main() {
	str := "ttcid=6d4d27f9e8b043a89989d20c0ef1e8db66; s_v_web_id=verify_kq6nl3ty_K1xTg0la_M3R2_41wa_ADQ8_gtv3rFHjuwxu; Hm_lvt_45173f3eae0174bc5b02a4973fe5a872=1624282034,1624282074; passport_csrf_token_default=1f5275d5de17a057b68c9807ad863f17; n_mh=pM_M_W6plOo31O857LuNdvErIVF3raWRr2sN34yZCJU; passport_auth_status=37c16772f65d2cb975d6194b6b6ff119%2C; passport_auth_status_ss=37c16772f65d2cb975d6194b6b6ff119%2C; sid_guard=a58ebce121745500238aa61879313f04%7C1624282075%7C5184000%7CFri%2C+20-Aug-2021+13%3A27%3A55+GMT; uid_tt=2fc23a82fd730760a781de45e672d445; uid_tt_ss=2fc23a82fd730760a781de45e672d445; sid_tt=a58ebce121745500238aa61879313f04; sessionid=a58ebce121745500238aa61879313f04; sessionid_ss=a58ebce121745500238aa61879313f04; passport_csrf_token=1f5275d5de17a057b68c9807ad863f17; PHPSESSID=86065f2725d27fa555e0ad28d5de8087; PHPSESSID_SS=86065f2725d27fa555e0ad28d5de8087; LUOPAN_DT=session_6976234919013990664; gfsitesid=YTU4ZWJjZTEyMXwxNjI0MjgyMDc1OTN8fDM2MDY4NTE5MTA5MDE0ODYHBwcHBwcH; gftoken=YTU4ZWJjZTEyMXwxNjI0MjgyMDc1OTN8fDAGBgYGBgY; csrftoken=I5heplOarqYqyS8LrrssuWmI; tt_scid=KaTu9xzFgIxl1FWD3docR5qfXCUAQx3uA5XTFNlxZpTTTKRhqUL8Ex5K6vbUJBnF362f; Hm_lpvt_45173f3eae0174bc5b02a4973fe5a872=1624291209; MONITOR_WEB_ID=f0132c30-bf06-4362-9e3f-41b549607bb3"
	arr := strings.Split(str, ";")
	cookies := []Export{}
	for _, a := range arr {
		as := strings.Split(a, "=")
		if len(as) >= 2 {
			cookies = append(cookies, Export{
				Name:     strings.TrimSpace(as[0]),
				Value:    strings.TrimSpace(as[1]),
				Domain:   ".jinritemai.com",
				Path:     "/",
				Priority: "Medium",
				Secure:   false,
				HTTPOnly: true,
				SameSite: "",
			})
		}
	}


	//im, err := ReadAll("./import.tmp")
	//if err != nil {
	//	panic(err)
	//}
	//imports := []Import{}
	//json.Unmarshal(im, &imports)
	//cookies := []Export{}
	//for _, i := range imports {
	//	cookies = append(cookies, Export{
	//		Name:         i.Name,
	//		Value:        i.Value,
	//		Domain:       i.Domain,
	//		Path:         i.Path,
	//		Expires: int(i.ExpirationDate),
	//		Size:         0,
	//		HTTPOnly:     i.HTTPOnly,
	//		Secure:       i.Secure,
	//		Session:      i.Session,
	//		SameSite:     i.SameSite,
	//		Priority:     "Medium",
	//		SameParty:    false,
	//		SourceScheme: "",
	//		SourcePort:   0,
	//	})
	//}
	//exports := map[string][]Export{
	//	"cookies": cookies,
	//}
	// 写入
	f, err := os.Create("cookies.tmp")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, _ := json.Marshal(cookies)
	f.Write(b)
	fmt.Println("done")
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}