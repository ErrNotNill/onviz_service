package tuya

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func LoginUser() {

	userName := "onvizbitrix@gmail.com"
	userPass := "htZHtFxG5728"

	uri := fmt.Sprintf("/v1.0/iot-03/users/login")

	method := "POST"
	encryptedPassword := encryptPassword(userPass)

	fmt.Println("encryptedPassword", encryptedPassword)
	credentials := Credentials{
		Username: userName,
		Password: encryptedPassword,
	}
	fmt.Println("credentials", credentials)
	// Marshal the struct into a JSON string
	body, err := json.Marshal(credentials)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, _ := http.NewRequest(method, Host+uri, bytes.NewReader(body))

	buildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	var result LoginResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result::LoginResponse:::", result)

}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func encryptPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashed := hex.EncodeToString(hash.Sum(nil))
	lowercaseHashed := strings.ToLower(hashed)
	return lowercaseHashed
}

type LoginResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UID          string `json:"uid"`
	Expire       int    `json:"expire"`
}

type LoginResponse struct {
	Result  Result `json:"result"`
	T       int64  `json:"t"`
	Success bool   `json:"success"`
}
