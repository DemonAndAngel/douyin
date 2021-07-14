package main

import (
	"douyin/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strconv"
	"time"

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
	Name         string              `json:"name"`                   // Cookie name.
	Value        string              `json:"value"`                  // Cookie value.
	URL          string              `json:"url,omitempty"`          // The request-URI to associate with the setting of the cookie. This value can affect the default domain, path, source port, and source scheme values of the created cookie.
	Domain       string              `json:"domain,omitempty"`       // Cookie domain.
	Path         string              `json:"path,omitempty"`         // Cookie path.
	Secure       bool                `json:"secure,omitempty"`       // True if cookie is secure.
	HTTPOnly     bool                `json:"httpOnly,omitempty"`     // True if cookie is http-only.
	SameSite     string      `json:"sameSite,omitempty"`     // Cookie SameSite type.
	Expires      time.Time `json:"expires,omitempty"`      // Cookie expiration date, session cookie if not set
	Priority     string      `json:"priority,omitempty"`     // Cookie Priority.
	SameParty    bool                `json:"sameParty,omitempty"`    // True if cookie is SameParty.
	SourceScheme string  `json:"sourceScheme,omitempty"` // Cookie source scheme type.
	SourcePort   int64               `json:"sourcePort,omitempty"`
}

func init() {
	os.Setenv("FYNE_FONT", utils.FolderPath + "/simhei.ttf")
	// 加载配置
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(utils.FolderPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}


func main() {
	a := app.New()
	w := a.NewWindow("软件")
	grabS := widget.NewEntry()
	grabS.SetText(strconv.FormatInt(viper.GetInt64("Interval.GrabS"), 10))
	form := widget.NewForm(widget.NewFormItem("抓取时间间隔: ", grabS))
	form.SubmitText = "应用配置"
	form.OnSubmit = func () {
		if grabS.Text != "" {

		}
	}
	var run *widget.Button
	run = widget.NewButton("启动", func() {
		run.Text = "停止"
	})

	w.SetContent(container.NewVBox(
		form,
		run,
	))
	w.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})
	w.ShowAndRun()

	//str := "x-jupiter-uuid=16247011957686924; passport_csrf_token_default=37d08ed88843f6d6285663738f0ecf3b; passport_csrf_token=37d08ed88843f6d6285663738f0ecf3b; ttcid=76f01a98fdef4d6f9a8f69715a55d54017; MONITOR_WEB_ID=bd72b778-b447-4edd-8828-65e0ce9dc3b3; business-account-center-csrf-secret=df3146ee1efa5157259af14edcbec7e2; business-account-center-csrf-token=VA3OVVkb-IsX8ZlcdIaVBiITRZtP_7bOKGzM; x-jupiter-uuid=1624701208644482; s_v_web_id=verify_kqdl5guf_lgR7GHnn_91u6_4aK1_8XEk_O65p0DfNedRZ; n_mh=pM_M_W6plOo31O857LuNdvErIVF3raWRr2sN34yZCJU; sso_uid_tt=466bd8e433645dfb5517faf115f5bfa4; sso_uid_tt_ss=466bd8e433645dfb5517faf115f5bfa4; toutiao_sso_user=2131cd5b12f6370ef58c35ee0541320c; toutiao_sso_user_ss=2131cd5b12f6370ef58c35ee0541320c; ttwid=1%7CuKBGRch7YuNJMl60NYxrA3Dmk6zg3XRrBZ1JYArWsak%7C1624708656%7C2122eb2b57df8b5195467a8644a1ba2c432091697fa40ae36a40211b50952b1f; odin_tt=52fe10847b17d5d4470cf92993d5859d973030ebd76e0bd29c184c80b009ff6a584b65dc68125d8edb848e689fc4f0cfa876fd878de727ea487bfc3163775eec; passport_auth_status_ss=109ec334950081846c86bf8f89b63320%2Ca11a9c9124b7437cda4304b372f373e2; passport_auth_status=109ec334950081846c86bf8f89b63320%2Ca11a9c9124b7437cda4304b372f373e2; ucas_sso_c0_ss=CkEKBTEuMC4wEI2IjdamxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUvg_lK7Dtekj2nvmZBAPn-8BDeFw; ucas_sso_c0=CkEKBTEuMC4wEI2IjdamxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUvg_lK7Dtekj2nvmZBAPn-8BDeFw; ucas_c0_ss_buyin=CkEKBTEuMC4wEI2Ij9amxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUeWwFnWMhb3EtHpiLGWKqWErNpXk; sid_guard=1773acd671660166e570ae5eac30de2a%7C1624708660%7C5184000%7CWed%2C+25-Aug-2021+11%3A57%3A40+GMT; uid_tt=85e4fd69b205ca1f55b7df138724eba6; uid_tt_ss=85e4fd69b205ca1f55b7df138724eba6; sid_tt=1773acd671660166e570ae5eac30de2a; sessionid=1773acd671660166e570ae5eac30de2a; sessionid_ss=1773acd671660166e570ae5eac30de2a; ucas_c0_buyin=CkEKBTEuMC4wEI2Ij9amxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUeWwFnWMhb3EtHpiLGWKqWErNpXk; gftoken=MTc3M2FjZDY3MXwxNjI0NzA4Njc5ODZ8fDAGBgYGBgY; tt_scid=XygQ5hyALTmzQFi-PpcGsyCeadrpcnGDO9u4MMJOA9qUe1gtcIAZN2Irn2S3Z8.69fb0; buyin_shop_type=24;passport_csrf_token_default=37d08ed88843f6d6285663738f0ecf3b; passport_csrf_token=37d08ed88843f6d6285663738f0ecf3b; ttcid=76f01a98fdef4d6f9a8f69715a55d54017; MONITOR_WEB_ID=bd72b778-b447-4edd-8828-65e0ce9dc3b3; business-account-center-csrf-secret=df3146ee1efa5157259af14edcbec7e2; business-account-center-csrf-token=VA3OVVkb-IsX8ZlcdIaVBiITRZtP_7bOKGzM; x-jupiter-uuid=1624701208644482; s_v_web_id=verify_kqdl5guf_lgR7GHnn_91u6_4aK1_8XEk_O65p0DfNedRZ; n_mh=pM_M_W6plOo31O857LuNdvErIVF3raWRr2sN34yZCJU; sso_uid_tt=466bd8e433645dfb5517faf115f5bfa4; sso_uid_tt_ss=466bd8e433645dfb5517faf115f5bfa4; toutiao_sso_user=2131cd5b12f6370ef58c35ee0541320c; toutiao_sso_user_ss=2131cd5b12f6370ef58c35ee0541320c; buyin_shop_type=24; buyin_app_id=1128; ttwid=1%7CuKBGRch7YuNJMl60NYxrA3Dmk6zg3XRrBZ1JYArWsak%7C1624708656%7C2122eb2b57df8b5195467a8644a1ba2c432091697fa40ae36a40211b50952b1f; odin_tt=52fe10847b17d5d4470cf92993d5859d973030ebd76e0bd29c184c80b009ff6a584b65dc68125d8edb848e689fc4f0cfa876fd878de727ea487bfc3163775eec; passport_auth_status_ss=109ec334950081846c86bf8f89b63320%2Ca11a9c9124b7437cda4304b372f373e2; passport_auth_status=109ec334950081846c86bf8f89b63320%2Ca11a9c9124b7437cda4304b372f373e2; ucas_sso_c0_ss=CkEKBTEuMC4wEI2IjdamxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUvg_lK7Dtekj2nvmZBAPn-8BDeFw; ucas_sso_c0=CkEKBTEuMC4wEI2IjdamxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUvg_lK7Dtekj2nvmZBAPn-8BDeFw; ucas_c0_ss_buyin=CkEKBTEuMC4wEI2Ij9amxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUeWwFnWMhb3EtHpiLGWKqWErNpXk; sid_guard=1773acd671660166e570ae5eac30de2a%7C1624708660%7C5184000%7CWed%2C+25-Aug-2021+11%3A57%3A40+GMT; uid_tt=85e4fd69b205ca1f55b7df138724eba6; uid_tt_ss=85e4fd69b205ca1f55b7df138724eba6; sid_tt=1773acd671660166e570ae5eac30de2a; sessionid=1773acd671660166e570ae5eac30de2a; sessionid_ss=1773acd671660166e570ae5eac30de2a; ucas_c0_buyin=CkEKBTEuMC4wEI2Ij9amxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUeWwFnWMhb3EtHpiLGWKqWErNpXk; SASID=SID2_3489035282219721424; BUYIN_SASID=SID2_3489035282219721424; gftoken=MTc3M2FjZDY3MXwxNjI0NzA4Njc5ODZ8fDAGBgYGBgY; tt_scid=gh34zgfhBY-spFNcBTHZvk61UTV4LEvfj1puNP96SY6iSLDQH4wX7p2wOEw.5fdOb8ed;gfsitesid=MTc3M2FjZDY3MXwxNjI0NzExOTIzMzR8fDM2MDY4NTE5MTA5MDE0ODYHBwcHBwcH; ttcid=74ee64c849774ed9a01279d2e8c45c3916; passport_csrf_token_default=37d08ed88843f6d6285663738f0ecf3b; passport_csrf_token=37d08ed88843f6d6285663738f0ecf3b; n_mh=pM_M_W6plOo31O857LuNdvErIVF3raWRr2sN34yZCJU; sso_uid_tt=466bd8e433645dfb5517faf115f5bfa4; sso_uid_tt_ss=466bd8e433645dfb5517faf115f5bfa4; toutiao_sso_user=2131cd5b12f6370ef58c35ee0541320c; toutiao_sso_user_ss=2131cd5b12f6370ef58c35ee0541320c; csrftoken=7gE1YqqzSYvxOSaBXc-XSZvQ; gfsitesid=OWEzN2I5YTM1N3wxNjI0NzAzMTI1Mjl8fDM2MDY4NTE5MTA5MDE0ODYHBwcHBwcH; Hm_lvt_45173f3eae0174bc5b02a4973fe5a872=1624288917,1624288950,1624455214,1624703126; s_v_web_id=verify_kqdmajws_HNOs6BTY_yh5L_4niW_8GeV_KWbDnRDytBZg; Hm_lpvt_45173f3eae0174bc5b02a4973fe5a872=1624703128; MONITOR_WEB_ID=143fc66c-b1b7-4ba9-8982-51118b1c8d1e; ttwid=1%7CuKBGRch7YuNJMl60NYxrA3Dmk6zg3XRrBZ1JYArWsak%7C1624708656%7C2122eb2b57df8b5195467a8644a1ba2c432091697fa40ae36a40211b50952b1f; odin_tt=52fe10847b17d5d4470cf92993d5859d973030ebd76e0bd29c184c80b009ff6a584b65dc68125d8edb848e689fc4f0cfa876fd878de727ea487bfc3163775eec; passport_auth_status_ss=109ec334950081846c86bf8f89b63320%2Ca11a9c9124b7437cda4304b372f373e2; passport_auth_status=109ec334950081846c86bf8f89b63320%2Ca11a9c9124b7437cda4304b372f373e2; ucas_sso_c0_ss=CkEKBTEuMC4wEI2IjdamxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUvg_lK7Dtekj2nvmZBAPn-8BDeFw; ucas_sso_c0=CkEKBTEuMC4wEI2IjdamxMXrYBi9LyC4ovCKt43oByiPETDuvaC3mo20BkC0rNyGBki04JiJBlCHvMjirMqw22BYfhIUvg_lK7Dtekj2nvmZBAPn-8BDeFw; sid_guard=1773acd671660166e570ae5eac30de2a%7C1624708660%7C5184000%7CWed%2C+25-Aug-2021+11%3A57%3A40+GMT; uid_tt=85e4fd69b205ca1f55b7df138724eba6; uid_tt_ss=85e4fd69b205ca1f55b7df138724eba6; sid_tt=1773acd671660166e570ae5eac30de2a; sessionid=1773acd671660166e570ae5eac30de2a; sessionid_ss=1773acd671660166e570ae5eac30de2a; BUYIN_SASID=SID2_3489035282219721424; LUOPAN_DT=session_6978098066997969183; gftoken=MTc3M2FjZDY3MXwxNjI0NzE4MjY5Mzd8fDAGBgYGBgY; tt_scid=rI6HbgK0bpbVJJLzNPAwwbelzycMkvOKd6ObJF4nrClYiky3xirSd1cL.CRt6cD-4783"
	//arr := strings.Split(str, ";")
	//cookies := []network.CookieParam{}
	//for _, a := range arr {
	//	as := strings.Split(a, "=")
	//	if len(as) >= 2 {
	//		cookies = append(cookies, network.CookieParam{
	//			Name:     strings.TrimSpace(as[0]),
	//			Value:    strings.TrimSpace(as[1]),
	//			Domain:   ".jinritemai.com",
	//			Path:     "/",
	//			Priority: "Medium",
	//			Secure:   false,
	//			HTTPOnly: true,
	//			SameSite: "",
	//		})
	//	}
	//}


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
	//		HTTPOnly:     i.HTTPOnly,
	//		Secure:       i.Secure,
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
	//f, err := os.Create("cookies.tmp")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//b, _ := json.Marshal(cookies)
	//f.Write(b)
	//fmt.Println("done")
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}