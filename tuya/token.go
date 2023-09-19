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
	"sort"
	"strings"
	"time"
)

func RefreshToken() {

	token := &TokenResponse{}
	fmt.Println("RefreshToken Token is : ", RefreshTokenVal)
	//values := url.Values{}
	//values.Set("grant_type", "refresh_token")
	//values.Set("client_id", clientID)
	//values.Set("refresh_token", refreshToken)
	uri := `https://openapi.tuyaeu.com/v1.0/token/` + RefreshTokenVal

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
	}
	//req.Header.Add("grant_type", "refresh_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating client: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	fmt.Println("string(body).RefreshToken : ", string(body))
	fmt.Println("token.Result.AccessToken: ", token.Result.AccessToken)
	fmt.Println("token.Result.RefreshToken: ", token.Result.RefreshToken)

}

func GetToken() {
	//token := os.Getenv("TOKEN")
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.0/token?grant_type=1", bytes.NewReader(body))

	buildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	ret := TokenResponse{}
	json.Unmarshal(bs, &ret)
	log.Println("resp:", string(bs))

	if v := ret.Result.AccessToken; v != "" {
		Token = v
	}
	//here we got AccessToken and UID / clientID
	AccessToken = ret.Result.AccessToken

	Uid = ret.Result.UID

	refToken := ret.Result.RefreshToken
	fmt.Println("refToken :", string(refToken))

	RefreshTokenVal = refToken

	log.Println("Token:", Token)
	log.Println("Refresh Token:", RefreshTokenVal)
	list, err := GetDevicesList()
	if err != nil {
		fmt.Println("Error getting devices list")
		return
	}
	for _, v := range list {
		fmt.Println("Device:", v)
	}

}

func buildHeader(req *http.Request, body []byte) {
	req.Header.Set("client_id", ClientID)
	req.Header.Set("sign_method", "HMAC-SHA256")

	ts := fmt.Sprintf("%d", time.Now().UTC().UnixNano()/int64(time.Millisecond))
	fmt.Println("ts:", ts)
	req.Header.Set("t", ts)

	if Token != "" {
		req.Header.Set("access_token", Token)
	}

	sign := buildSign(req, body, ts)
	req.Header.Set("sign", sign)
}

func buildSign(req *http.Request, body []byte, t string) string {
	headers := getHeaderStr(req)
	urlStr := getUrlStr(req)
	contentSha256 := Sha256(body)
	stringToSign := req.Method + "\n" + contentSha256 + "\n" + headers + "\n" + urlStr
	signStr := ClientID + Token + t + stringToSign
	sign := strings.ToUpper(HmacSha256(signStr, Secret))
	return sign
}

func Sha256(data []byte) string {
	sha256Contain := sha256.New()
	sha256Contain.Write(data)
	return hex.EncodeToString(sha256Contain.Sum(nil))
}

func getUrlStr(req *http.Request) string {
	url := req.URL.Path
	keys := make([]string, 0, 10)

	query := req.URL.Query()
	for key, _ := range query {
		keys = append(keys, key)
	}
	if len(keys) > 0 {
		url += "?"
		sort.Strings(keys)
		for _, keyName := range keys {
			value := query.Get(keyName)
			url += keyName + "=" + value + "&"
		}
	}

	if url[len(url)-1] == '&' {
		url = url[:len(url)-1]
	}
	return url
}

func getHeaderStr(req *http.Request) string {
	signHeaderKeys := req.Header.Get("Signature-Headers")
	if signHeaderKeys == "" {
		return ""
	}
	keys := strings.Split(signHeaderKeys, ":")
	headers := ""
	for _, key := range keys {
		headers += key + ":" + req.Header.Get(key) + "\n"
	}
	return headers
}

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
