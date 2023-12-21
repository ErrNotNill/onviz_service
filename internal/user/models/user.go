package models

type UserData struct {
	ID       int    `json:"id omitempty"`
	Username string `json:"username omitempty"`
	Email    string `json:"email omitempty"`
	Password string `json:"password omitempty"`
	Token    string `json:"token omitempty"`
	Country  string `json:"country omitempty"`
}

type UserDataOnviz struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Country      string `json:"country"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
