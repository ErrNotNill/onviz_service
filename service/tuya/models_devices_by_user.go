package tuya

type StatusItem struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}

type ResultItem struct {
	Sub         bool         `json:"sub"`
	CreateTime  int64        `json:"create_time"`
	LocalKey    string       `json:"local_key"`
	OwnerID     string       `json:"owner_id"`
	IP          string       `json:"ip"`
	BizType     int          `json:"biz_type"`
	Icon        string       `json:"icon"`
	TimeZone    string       `json:"time_zone"`
	UUID        string       `json:"uuid"`
	ProductName string       `json:"product_name"`
	ActiveTime  int64        `json:"active_time"`
	UID         string       `json:"uid"`
	UpdateTime  int64        `json:"update_time"`
	ProductID   string       `json:"product_id"`
	Name        string       `json:"name"`
	Online      bool         `json:"online"`
	ID          string       `json:"id"`
	Category    string       `json:"category"`
	Status      []StatusItem `json:"status"`
}

type ResponseDevices struct {
	Result  []ResultItem `json:"result"`
	T       int64        `json:"t"`
	Success bool         `json:"success"`
}
