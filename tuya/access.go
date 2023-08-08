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

const (
	Host     = "https://openapi.tuyaeu.com"
	ClientID = "9x8wfym7m5vyck7tdwwt"
	Secret   = "TUYA_SECRET"
	DeviceID = ""
)

var (
	Token string
)

type TokenResponseOEM struct {
	Result struct {
		AccessToken  string `json:"access_token"`
		ExpireTime   int    `json:"expire_time"`
		RefreshToken string `json:"refresh_token"`
		UID          string `json:"uid"`
	} `json:"result"`
	Success bool  `json:"success"`
	T       int64 `json:"t"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetDevices(accessToken string) ([]Device, error) {
	method := "GET"
	client := &http.Client{}
	body := []byte(``)
	req, err := http.NewRequest(method, Host+"/v1.2/iot-03/devices", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var devices []Device
	if err := json.NewDecoder(resp.Body).Decode(&devices); err != nil {
		return nil, err
	}

	return devices, nil
}

/*func GetDevices() {
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.2/iot-03/devices", bytes.NewReader(body))
	buildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	ret := TokenResponseOEM{}
	err = json.Unmarshal(bs, &ret)
	if err != nil {
		fmt.Println(" unmarshal error:", err)
		return
	}
	log.Println("resp:", string(bs))

	if v := ret.Result.AccessToken; v != "" {
		Token = v
	}
}*/

func GetToken(Host string) string {
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.0/token?grant_type=1", bytes.NewReader(body))

	buildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	ret := TokenResponseOEM{}
	err = json.Unmarshal(bs, &ret)
	if err != nil {
		fmt.Println(" unmarshal error:", err)
		return ""
	}
	log.Println("resp:", string(bs))

	if v := ret.Result.AccessToken; v != "" {
		Token = v
	}
	return Token
}

func GetDevice(deviceId string) {
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.0/devices/"+deviceId, bytes.NewReader(body))

	buildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(bs))
}

func buildHeader(req *http.Request, body []byte) {
	req.Header.Set("client_id", ClientID)
	req.Header.Set("sign_method", "HMAC-SHA256")

	ts := fmt.Sprint(time.Now().UnixNano() / 1e6)
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
