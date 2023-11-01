package models

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	SessionId int64  `json:"session"`
	ChannelId int64  `json:"channel_id"`
}
