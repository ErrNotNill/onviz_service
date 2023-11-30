package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetInfoAboutUserTuya() {
	// /v2.0/apps/{schema}/users

	method := "GET"
	body := []byte(``)
	var uid = ""
	uri := fmt.Sprintf("/v1.0/users/%v/infos", uid)
	req, _ := http.NewRequest(method, Host+uri, bytes.NewReader(body))
	BuildHeader(req, body)

	req.Header.Add("client_id", ClientID)
	clientSecret := os.Getenv("TUYA_SECRET_KEY")
	req.Header.Add("access_token", AccessToken)
	req.Header.Add("secret", clientSecret)

	reader, err := io.ReadAll(req.Body)
	var i interface{}
	err = json.Unmarshal(reader, &i)
	if err != nil {
		fmt.Println("CANT UNMARSHALL :", err)
		//return
	}

	fmt.Println("i:interface::: ", i)
	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(bs))
}
