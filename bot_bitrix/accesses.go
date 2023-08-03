package bot_bitrix

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	redirectURI = "https://onviz-api.ru/redir"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("OK")))
	fmt.Println(r.Body)
}

func GetAccess(w http.ResponseWriter, r *http.Request) {

	// Step 1: Redirect the user to the Bitrix24 authorization page
	authURL := fmt.Sprintf("https://onviz.bitrix24.ru/oauth/authorize"+
		"?client_id=%s&redirect_uri=%s&response_type=code", clientID, redirectURI)
	http.Redirect(w, r, authURL, http.StatusFound)
}

// Step 2: Handle the callback from Bitrix24
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the authorization code from the query parameters
	code := r.URL.Query().Get("code")

	// Step 3: Exchange the authorization code for the access token
	tokenURL := fmt.Sprintf("https://onviz.bitrix24.com/oauth/token"+
		"?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		clientID, clientSecret, code, redirectURI)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response to get the access token
	var tokenResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		fmt.Println("Failed to parse token response:", err)
		return
	}

	accessToken := tokenResponse["access_token"].(string)

	// Step 4: Use the access token to make API requests on behalf of the user
	// Your code to make API requests goes here...

	fmt.Println(accessToken)
}
