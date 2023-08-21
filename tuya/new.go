package tuya

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Config struct {
	Host      string
	AccessKey string
	SecretKey string
	DeviceID  string
}

type App struct {
	Config Config
	Client *http.Client
	Token  string
}

var Deviceid = "bay1675253638992fSA4"

func Cfg() {
	config := Config{
		Host:      "https://openapi.tuyaeu.com",
		AccessKey: "9x8wfym7m5vyck7tdwwt",             //TUYA_CLIENT_ID
		SecretKey: "d8205ed66f15471fa969aecab48ab495", //TUYA_SECRET_KEY
		DeviceID:  Deviceid,
	}

	app := App{
		Config: config,
		Client: &http.Client{Timeout: 5 * time.Second},
	}

	switchValue := true // Set your switch value here
	err := app.getToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := app.getDeviceInfo(config.DeviceID, switchValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("fetch success: %+v\n", data)
}

func (app *App) getToken() error {
	method := "GET"
	timestamp := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	signURL := "/v1.0/token?grant_type=1"
	contentHash := sha256.Sum256([]byte(""))
	stringToSign := strings.Join([]string{method, hex.EncodeToString(contentHash[:]), "", signURL}, "\n")
	signStr := app.Config.AccessKey + timestamp + stringToSign

	headers := map[string]string{
		"t":           timestamp,
		"sign_method": "HMAC-SHA256",
		"client_id":   app.Config.AccessKey,
		"sign":        app.encryptStr(signStr, app.Config.SecretKey),
	}

	loginURL := app.Config.Host + "/v1.0/token?grant_type=1"
	loginResp, err := app.makeRequest("GET", loginURL, headers, nil)
	if err != nil {
		return fmt.Errorf("fetch failed: %v", err)
	}
	defer loginResp.Body.Close()

	body, err := io.ReadAll(loginResp.Body)
	if err != nil {
		fmt.Println("read failed: ", err)
		return err
	}
	bs, _ := io.ReadAll(loginResp.Body)
	ret := TokenResponse{}
	json.Unmarshal(bs, &ret)
	log.Println("resp:", string(bs))

	if v := ret.Result.AccessToken; v != "" {
		Token = v
	}
	//here we got AccessToken and UID / clientID
	AccessToken = ret.Result.AccessToken
	Uid = ret.Result.UID

	if refToken := ret.Result.RefreshToken; refToken != "" {
		RefreshTokenVal = refToken
	}
	log.Println("Token:", Token)
	log.Println("Refresh Token:", RefreshTokenVal)
	// Parse and use response as needed
	_ = body
	return nil
}

func (app *App) getDeviceInfo(deviceID string, switchValue bool) (interface{}, error) {
	method := "GET"
	url := fmt.Sprintf("/v1.0/devices/%s/commands", deviceID)
	reqHeaders, err := app.getRequestSign(url, method, map[string]string{}, map[string]string{})
	if err != nil {
		fmt.Println("Error getting getDeviceInfo")
		return nil, err
	}

	data := map[string]interface{}{
		"commands": []map[string]interface{}{
			{"code": "countdown_1", "value": 0},
			{"code": "switch", "value": switchValue},
		},
	}

	resp, err := app.makeRequest(method, reqHeaders["path"], reqHeaders, data)
	if err != nil {
		return nil, fmt.Errorf("request api failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request read failed:")
		return nil, err
	}

	// Parse and use response as needed
	_ = respBody

	return nil, nil
}

func (app *App) encryptStr(str, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (app *App) getRequestSign(path, method string, headers, query map[string]string) (map[string]string, error) {
	timestamp := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	uri := path
	queryMerged := make(map[string]string)
	for k, v := range query {
		queryMerged[k] = v
	}

	sortedQuery := make([]string, 0, len(queryMerged))
	for k := range queryMerged {
		sortedQuery = append(sortedQuery, k)
	}
	sort.Strings(sortedQuery)

	queryString := url.Values{}
	for _, k := range sortedQuery {
		queryString[k] = []string{queryMerged[k]}
	}

	fullURL := uri
	if len(queryString) > 0 {
		fullURL += "?" + queryString.Encode()
	}

	contentHash := sha256.Sum256([]byte("{}"))
	clientID := app.Config.AccessKey
	accessToken := app.Token
	stringToSign := strings.Join([]string{method, hex.EncodeToString(contentHash[:]), "", fullURL}, "\n")
	signStr := clientID + accessToken + timestamp + stringToSign

	reqHeaders := map[string]string{
		"t":            timestamp,
		"path":         fullURL,
		"client_id":    clientID,
		"sign":         app.encryptStr(signStr, app.Config.SecretKey),
		"sign_method":  "HMAC-SHA256",
		"sign_version": "2.0",
		"access_token": accessToken,
	}

	return reqHeaders, nil
}

func (app *App) makeRequest(method, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if method == "POST" && body != nil {
		// Serialize the request body and set content type header
		// Here you should serialize the 'body' parameter according to the API's requirements
		// and set the appropriate content type header
		req.Header.Set("Content-Type", "application/json")
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(reqBody))
	}

	return app.Client.Do(req)
}
