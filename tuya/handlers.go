package tuya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func TheTuyaAllFunctions() {
	//tuya.Cfg()
	//706020145002916b779e //my device
	//bf223afc6530ce5259mfcx
	myDeviceId := `706020145002916b779e`
	GetToken()
	GetDevice(myDeviceId)
	DeviceControl(myDeviceId)
	//GetDeviceWithConnector()
	//GetUsers()
	//RefreshToken()

	project, err := GetDevicesInProject()
	if err != nil {
		fmt.Println("Couldn't get devices in project: ", project)

	}
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
