package service

type Result struct {
	LastRowKey string `json:"last_row_key"`
	List       []Item `json:"list"`
	Total      int    `json:"total"`
	HasMore    bool   `json:"has_more"`
}

type Item struct {
	ID           string `json:"id"`
	GatewayID    string `json:"gateway_id"`
	NodeID       string `json:"node_id"`
	UUID         string `json:"uuid"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	Name         string `json:"name"`
	ProductID    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	LocalKey     string `json:"local_key"`
	Sub          bool   `json:"sub"`
	AssetID      string `json:"asset_id"`
	OwnerID      string `json:"owner_id"`
	IP           string `json:"ip"`
	Lon          string `json:"lon"`
	Lat          string `json:"lat"`
	Model        string `json:"model"`
	TimeZone     string `json:"time_zone"`
	ActiveTime   int64  `json:"active_time"`
	UpdateTime   int64  `json:"update_time"`
	CreateTime   int64  `json:"create_time"`
	Online       bool   `json:"online"`
	Icon         string `json:"icon"`
}

type Response struct {
	Result  Result `json:"result"`
	T       int64  `json:"t"`
	Success bool   `json:"success"`
}
