package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	models2 "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"onviz/DB"
	"onviz/internal/user/models"
	"onviz/service/tuya/service"
	"os"
)

func CreateAccount(name, email, password string) error {
	fmt.Println("db connected")
	tokenData := email + password // Customize this according to your needs
	sha256Hash := sha256.Sum256([]byte(tokenData))
	token := hex.EncodeToString(sha256Hash[:])
	result, err := DB.Db.Exec(`insert into Users (Name, Email, Password, Token) values (?,?,?,?)`,
		name, email, password, token)

	if err != nil {
		fmt.Println("cant insert data to dbase")
		//panic(err)
		return err
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
	return nil
}

func GetAccount(email, password string) models.UserData {

	rows, err := DB.Db.Query(`SELECT Email, Password FROM Users WHERE Email = ? AND Password = ?`, email, password)
	if err != nil {
		fmt.Println("cant get data from dbase:", err)
		return models.UserData{}
	}
	defer rows.Close()

	p := models.UserData{}

	for rows.Next() {
		err := rows.Scan(&p.Email, &p.Password)
		if err != nil {
			fmt.Println("Error scanning data:", err)
			return models.UserData{}
		}
	}
	fmt.Println("Email is not nil>>>: ", p.Email)
	if p.Email == "" {
		fmt.Println("Email is>>>: ", p.Email)
		return models.UserData{}
	}
	return p
	//fmt.Println("products: ", p)

}

func AuthPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var userData models.UserData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	tokenData := userData.Email + userData.Password // Customize this according to your needs
	sha256Hash := sha256.Sum256([]byte(tokenData))
	token := hex.EncodeToString(sha256Hash[:])
	// Process user registration data (userData) as needed
	fmt.Printf("Received registration data: %+v\n", userData)
	// You can now handle the registration logic, such as storing the data in a database

	// Handle registration error
	if err := CreateAccount(userData.Username, userData.Email, userData.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "User exist"})
		fmt.Println("Error creating, user exists")
		return
	}

	// Send a response back to the client
	response := map[string]string{"token": token}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set the response status once, indicating a successful response
	w.Write(responseJSON)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the callback after the user authorizes your application
	code := r.URL.Query().Get("code")

	// Exchange the authorization code for an access token
	token, err := OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging code for token: %v", err)
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		return
	}

	// You can now use the access token to make authenticated requests to the OAuth provider's API
	fmt.Printf("Access Token: %v\n", token.AccessToken)

	// Your existing code...
}

func LoginPage(w http.ResponseWriter, r *http.Request) {

	var userData models.UserData
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	rdr, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(rdr, &userData)
	if err != nil {
		log.Println("Error decoding JSON")
	}

	fmt.Println("userData: ", userData)
	// Process user registration data (userData) as needed
	fmt.Printf("Received registration data: %+v\n", userData)
	// You can now handle the registration logic, such as storing the data in a database
	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("userData.Email", userData.Email)
	fmt.Println("userData.Country", userData.Country)

	countryCode := service.GetCountryCodeFromDbase(userData.Country)

	fmt.Printf("countryCode : %v, userData.Email : %v, userData.Password : %v", countryCode, userData.Email, userData.Password)

	if userData.Email != "" {

		uid := GetUserFromDbase(userData.Email)
		if uid != "" {
			oau := OauthConfig
			oau.ClientID = uid
			oau.ClientSecret = os.Getenv("TUYA_SECRET_KEY")
			oau.RedirectURL = fmt.Sprintf("https://social.yandex.net/broker/redirect")

			url := oau.AuthCodeURL("state", oauth2.AccessTypeOffline)

			//todo probably need to parse `state` from yandex response
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			bs, _ := io.ReadAll(r.Body)
			fmt.Println("REDIRECT BODY >>>> ::: ", string(bs))
			fmt.Println("URL>>>>>>>>>>:::::", url)
			fmt.Println("redirect started")
		}

		service.GetDevicesFromUser(uid)
		fmt.Println("uid_uid_uid::::", uid)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	/*uid := service.SynchronizeUser(countryCode, userData.Email, userData.Password)
	if uid != "" {
		fmt.Println("UID_UID_UID:::", uid)
	}
	*/

	/*if r.Body != nil && r.Method == "POST" && r.Header != nil {
		// Check if the email exists
		var dbEmail string
		var dbPassword string
		err := DB.Db.QueryRow("SELECT Email, Password FROM Users WHERE Email = ?", userData.Email).Scan(&dbEmail, &dbPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				// Email doesn't exist in the database
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Email not found in the database:", err)
			} else {
				// Some other error occurred
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Error querying the database:", err)
			}
		} else {
			// Email exists in the database
			// Now, check if the passwords match
			if userData.Password != dbPassword {
				// Passwords do not match
				w.WriteHeader(http.StatusUnauthorized) // You can use 401 (Unauthorized) for this
				fmt.Println("Passwords do not match")
			} else {
				// Passwords match
				w.WriteHeader(http.StatusOK)
			}
		}
	}*/
}

func GetUserFromDbase(email string) string {

	rows, err := DB.Db.Query(`SELECT uid FROM TuyaUsers WHERE Email = ?`, email)
	if err != nil {
		fmt.Println("cant get data from dbase:", err)
		return ""
	}
	defer rows.Close()

	var uid string

	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			fmt.Println("Error scanning data:", err)
			return ""
		}
	}
	fmt.Println("Email is not nil>>>: ", uid)
	if uid == "" {
		fmt.Println("Email is nil?>>>: ", uid)
		return ""
	}
	return uid
	//fmt.Println("products: ", p)

}

func ExchangeAuthorizationCodeForToken(code string) (string, string, error) {
	conf := &oauth2.Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURL:  "your-redirect-uri",
		Endpoint: oauth2.Endpoint{
			TokenURL: "token-url",
		},
	}

	ctx := context.Background()
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return "", "", err
	}

	return token.AccessToken, token.RefreshToken, nil
}

func Auth() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models2.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:9090",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
}

func GetAuthTokenYandex(w http.ResponseWriter, r *http.Request) {
	//http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read
	method := "GET"
	clientID := models.UserData{}
	clientSecret := r.FormValue("client_secret")
	uri := fmt.Sprintf("http://localhost:9090/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=read", clientID, clientSecret)
	//body := []byte(``)
	req, _ := http.NewRequest(method, uri, nil)
	//req.Header.Set("client_id", "000000")
	//req.Header.Set("client_secret", "999999")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(bs))
}
