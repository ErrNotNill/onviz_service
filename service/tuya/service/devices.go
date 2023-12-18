package service

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

var UserUID string

func SynchronizeUser(countryCode string, username string, password string) string {
	schema := os.Getenv("TUYA_APP_KEY")
	fmt.Println("SCHEMA SCHEMA: ", schema)
	uri := fmt.Sprintf("/v1.0/apps/%v/user", schema)
	pass := encryptPassword(password)

	userInfo := fmt.Sprintf(`{
      "country_code":"%v",
      "username":"%v",
      "password":"%v",
      "username_type":2,
      "time_zone_id": ""
}`, countryCode, username, pass)

	method := "POST"
	body := []byte(userInfo)
	req, err := http.NewRequest(method, Host+uri, bytes.NewReader(body))
	if err != nil {
		log.Println("Error creating request:", err)
		return ""
	}

	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ""
	}

	fmt.Println("Response body ___ SynchronizeUser ___:", string(bs))
	var result SynchronizeResult
	err = json.Unmarshal(bs, &result)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return ""
	}
	UserUID = result.Result.UID
	fmt.Println("result: result: UID: ", result.Result.UID)
	fmt.Println("result: result: T: ", result.T)
	fmt.Println("result: result: TID: ", result.TID)
	fmt.Println("result: result: Success: ", result.Success)

	return result.Result.UID

}

func GetInfoAboutUser() {
	uid := "eu1676886479349LEc2f"
	uid = "eu1678182453992SotFy"
	uri := fmt.Sprintf("/v1.0/users/%v/infos", uid)

	method := "GET"
	body := []byte(``)
	req, err := http.NewRequest(method, Host+uri, bytes.NewReader(body))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response body ___ GetInfoAboutUser ___:", string(bs))

	var users TuyaUsers
	err = json.Unmarshal(bs, &users)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return
	}

}

func GetUsersInfo() {
	appKey := os.Getenv("TUYA_APP_KEY")
	uri := fmt.Sprintf("/v1.0/apps/%v/users?page_no=5&page_size=500&access_token=%v&sign=&t=", appKey, AccessToken)

	method := "GET"
	body := []byte(``)
	req, err := http.NewRequest(method, Host+uri, bytes.NewReader(body))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response body BS:", string(bs))

	var users TuyaUsers
	err = json.Unmarshal(bs, &users)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return
	}
	var i int
	var k int
	for i, v := range users.Result.List {
		for k = range v.Email {
			k++
		}
		i++
	}
	fmt.Println("TUYA USERS >>>>>", users)
	fmt.Println("RESULT: i: ", i)
	fmt.Println("RESULT: k: ", k)
	TransactionTuyaUsersToDb(users)
}

func GetDevicesFromUser(uid string) {
	//uid := GetUidFromTuyaUsersByEmail(userEmail)

	fmt.Println(uid, "uid_uid_uid_uid")
	uri := fmt.Sprintf("/v1.0/users/%v/devices", uid)
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

	fmt.Println("bs_bs_bs_bs_GetDevicesFromUser:::::", string(bs))
	var result ResponseDevices
	err = json.Unmarshal(bs, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//fmt.Println("result::GetDevicesFromUser:::", result)

	//TransactionDeviceToDb(result)

	for _, val := range result.Result {
		fmt.Println("val.ID", val.ID)
		fmt.Println("val.UID", val.UID)
		fmt.Println("val.CreateTime", val.CreateTime)
		fmt.Println("val.UpdateTime", val.UpdateTime)
		fmt.Println("val.Name", val.Name)
		fmt.Println("val.Status", val.Status)
		for _, status := range val.Status {
			fmt.Println("status.Code", status.Code)
			fmt.Println("status.Value", status.Value)
		}
		//fmt.Printf("val.Status type: %T", val.Status)
		fmt.Println("val.ActiveTime", val.ActiveTime)
		fmt.Println("val.BizType", val.BizType)
		fmt.Println("val.Category", val.Category)
		fmt.Println("val.Icon", val.Icon)
		fmt.Println("val.IP", val.IP)
		fmt.Println("val.LocalKey", val.LocalKey)
		fmt.Println("val.Online", val.Online)
		fmt.Println("val.OwnerID", val.OwnerID)
		fmt.Println("val.ProductID", val.ProductID)
		fmt.Println("val.ProductName", val.ProductName)
		fmt.Println("val.Sub", val.Sub)
		fmt.Println("val.TimeZone", val.TimeZone)
		fmt.Println("val.UUID", val.UUID)
	}

}

// todo fix this func, cause Response is wrong structure pointer
func GetDeviceList() {
	fmt.Println("GetDeviceList..........")
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.2/iot-03/devices/", bytes.NewReader(body))

	BuildHeader(req, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	fmt.Println("bs_bs_bs_bs_GetDevicesList:::::", string(bs))
	var response Response
	err = json.Unmarshal(bs, &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result::GetDeviceList:::", response)
	//fmt.Printf("Device ID: %s, Name: %s, Online: %v\n", deviceInfo.Result.ID, deviceInfo.Result.Name, deviceInfo.Result.Online)
	//log.Println("resp devices:", string(bs))
}

type Control struct {
	Commands []string `json:"commands"`
}

func DeviceControl(deviceId string) {
	method := "POST"
	command := `{
  "commands": [
    {
      "code": "control",
      "value": "stop"
    },
    {
      "code": "percent_control",
      "value": 0
    }
  ]
}`
	body := []byte(command)
	uriDevice := fmt.Sprintf("/v1.0/devices/%s/commands", deviceId)
	req, _ := http.NewRequest(method, Host+uriDevice, bytes.NewReader(body))
	fmt.Println("req STRUCTURE:::", req)
	BuildHeader(req, body)

	req.Header.Add("client_id", ClientID)
	clientSecret := os.Getenv("TUYA_SECRET_KEY")
	req.Header.Add("access_token", AccessToken)
	req.Header.Add("secret", clientSecret)
	//fmt.Println("deviceId:::", deviceId)
	//reader, err := io.ReadAll(req.Body)
	/*var i interface{}
	err = json.Unmarshal(reader, &i)
	if err != nil {
		fmt.Println("CANT UNMARSHALL :", err)
		//return
	}
	fmt.Println("i:interface::: ", i)
	buildHeader(req, body)*/
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	log.Println("responze::", string(bs))
}

func GetDevice(deviceId string) {

	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, Host+"/v1.0/devices/"+deviceId, bytes.NewReader(body))
	BuildHeader(req, body)

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

/*func GetDevicesList() ([]Device, error) {
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

}*/

/*func GetDevicesInProject() ([]Device, error) {
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

}*/

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
