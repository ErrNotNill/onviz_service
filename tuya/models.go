package tuya

const (
	Host     = "https://openapi.tuyaeu.com"
	ClientID = "9x8wfym7m5vyck7tdwwt"
	Secret   = "d8205ed66f15471fa969aecab48ab495"
	DeviceID = "bf85de23e4cf1c10fb6bsn" //example
)

var (
	Token           string
	RefreshTokenVal string
	ResultPolicy    string
)

type TokenResponse struct {
	Result struct {
		AccessToken  string `json:"access_token"`
		ExpireTime   int    `json:"expire_time"`
		RefreshToken string `json:"refresh_token"`
		UID          string `json:"uid"`
	} `json:"result"`
	Success bool  `json:"success"`
	T       int64 `json:"t"`
}

type DeviceModel struct {
	UUID   string `json:"uuid"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Sub    bool   `json:"sub"`
	Model  string `json:"model"`
	Status []struct {
		Code  string      `json:"code"`
		Value interface{} `json:"value"`
	} `json:"status"`
	Category    string `json:"category"`
	Online      bool   `json:"online"`
	ID          string `json:"id"`
	TimeZone    string `json:"time_zone"`
	LocalKey    string `json:"local_key"`
	UpdateTime  int    `json:"update_time"`
	ActiveTime  int    `json:"active_time"`
	OwnerID     string `json:"owner_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
}

type GetDeviceResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Result  DeviceModel `json:"result"`
	T       int64       `json:"t"`
}

type PostDeviceCmdResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Result  bool   `json:"result"`
	T       int64  `json:"t"`
}

type Device struct {
	Result  DeviceInfo `json:"result"`
	Success bool       `json:"success"`
	T       int64      `json:"t"`
	TID     string     `json:"tid"`
}

type DeviceInfo struct {
	ActiveTime  int64        `json:"active_time"`
	BizType     int64        `json:"biz_type"`
	Category    string       `json:"category"`
	CreateTime  int64        `json:"create_time"`
	Icon        string       `json:"icon"`
	ID          string       `json:"id"`
	IP          string       `json:"ip"`
	Lat         string       `json:"lat"`
	LocalKey    string       `json:"local_key"`
	Lon         string       `json:"lon"`
	Model       string       `json:"model"`
	Name        string       `json:"name"`
	NodeID      string       `json:"node_id"`
	Online      bool         `json:"online"`
	OwnerID     string       `json:"owner_id"`
	ProductID   string       `json:"product_id"`
	ProductName string       `json:"product_name"`
	Status      []StatusInfo `json:"status"`
	Sub         bool         `json:"sub"`
	TimeZone    string       `json:"time_zone"`
	UID         string       `json:"uid"`
	UpdateTime  int64        `json:"update_time"`
	UUID        string       `json:"uuid"`
}

type StatusInfo struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}
