package service

import (
	"net/url"
)

func GenerateAuthorizationURL(clientID, redirectURI string) string {
	authEndpoint := "https://openapi.tuyaeu.com/v1.0/authorize"
	queryParams := url.Values{}
	queryParams.Set("client_id", clientID)
	queryParams.Set("response_type", "code")
	queryParams.Set("redirect_uri", redirectURI)
	authURL := authEndpoint + "?" + queryParams.Encode()
	return authURL
}
