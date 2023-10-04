package tuya

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/tuya/tuya-connector-go/connector"
	"github.com/tuya/tuya-connector-go/connector/logger"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetDeviceList() {
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

	req.Header.Add("client_id", ClientID)
	clientSecret := os.Getenv("TUYA_SECRET_KEY")
	req.Header.Add("access_token", AccessToken)
	req.Header.Add("secret", clientSecret)
	fmt.Println("deviceId:::", deviceId)
	reader, err := io.ReadAll(req.Body)
	var i interface{}
	err = json.Unmarshal(reader, &i)
	if err != nil {
		fmt.Println("CANT UNMARSHALL :", err)
		//return
	}
	fmt.Println("i:interface::: ", i)
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
	urlz := "https://openapi.tuyaeu.com/v2.0/apps/schema/users"

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

	} else {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
	}

}

const (
	BaseURL     = "https://openapi.tuyaeu.com/v1.0"
	DevicesPath = "/devices"
)

type Devices struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func generateSignature(accessID, path, secretKey string) string {
	data := accessID + path
	hmacHash := hmac.New(sha256.New, []byte(secretKey))
	hmacHash.Write([]byte(data))
	signature := fmt.Sprintf("%x", hmacHash.Sum(nil))
	fmt.Println("SIGNATURE : ", signature)
	return signature
}

func GetDevicesList() ([]Device, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://openapi.tuyaeu.com/v1.0/devices", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("access_token is: ", AccessToken)
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	fmt.Println("Uid : ", Uid)
	req.Header.Add("client_id", Uid)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to create client: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK", http.StatusOK)
		body, _ := io.ReadAll(resp.Body)
		var deviceListResponse DeviceListResponse
		if err := json.Unmarshal(body, &deviceListResponse); err != nil {
			return nil, err
		}
		if deviceListResponse.Success {
			return deviceListResponse.Result, nil
		} else {
			return nil, fmt.Errorf("API request failed: %s", string(body))
		}
	} else {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

}

func GetDevicesInProject() ([]Device, error) {
	//fmt.Println("Getting devices")
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://openapi.tuyaeu.com/v2.0/cloud/thing/device", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("access_token is: ", AccessToken)
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	fmt.Println("Uid : ", Uid)
	//req.Header.Add("client_id", Uid)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to create client: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK", http.StatusOK)
		body, _ := io.ReadAll(resp.Body)
		var deviceListResponse DeviceListResponse
		if err := json.Unmarshal(body, &deviceListResponse); err != nil {
			return nil, err
		}
		if deviceListResponse.Success {
			return deviceListResponse.Result, nil
		} else {
			return nil, fmt.Errorf("API request failed: %s", string(body))
		}
	} else {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

}

// GetDevice  Call OpenAPI (taking Gin framework for example)
func GetDeviceWithConnector() {
	resp := &GetDeviceResponse{}
	// Initiate an API request
	err := connector.MakeGetRequest(
		context.Background(),
		connector.WithAPIUri(fmt.Sprintf("/v1.0/devices/%s", DeviceID)),
		connector.WithResp(resp))
	connector.WithErrProc(1102, &DeviceError{})
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}

// GetDeviceResponse Data structure returned by OpenAPI
type GetDeviceResponseWithConnector struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	T       int64       `json:"t"`
}

// DeviceError the struct class that implements the IError interface.
type DeviceError struct {
}

func (d *DeviceError) Process(ctx context.Context, code int, msg string) {
	logger.Log.Error(code, msg)
}
