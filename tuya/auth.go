package tuya

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	clientID := ClientID
	redirectURI := YaRedirectUri

	reader := r.Body
	var yaparams = YandexAuthParams{}

	js, err := json.MarshalIndent(reader, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(js, &yaparams)
	w.Write(js)
	w.Write([]byte(yaparams.State))

	state := yaparams.State
	http.Redirect(w, r, fmt.Sprintf("https://social.yandex.net/broker/redirect?state=%s", state), http.StatusTemporaryRedirect)

	authURL := GenerateAuthorizationURL(clientID, redirectURI)
	response := struct {
		AuthURL string `json:"auth_url"`
	}{
		AuthURL: authURL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GenerateAuthorizationURL(clientID, redirectURI string) string {
	authEndpoint := "https://openapi.tuyaeu.com/v1.0/authorize"
	queryParams := url.Values{}
	queryParams.Set("client_id", clientID)
	queryParams.Set("response_type", "code")
	queryParams.Set("redirect_uri", redirectURI)
	authURL := authEndpoint + "?" + queryParams.Encode()
	return authURL
}
