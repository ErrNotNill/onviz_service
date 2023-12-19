package models

import "golang.org/x/oauth2"

var OauthConfig = &oauth2.Config{
	ClientID:     "your-client-id",
	ClientSecret: "your-client-secret",
	RedirectURL:  "your-redirect-url",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "authorization-url",
		TokenURL: "token-url",
	},
	Scopes: []string{"scope"},
}
