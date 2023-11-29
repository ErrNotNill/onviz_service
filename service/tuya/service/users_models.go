package service

type TuyaUsers struct {
	Result struct {
		HasMore bool `json:"has_more"`
		List    []struct {
			CreateTime int    `json:"create_time"`
			Email      string `json:"email"`
			Mobile     string `json:"mobile"`
			UID        string `json:"uid"`
			UpdateTime int    `json:"update_time"`
			Username   string `json:"username"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"result"`
	Success bool   `json:"success"`
	T       int64  `json:"t"`
	Tid     string `json:"tid"`
}
