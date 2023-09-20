package bot_bitrix

import (
	"encoding/json"
	"fmt"
	"io"
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

func CallbackHandlerOld(w http.ResponseWriter, r *http.Request) {
	// Extract the authorization code from the query parameters
	code := r.URL.Query().Get("code")

	//just checks git
	// Step 3: Exchange the authorization code for the access token
	tokenURL := fmt.Sprintf("https://onviz.bitrix24.com/oauth/token"+
		"?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		clientID, clientSecret, code, redirectURI)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}
	fmt.Println(r.Body)
	defer resp.Body.Close()

	var newToken string
	err = json.NewDecoder(resp.Body).Decode(&newToken)
	fmt.Println(newToken)

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

	fmt.Println("access token: ", accessToken)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	//client_id := `FM9Vb9sOseDFnAx4BNOvgbpr7r37dBmL`
	//client_secret := `mM20SgbV7mWjlSLvQ1IR8UpMsFSlUXtl`

	//url := fmt.Sprintf("https://api.cdek.ru/v2/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", client_id, client_secret)
	tokenURL := fmt.Sprintf(`https://onviz.bitrix24.ru/oauth/authorize/?
	client_id=%s`,
		clientID)

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

func NextStepAuth(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	//client_id := `FM9Vb9sOseDFnAx4BNOvgbpr7r37dBmL`
	//client_secret := `mM20SgbV7mWjlSLvQ1IR8UpMsFSlUXtl`
	code := r.URL.Query().Get("")
	//url := fmt.Sprintf("https://api.cdek.ru/v2/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", client_id, client_secret)
	tokenURL := fmt.Sprintf("https://onviz.bitrix24.com/oauth/token"+
		"?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s",
		clientID, clientSecret, code)

	post, err := client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}
	body, err := io.ReadAll(post.Body)
	fmt.Println("post.Body", post.Body)
	json.Unmarshal(body, &post.Body)
	fmt.Fprint(w, string(body))

	/*resp := &Response{}
	//fmt.Println("post.Body", post.Body)
	json.Unmarshal(body, &resp)

	fmt.Println(post.StatusCode)
	fmt.Println("access_token", resp.AccessToken)
	fmt.Println("token_type", resp.TokenType)
	fmt.Println("expires_in", resp.ExpiresIn)
	fmt.Println("scope", resp.Scope)
	fmt.Println("jti", resp.Jti)

	//order := &Order{}
	var bearer = "Bearer " + resp.AccessToken
	req, err := http.NewRequest("GET", `https://api.cdek.ru/v2/orders?cdek_number=1460493586`, nil)
	req.Header.Add("Authorization", bearer)
	newclient := &http.Client{}
	rez, err := newclient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer rez.Body.Close()

	newbody, err := io.ReadAll(rez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))

	newReq, err := http.NewRequest("GET", `https://api.cdek.ru/v2/registries?date=2023-09-12`, nil)
	newReq.Header.Add("Authorization", bearer)
	newPay := &http.Client{}
	newRez, err := newPay.Do(newReq)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer newRez.Body.Close()

	newBody, err := io.ReadAll(newRez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newBody)))*/

	/*get, err := client.Get(`https://api.cdek.ru/v2/orders?cdek_number=1463958253`)
	get.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	get.Header.Add("Authorization", bearer)
	//fmt.Println("bearer::: ", bearer)

	body, err = io.ReadAll(get.Body)
	fmt.Println("StatusCode: ", get.StatusCode)
	json.Unmarshal(body, &order)
	fmt.Println("order", order)
	fmt.Println("get.Body", get.Body)*/

	//client.Get(`https://api.cdek.ru/v2/payment`)

	//req, err := http.NewRequest("POST", `https://api.cdek.ru/v2/oauth/token?grant_type=client_credentials&client_id=FM9Vb9sOseDFnAx4BNOvgbpr7r37dBmL&client_secret=mM20SgbV7mWjlSLvQ1IR8UpMsFSlUXtl`, nil)
	//req.Header.Set("Content-Type", "x-www-form-urlencoded")
	//fmt.Println("req.Body", req.Body)
}
