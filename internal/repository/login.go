package repository

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"net/url"
	"onviz/DB"
	"onviz/internal/user/models"
	"os"
	"strings"
)

func CreateAccount(name, email, password string) error {
	fmt.Println("db connected")
	tokenData := email + password // Customize this according to your needs
	sha256Hash := sha256.Sum256([]byte(tokenData))
	token := hex.EncodeToString(sha256Hash[:])
	result, err := DB.Db.Exec(`insert into Users (Name, Email, Password, Token) values (?,?,?,?)`,
		name, email, password, token)

	if err != nil {
		fmt.Println("cant insert data to dbase")
		//panic(err)
		return err
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
	return nil
}

func GetAccount(email, password string) models.UserData {

	rows, err := DB.Db.Query(`SELECT Email, Password FROM Users WHERE Email = ? AND Password = ?`, email, password)
	if err != nil {
		fmt.Println("cant get data from dbase:", err)
		return models.UserData{}
	}
	defer rows.Close()

	p := models.UserData{}

	for rows.Next() {
		err := rows.Scan(&p.Email, &p.Password)
		if err != nil {
			fmt.Println("Error scanning data:", err)
			return models.UserData{}
		}
	}
	fmt.Println("Email is not nil>>>: ", p.Email)
	if p.Email == "" {
		fmt.Println("Email is>>>: ", p.Email)
		return models.UserData{}
	}
	return p
	//fmt.Println("products: ", p)

}

func AuthPage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var userData models.UserData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	tokenData := userData.Email + userData.Password // Customize this according to your needs
	sha256Hash := sha256.Sum256([]byte(tokenData))
	token := hex.EncodeToString(sha256Hash[:])
	// Process user registration data (userData) as needed
	fmt.Printf("Received registration data: %+v\n", userData)
	// You can now handle the registration logic, such as storing the data in a database

	// Handle registration error
	if err := CreateAccount(userData.Username, userData.Email, userData.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{"message": "User exist"})
		if err != nil {
			log.Println("Error encode account", err.Error())
			return
		}
		fmt.Println("Error creating, user exists")
		return
	}

	// Send a response back to the client
	response := map[string]string{"token": token}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set the response status once, indicating a successful response
	w.Write(responseJSON)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the callback after the user authorizes your application
	code := r.URL.Query().Get("code")

	// Exchange the authorization code for an access token
	token, err := models.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging code for token: %v", err)
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		return
	}

	// You can now use the access token to make authenticated requests to the OAuth provider's API
	fmt.Printf("Access Token: %v\n", token.AccessToken)

	// Your existing code...
}

func GenerateYandexAuthURL(clientID, redirectURI, scope, state string) string {
	baseURL := "https://oauth.yandex.com/authorize"
	authURL, _ := url.Parse(baseURL)

	queryParams := url.Values{}
	queryParams.Add("client_id", clientID)
	queryParams.Add("redirect_uri", redirectURI)
	queryParams.Add("response_type", "code")
	queryParams.Add("scope", scope)
	queryParams.Add("state", state)

	authURL.RawQuery = queryParams.Encode()

	return authURL.String()
}

var UserFromTuya string

func generateRandomCode() (string, error) {
	length := 32
	byteSize := length / 4
	randomBytes := make([]byte, byteSize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	code := base64.URLEncoding.EncodeToString(randomBytes)
	code = code[:len(code)-1]
	return code, nil
}

func SplitString(path string) string {
	newSpl := strings.Split(path, "https://social.yandex.ru/broker2/authz_in_web/")
	var FinalStr string
	for _, s := range newSpl {
		nextSpl := strings.Split(s, "/callback")
		FinalStr = nextSpl[0]
	}
	return FinalStr
}

func extractCallbackID(path string) string {
	// Find the position of /web/ in the path
	strings.Cut(path, "https://social.yandex.ru/broker2/authz_in_web/")
	_, newStr, _ := strings.Cut(path, "/callback")
	return newStr
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Client struct {
	ClientId     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Code         string `json:"code"`
}

type TokenReq struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

var TransferData interface{}

func ReturnTokenToYandex(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(TransferData)
	if err != nil {
		panic(err)
	}
	w.Write(b)
	fmt.Println("string(b):>> ", string(b))
}

func TokenOauthWithCode(w http.ResponseWriter, r *http.Request) {

	nextcode := r.URL.Query().Get("code")
	fmt.Println("nextCODE:???", nextcode)

	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")

	// Encode client ID and client secret for Basic Authentication
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))

	// Prepare the token request data
	tokenEndpoint := "https://oauth.yandex.ru/token"
	tokenData := url.Values{}
	tokenData.Set("grant_type", "authorization_code")
	tokenData.Set("code", nextcode)

	// Make the POST request
	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBufferString(tokenData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authHeader)

	// Make the request
	clientd := &http.Client{}
	resp, err := clientd.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body:", err.Error())
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	TransferData = body
	var token TokenReq
	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return
	}
	err = json.Unmarshal(body, &TransferData)
	if err != nil {
		log.Println("Error unmarshalling transfer_data response:", err)
		return
	}
	// Print the response body
	fmt.Println(string(body))

	fmt.Println("token.AccessToken>", token.AccessToken)
	fmt.Println("token.RefreshToken>", token.RefreshToken)
	fmt.Println("token.TokenType>", token.TokenType)
	fmt.Println("token.ExpiresIn>", token.ExpiresIn)

	http.Redirect(w, r, "https://social.yandex.net/broker/redirect?code="+nextcode, http.StatusFound)

}

var CodeAuth string

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func ExchangeCodeForToken(code string, clientID string, clientSecret string) (*TokenResponse, error) {
	tokenEndpoint := "https://oauth.yandex.ru/token"

	// Prepare the form data for the token request
	formData := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}

	// Make the POST request to the token endpoint
	resp, err := http.PostForm(tokenEndpoint, formData)
	if err != nil {
		fmt.Println("Error creating post request to token endpoint")
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read json")
		return nil, err
	}

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Token request failed with status: %d, response: %s", resp.StatusCode, body)
	}

	// Parse the JSON response into a TokenResponse struct
	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
func TokenOauth(w http.ResponseWriter, r *http.Request) {
	/*'grant_type'    => 'authorization_code',
	'code'          => $_GET['code'],
	'client_id'     => $clientId,
	'client_secret' => $clientSecret*/

	randomcode, _ := generateRandomCode()
	CodeAuth = randomcode
	fmt.Println("random_code:>>", randomcode)

	clientID := "4fed8408c435482b950afeb2d6e0f3cc"
	clientSecret := "dbb4420ab51f41fc86a2dedd37d2302b"

	// Replace this with the authorization code you received
	authorizationCode := randomcode

	fmt.Println("authorization_code:>>", authorizationCode)
	fmt.Println("clientID:>>", clientID)
	fmt.Println("clientSecret:>>", clientSecret)
	// Exchange the authorization code for a token
	tokenResponse, err := ExchangeCodeForToken(authorizationCode, clientID, clientSecret)
	if err != nil {
		fmt.Println("Error exchanging code for token:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Encode the TokenResponse as JSON and write it to the response
	err = json.NewEncoder(w).Encode(tokenResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %s", err), http.StatusInternalServerError)
		return
	}

	rdrNew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading response")
	}
	fmt.Println("string(rdr First Query):>", string(rdrNew))

	codeNew := r.URL.Query().Get("code")
	fmt.Println("CODE:???", codeNew)
	http.Redirect(w, r, "https://social.yandex.net/broker/redirect?code="+randomcode, http.StatusFound)

	//http.Redirect(w, r, "https://oauth.yandex.ru/authorize?client_id=4fed8408c435482b950afeb2d6e0f3cc", http.StatusFound)

	/*method := "POST"
	body := []byte(``)
	req, err := http.NewRequest(method, "https://oauth.yandex.ru/authorize?response_type=code&client_id=4fed8408c435482b950afeb2d6e0f3cc", bytes.NewReader(body))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	nextcode := r.URL.Query().Get("code")
	fmt.Println("CODE:???", nextcode)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	rdr, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response")
	}
	fmt.Println("string(rdr):>", string(rdr))
	var client map[string]interface{}
	err = json.Unmarshal(rdr, &client)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return
	}*/
	//fmt.Println("users:>", client)*/
	/*fmt.Println("users.AccessToken:>", client.ClientId)
	fmt.Println("client.RedirectURI:>", client.RedirectURI)
	fmt.Println("client.ResponseType:>", client.ResponseType)
	fmt.Println("client.Code:>", client.Code)*/

}

func RedirectPage(w http.ResponseWriter, r *http.Request) {

	r.Header.Add("X-Request-Id", "fb3f2807-3af6-4fbd-aaf2-42b5402d15e4")
	http.Redirect(w, r, "https://social.yandex.net/broker/redirect/", http.StatusFound)

	oau := models.OauthConfig
	oau.ClientID = os.Getenv("TUYA_CLIENT_ID")
	oau.ClientSecret = os.Getenv("TUYA_SECRET_KEY")
	oau.RedirectURL = fmt.Sprintf("https://social.yandex.net/broker/redirect/")

	authUrl := GenerateYandexAuthURL(oau.ClientID, oau.RedirectURL, "scope", "state")

	//url := oau.AuthCodeURL("state", oauth2.AccessTypeOffline)

	//convUrl := oau.RedirectURL + url

	//todo probably need to parse `state` from yandex response
	//http.Redirect(w, r, authUrl, http.StatusSeeOther)
	bs, _ := io.ReadAll(r.Body)
	fmt.Println("REDIRECT BODY >>> ::: ", string(bs))
	fmt.Println("URL>>>>>>>>>>:::::", authUrl)

}

func GetUserFromDbase(email string) string {

	rows, err := DB.Db.Query(`SELECT uid FROM TuyaUsers WHERE Email = ?`, email)
	if err != nil {
		fmt.Println("cant get data from dbase:", err)
		return ""
	}
	defer rows.Close()

	var uid string

	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			fmt.Println("Error scanning data:", err)
			return ""
		}
	}
	fmt.Println("Email is not nil>>>: ", uid)
	if uid == "" {
		fmt.Println("Email is nil?>>>: ", uid)
		return ""
	}
	return uid
	//fmt.Println("products: ", p)

}

func ExchangeAuthorizationCodeForToken(code string) (string, string, error) {
	conf := &oauth2.Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURL:  "your-redirect-uri",
		Endpoint: oauth2.Endpoint{
			TokenURL: "token-url",
		},
	}

	ctx := context.Background()
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return "", "", err
	}

	return token.AccessToken, token.RefreshToken, nil
}

func generateRandomState() (string, error) {
	stateLength := 32
	byteSize := stateLength / 4
	randomBytes := make([]byte, byteSize)
	_, err := rand.Read(randomBytes) //nolint
	if err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(randomBytes)
	state = state[:len(state)-1]
	return state, nil
}

func GetAuthTokenYandex(w http.ResponseWriter, r *http.Request) {
	//http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read
	method := "GET"
	clientID := models.UserData{}
	clientSecret := r.FormValue("client_secret")
	uri := fmt.Sprintf("http://localhost:9090/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=read", clientID, clientSecret)
	//body := []byte(``)
	req, _ := http.NewRequest(method, uri, nil)
	//req.Header.Set("client_id", "000000")
	//req.Header.Set("client_secret", "999999")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(bs))
}

func LogicForSynchronizeUser() {
	/*uid := service.SynchronizeUser(countryCode, userData.Email, userData.Password)
	if uid != "" {
		fmt.Println("UID_UID_UID:::", uid)
	}
	*/

	/*if r.Body != nil && r.Method == "POST" && r.Header != nil {
		// Check if the email exists
		var dbEmail string
		var dbPassword string
		err := DB.Db.QueryRow("SELECT Email, Password FROM Users WHERE Email = ?", userData.Email).Scan(&dbEmail, &dbPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				// Email doesn't exist in the database
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Email not found in the database:", err)
			} else {
				// Some other error occurred
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Error querying the database:", err)
			}
		} else {
			// Email exists in the database
			// Now, check if the passwords match
			if userData.Password != dbPassword {
				// Passwords do not match
				w.WriteHeader(http.StatusUnauthorized) // You can use 401 (Unauthorized) for this
				fmt.Println("Passwords do not match")
			} else {
				// Passwords match
				w.WriteHeader(http.StatusOK)
			}
		}
	}*/
}
