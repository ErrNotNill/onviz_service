package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func TheTuyaAllFunctions() {
	//tuya.Cfg()
	//706020145002916b779e //my device
	//bf223afc6530ce5259mfcx
	myDeviceId := `706020145002916b779e`
	//DeviceControl(myDeviceId)
	GetToken()
	GetDevice(myDeviceId)
	GetDeviceList()
	/*userEmail := "winni3112@gmail.com"
	userEmail = "standarttechnology8891@gmail.com"
	userEmail = "vasya_57@mail.ru"
	userEmail = "zdoralex@inbox.ru"
	userEmail = "iln2005@yandex.ru"
	userEmail = "a9022798377@mail.ru"
	userEmail = "pavelrogoznikov@bk.ru"
	userEmail = "syncop@mail.ru"
	userEmail = "dorovskiy2@yandex.ru"
	userEmail = "89658088878@yandex.ru"
	userEmail = "barkov-sergey@yandex.ru"
	userEmail = "cnnmm29@gmail.com"
	userEmail = "victoriyafesenko@yandex.ru"
	userEmail = "kotchergin50@gmail.com"
	userEmail = "seniorleytan-96@yandex.ru"
	userEmail = "garibyan2006@mail.ru"
	userEmail = "shahinyan.artur@yandex.ru"
	userEmail = "kamera214@mail.ru"
	userEmail = "pasl@inbox.ru"
	userEmail = "sprokorina@mail.ru"*/

	//LoginUser()

	//fmt.Println("SYNCHRONIZE USER STARTED.............................")

	/*countryCode := "7"
	username := "standarttechnology8891@gmail.com"
	password := "htZHtFxG5728"

	SynchronizeUser(countryCode, username, password)*/

	GetUsersInfo()

	//GetInfoAboutUserTuya()

	//GetActives()
	//GetDevicesFromUser(userEmail)

	//DeviceControl(myDeviceId)
	//GetDeviceWithConnector()
	//GetUsers()
	//RefreshToken()

	/*project, err := GetDevicesInProject()
	if err != nil {
		fmt.Println("Couldn't get devices in project: ", project)

	}*/
	//tuya.PolicyAction()

	//tuya.GetUsers()
	/*devices, err := tuya.GetDevicesList()
	if err != nil {
		log.Println("No devices")
	}
	for _, device := range devices {
		fmt.Printf("ID: %v, Name: %v, Online: %v\n", device.Result, device.Success, device.T)
	}*/

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		token := &TokenResponse{}
		fmt.Println("RefreshToken Token is : ", RefreshTokenVal)
		uri := `https://openapi.tuyaeu.com/v1.0/token/` + RefreshTokenVal

		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			fmt.Println("Error creating request: ", err)
		}
		ClientID = os.Getenv("TUYA_CLIENT_ID")
		req.Header.Add("client_id", ClientID)
		clientSecret := os.Getenv("TUYA_SECRET_KEY")
		req.Header.Add("access_token", AccessToken)
		req.Header.Add("secret", clientSecret)
		ts := fmt.Sprintf("%d", time.Now().UTC().UnixNano()/int64(time.Millisecond))
		req.Header.Add("sign_method", "HMAC-SHA256")
		fmt.Println("ts:", ts)
		fmt.Println("strconv.Itoa(int(TimeToken))", strconv.Itoa(int(TimeToken)))
		req.Header.Add("t", strconv.Itoa(int(TimeToken)))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error creating client: ", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		json.Unmarshal(body, &token)
	}
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	clientID := os.Getenv("TUYA_CLIENT_ID")
	redirectURI := YaRedirectUri

	authURL := GenerateAuthorizationURL(clientID, redirectURI)
	response := struct {
		AuthURL string `json:"auth_url"`
	}{
		AuthURL: authURL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	reader := r.Body
	var yaParams = YandexAuthParams{}

	js, err := json.MarshalIndent(reader, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(js, &yaParams)
	w.Write(js)
	w.Write([]byte(yaParams.State))

	http.Redirect(w, r, "https://social.yandex.net/broker/redirect", http.StatusTemporaryRedirect)

}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	//token := os.Getenv("TOKEN")
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.0/token?grant_type=1", bytes.NewReader(body))

	BuildHeader(req, body)
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

	if refToken := ret.Result.RefreshToken; refToken != "" {
		RefreshTokenVal = refToken
	}
	log.Println("Token:", Token)
	log.Println("Refresh Token:", RefreshTokenVal)
	w.Write([]byte(Token))
	w.Write([]byte(RefreshTokenVal))

	if uid := ret.Result.UID; uid != "" {
		w.Write([]byte("Uid is: " + Uid))
		http.Redirect(w, r, "https://social.yandex.net/broker/redirect", http.StatusFound)
	}

}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	go GetTokenHandler(w, r)
	clientID := ClientID
	refreshToken := RefreshTokenVal
	if clientID == "" || refreshToken == "" {
		http.Error(w, "Missing client_id or refresh_token", http.StatusBadRequest)
		return
	}
	response := struct {
		AccessToken string `json:"refresh_token"`
	}{
		AccessToken: RefreshTokenVal,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
