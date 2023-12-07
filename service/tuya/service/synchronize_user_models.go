package service

type SynchronizeResult struct {
	Result  SynchronizeResultData `json:"result"`
	Success bool                  `json:"success"`
	T       int64                 `json:"t"`
	TID     string                `json:"tid"`
}

type SynchronizeResultData struct {
	UID string `json:"uid"`
}
