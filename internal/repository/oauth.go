package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	models2 "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"io"
	"log"
	"net/http"
	"onviz/internal/user/models"
	"onviz/service/tuya/service"
)

var YandexClientId string

func LoginPage(w http.ResponseWriter, r *http.Request) {
	//r.Header.Get("X-Request-Id")
	client_id := r.URL.Query().Get("client_id")
	fmt.Println("client_id", client_id)
	YandexClientId = client_id

	http.Redirect(w, r, "https://onviz-api.ru", http.StatusFound)

	readerFromYandex, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading from yandex")
	}
	fmt.Println("string(readerFromYandex):::", string(readerFromYandex))

	//w.WriteHeader(http.StatusOK)
	var userData models.UserDataOnviz
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err = json.Unmarshal(readerFromYandex, &userData)
	if err != nil {
		log.Println("Error decoding JSON", err.Error())
	}

	countryCode := service.GetCountryCodeFromDbase(userData.Country)
	fmt.Printf("countryCode : %v, userData.Email : %v, userData.Password : %v, userData.ClientID: %v, userData.ClientSecret: %v", countryCode, userData.Email, userData.Password, userData.ClientID, userData.ClientSecret)

	if userData.Email != "" {

		fmt.Printf("countryCode : %v, userData.Email : %v, userData.Password : %v, userData.ClientID: %v, userData.ClientSecret: %v", countryCode, userData.Email, userData.Password, userData.ClientID, userData.ClientSecret)

		uid := GetUserFromDbase(userData.Email)

		if uid != "" {
			err := json.NewEncoder(w).Encode("tokenFromEncode")
			if err != nil {
				log.Println("Error encoding token")
			}
		}
		//http.Redirect(w, r, "https://social.yandex.net/broker/redirect?"+YandexClientId, http.StatusFound)
		//service.GetDevicesFromUser(uid)
		fmt.Println("uid_uid_uid::::", uid)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func Auth() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	err := clientStore.Set("9x8wfym7m5vyck7tdwwt", &models2.Client{
		ID:     "9x8wfym7m5vyck7tdwwt",
		Secret: "d8205ed66f15471fa969aecab48ab495",
		Domain: "https://social.yandex.net/broker/redirect",
	})
	if err != nil {
		log.Println("Error sett client", err.Error())
		return
	}
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

	http.HandleFunc("/api/authorize", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("client_id", "9x8wfym7m5vyck7tdwwt")
		r.Header.Add("client_secret", "d8205ed66f15471fa969aecab48ab495")
		err := srv.HandleAuthorizeRequest(w, r)
		//ExchangeAuthorizationCodeForToken() //todo where to get code for this method
		if err != nil {
			log.Println("HandeAuthorizeError")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			log.Println("Handler token error")
			return
		}
	})
}

func AccessToLoginPage(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	state := r.FormValue("state")
	redirectURI := r.FormValue("redirect_uri")
	responseType := r.FormValue("response_type")
	clientID := r.FormValue("client_id")
	scope := r.FormValue("scope")

	// Log the extracted parameters (you can customize this part)
	log.Printf("Received OAuth parameters:\nState: %s\nRedirect URI: %s\nResponse Type: %s\nClient ID: %s\nScope: %s\n",
		state, redirectURI, responseType, clientID, scope)
	log.Println("State is: ", state)
	log.Println("redirectURI is: ", redirectURI)
	log.Println("responseType is: ", responseType)
	log.Println("clientID is: ", clientID)
	log.Println("scope is: ", scope)
	//log.Println("State NEW: ", splState)

	redirectURL := fmt.Sprintf("%s?state=%s&response_type=%s&client_id=%s&scope=%s",
		redirectURI, state, responseType, clientID, "scope")

	log.Println("Redirect URL is: ", redirectURL)
	// Use http.Redirect to perform the redirect
	http.Redirect(w, r, redirectURL, http.StatusFound)

	bs, _ := io.ReadAll(r.Body)
	fmt.Println("rdr:::", string(bs))

}
