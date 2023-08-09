package bot_bitrix

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const FSPATH = "vite-project/index.html"

func BotBitrix(w http.ResponseWriter, r *http.Request) {

	fs := http.FileServer(http.Dir(FSPATH))
	fs.ServeHTTP(w, r)

	http.Redirect(w, r, "http://45.141.79.120:5173/", http.StatusMovedPermanently)
	t, err := template.ParseFiles("bot_bitrix/bot_bitrix.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//w.Write([]byte("OK"))
	//w.WriteHeader(http.StatusOK)
}

// Replace these with your actual credentials
const (
	clientID     = "local.64c7a198b979a3.49589400"
	clientSecret = "BHrEY2UHSpx8HcxaImxDIGghA7EZ0EFFa0empe9INxBtElOEpR"
	accessToken  = "YOUR_ACCESS_TOKEN"
	bitrixDomain = "onviz.bitrix24.ru"
)

// APIResponse represents the response from the Bitrix24 API
type APIResponse struct {
	Result []OpenLine `json:"result"`
}

// OpenLine represents the structure of an Open Line
type OpenLine struct {
	ID   int    `json:"ID"`
	Name string `json:"NAME"`
	// Add more fields here if needed
}

func OpenLines(w http.ResponseWriter, r *http.Request) {
	// API endpoint to get the list of Open Lines

	apiEndpoint := fmt.Sprintf("https://%s/rest/imopenlines.bot.list", bitrixDomain)

	// Set up the HTTP client
	client := &http.Client{}

	// Create the API request with the access token
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		fmt.Println("Failed to create API request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	// Make the API request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to make API request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode == http.StatusOK {
		// Parse the JSON response
		var apiResp APIResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResp)
		if err != nil {
			fmt.Println("Failed to parse JSON response:", err)
			return
		}

		// Extract the ID of the first Open Line (assuming it exists)
		if len(apiResp.Result) > 0 {
			openLineID := apiResp.Result[0].ID
			fmt.Printf("The ID of the first Open Line is: %d\n", openLineID)
		} else {
			fmt.Println("No Open Lines found.")
		}
	} else {
		fmt.Println("Failed to retrieve Open Lines:", resp.Status)
	}
}
