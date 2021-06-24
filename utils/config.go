package utils

var MyConfig *Config

type Config struct {
	Interval ConfigInterval `json:"interval"`
	Server   ConfigServer   `json:"server"`
	System   ConfigSystem   `json:"system"`
}
type ConfigInterval struct {
	GrabS         int64 `json:"grab_s"`
	SaveS         int64 `json:"save_s"`
	CheckLoginS   int64 `json:"check_login_s"`
	QrcodeExpireS int64 `json:"qrcode_expire_s"`
	UrlS          int64 `json:"url_s"`
}
type ConfigServer struct {
	Port int `json:"port"`
}
type ConfigSystem struct {
	UserAgent string `json:"user_agent"`
}