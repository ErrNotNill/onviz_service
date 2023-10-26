package bot_bitrix

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		"?client_id=%s&redirect_uri=%s&response_type=code", localBitrixAppID, redirectURI)
	http.Redirect(w, r, authURL, http.StatusFound)
}

/*func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	tokenURL := fmt.Sprintf(`https://onviz.bitrix24.ru/oauth/authorize/?
	client_id=%s`,
		BitrixClientId)

	post, err := client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}
	body, err := io.ReadAll(post.Body)
	fmt.Println("post.Body", post.Body)
	json.Unmarshal(body, &post.Body)
	fmt.Fprint(w, string(body))
}*/

func NextStepAuth(w http.ResponseWriter, r *http.Request) {
	BitrixClientId := os.Getenv("BITRIX_CLIENT_ID")
	BitrixClientSecret := os.Getenv("BITRIX_CLIENT_SECRET")
	client := &http.Client{}
	code := r.URL.Query().Get("")
	//url := fmt.Sprintf("https://api.cdek.ru/v2/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", client_id, client_secret)
	tokenURL := fmt.Sprintf("https://onviz.bitrix24.com/oauth/token"+
		"?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s",
		BitrixClientId, BitrixClientSecret, code)

	post, err := client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}
	body, err := io.ReadAll(post.Body)
	fmt.Println("post.Body", post.Body)
	json.Unmarshal(body, &post.Body)
	fmt.Fprint(w, string(body))
}
