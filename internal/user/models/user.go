package models

type UserData struct {
	ID       int    `json:"id omitempty"`
	Username string `json:"username omitempty"`
	Email    string `json:"email omitempty"`
	Password string `json:"password omitempty"`
	Token    string `json:"token omitempty"`
	Country  string `json:"country omitempty"`
}
