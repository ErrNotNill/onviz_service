package tuya

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Replace these with your actual Tuya credentials

// https://openapi.tuyaeu.com/v1.0/token?grant_type=1&client_id=9x8wfym7m5vyck7tdwwt&secret=d8205ed66f15471fa969aecab48ab495
type TokenResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	T       int64  `json:"t"`
	TID     string `json:"tid"`
}

func OpenConnectTuya(clientID, secretKey, baseURL, endpoint, grantType string) {
	// Create the request body parameters
	params := url.Values{}
	params.Set("client_id", clientID)
	params.Set("secret", secretKey)

	// Send the POST request
	//url = https://openapi.tuyaeu.com/v1.0/token?grant_type=1&client_id=your_client_id&secret=your_secret_key
	response, err := http.PostForm(baseURL+endpoint+grantType, params)
	//fmt.Println("response is: ", response)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println(http.StatusOK)
		// Parse the response JSON
		var data TokenResponse
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			fmt.Printf("Error decoding the response: %s\n", err)
			return
		}
		/*	{
			"code": 1108,
			"msg": "uri path invalid",
			"success": false,
			"t": 1691412745295,
			"tid": "4206274d352111ee97740a4306e7517e"
		}*/
		// Handle the API response data here
		if data.Success {
			// The API request was successful, process the data
			fmt.Printf("Code: %v\n", data.Code)
			fmt.Printf("Msg: %s\n", data.Msg)
			fmt.Printf("Success: %v\n", data.Success)
			fmt.Printf("Token Type: %v\n", data.T)
			fmt.Printf("Token ID: %s\n", data.TID)
		} else {
			// The API request encountered an error, handle the error
			fmt.Printf("Error Code: %d\n", data.Code)
			fmt.Printf("Error Message: %s\n", data.Msg)
		}
	} else {
		// Handle errors
		fmt.Printf("Error: %s\n", response.Status)
	}
}
