package tuya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetDevicesWithToken() {
	accessToken := Token // Replace with your actual access token

	req, err := http.NewRequest("GET", "https://openapi.tuyaeu.com/v1.0/devices", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("client_id", ClientID)
	timestamp := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	req.Header.Set("timestamp", timestamp)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	var deviceInfo Device
	err = json.Unmarshal(bs, &deviceInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Device ID: %s, Name: %s, Online: %v\n", deviceInfo.Result.ID, deviceInfo.Result.Name, deviceInfo.Result.Online)
	log.Println("resp devices:", string(bs))

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful!")
		// Handle the response body here
	} else {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
	}
}

func GetDeviceList() {
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.0/devices", bytes.NewReader(body))

	buildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	var deviceInfo Device
	err = json.Unmarshal(bs, &deviceInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//fmt.Printf("Device ID: %s, Name: %s, Online: %v\n", deviceInfo.Result.ID, deviceInfo.Result.Name, deviceInfo.Result.Online)
	//log.Println("resp devices:", string(bs))
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

type Device struct {
	Result  DeviceInfo `json:"result"`
	Success bool       `json:"success"`
	T       int64      `json:"t"`
	TID     string     `json:"tid"`
}

type DeviceInfo struct {
	ActiveTime  int64        `json:"active_time"`
	BizType     int64        `json:"biz_type"`
	Category    string       `json:"category"`
	CreateTime  int64        `json:"create_time"`
	Icon        string       `json:"icon"`
	ID          string       `json:"id"`
	IP          string       `json:"ip"`
	Lat         string       `json:"lat"`
	LocalKey    string       `json:"local_key"`
	Lon         string       `json:"lon"`
	Model       string       `json:"model"`
	Name        string       `json:"name"`
	NodeID      string       `json:"node_id"`
	Online      bool         `json:"online"`
	OwnerID     string       `json:"owner_id"`
	ProductID   string       `json:"product_id"`
	ProductName string       `json:"product_name"`
	Status      []StatusInfo `json:"status"`
	Sub         bool         `json:"sub"`
	TimeZone    string       `json:"time_zone"`
	UID         string       `json:"uid"`
	UpdateTime  int64        `json:"update_time"`
	UUID        string       `json:"uuid"`
}

type StatusInfo struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}
