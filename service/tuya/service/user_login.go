package service

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

func GetActives() {

	uri := fmt.Sprintf("/v1.0/iot-03/users/assets")

	method := "GET"
	body := []byte(``)

	req, _ := http.NewRequest(method, Host+uri, bytes.NewReader(body))

	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	fmt.Println("bs_bs_bs: ", string(bs))

	var result LoginResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result::LoginResponse:::", result)

}

func SchemaUser() {
	//p1683791449319ma4w9u
	//https://api.cakes.com/v1.0/apps/{schema}/users?page_no=&page_size=&access_token=&sign=&t=
}

func LoginUser() {

	userName := "commongoverygoodguy@gmail.com"
	userPass := "htZHtFxG5728"

	uri := fmt.Sprintf("/v1.0/iot-03/users/login")

	method := "POST"
	encryptedPassword := encryptPassword(userPass)

	b := fmt.Sprintf(`{
  "username": "%v",
  "password": "%v"
}`, userName, encryptedPassword)

	fmt.Println("b_b_b_b__b_b_b_b_b: ", b)
	fmt.Println("encryptedPassword", encryptedPassword)

	body := []byte(b)

	req, _ := http.NewRequest(method, Host+uri, bytes.NewReader(body))

	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	fmt.Println("bs_bs_LoginUser_bs_bs : ", string(bs))

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
