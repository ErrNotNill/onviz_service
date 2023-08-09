package tuya

import (
	"bytes"
	"encoding/json"
	"net/http"
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
