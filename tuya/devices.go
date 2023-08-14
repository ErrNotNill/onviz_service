package tuya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GetDevicesWithToken() {
	accessToken := Token // Replace with your actual access token

	req, err := http.NewRequest("GET", "https://openapi.tuyaeu.com/v1.0/devices?", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("client_id", ClientID)

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

func GetUsers() {
	urlz := "https://openapi.tuyacn.com/v2.0/apps/schema/users"

	// Configure start_time and end_time
	startTimeStr := "2023-05-22 20:36:55"
	endTime := time.Now()

	startTime, err := time.Parse("2006-01-02 15:04:05", startTimeStr)
	if err != nil {
		fmt.Println("Error parsing start_time:", err)
		return
	}

	params := url.Values{}
	params.Set("page_no", "1")
	params.Set("page_size", "50")
	params.Set("start_time", fmt.Sprintf("%d", startTime.Unix()))
	params.Set("end_time", fmt.Sprintf("%d", endTime.Unix()))

	urlz += "?" + params.Encode()

	req, err := http.NewRequest("GET", urlz, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful!")

		// Read and print the response body or process it as needed
	} else {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
	}

}