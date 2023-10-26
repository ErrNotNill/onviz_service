package bot_bitrix

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

const FSPATH = "vite-project/index.html"

type Data struct {
	Title string
	Size  string
}

func BotBitrix(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	tokenURL := fmt.Sprintf(`https://onviz.bitrix24.ru/oauth/authorize/?
	client_id=%s`,
		BitrixClientId)

	post, err := client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}
	body, err := io.ReadAll(post.Body)
	fmt.Println("post.Body", post.Body)
	json.Unmarshal(body, &post.Body)
	fmt.Fprint(w, string(body))
	//fs := http.FileServer(http.Dir(FSPATH))
	//fs.ServeHTTP(w, r)
	//serverUri := os.Getenv("SERVER_URL")
	//	http.Redirect(w, r, serverUri, http.StatusMovedPermanently)
	t, err := template.ParseFiles("bot_bitrix/bot_bitrix.html")
	if err != nil {
		fmt.Println(err)
	}
	var data = Data{Title: "Тип одежды", Size: "Размер"}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	//w.WriteHeader(http.StatusOK)
}

var (
	localBitrixAppID     = os.Getenv("LOCAL_BITRIX_APP_ID")
	localBitrixAppSecret = os.Getenv("LOCAL_BITRIX_APP_SECRET")
	accessToken          = "YOUR_ACCESS_TOKEN"
	bitrixDomain         = "onviz.bitrix24.ru"
)

type APIResponse struct {
	Result []OpenLine `json:"result"`
}

type OpenLine struct {
	ID   int    `json:"ID"`
	Name string `json:"NAME"`
}

func OpenLines(w http.ResponseWriter, r *http.Request) {
	apiEndpoint := fmt.Sprintf("https://%s/rest/imopenlines.bot.list", bitrixDomain)
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		fmt.Println("Failed to create API request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to make API request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var apiResp APIResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResp)
		if err != nil {
			fmt.Println("Failed to parse JSON response:", err)
			return
		}
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
