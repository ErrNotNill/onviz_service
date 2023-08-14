package tuya

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	if clientID == "" || redirectURI == "" {
		http.Error(w, "Missing client_id or redirect_uri", http.StatusBadRequest)
		return
	}
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
