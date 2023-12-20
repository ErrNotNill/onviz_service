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
	"os"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("X-Request-Id")

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
	fmt.Printf("countryCode : %v, userData.Email : %v, userData.Password : %v", userData.Country, userData.Email, userData.Password)

	countryCode := service.GetCountryCodeFromDbase(userData.Country)

	if userData.Email != "" {

		fmt.Printf("countryCode : %v, userData.Email : %v, userData.Password : %v", countryCode, userData.Email, userData.Password)
		uid := GetUserFromDbase(userData.Email)
		if uid != "" {
			fmt.Println("uid_uid_uid_uid_uid::: ", uid)
			w.WriteHeader(http.StatusOK)
			UserFromTuya = uid
		}
		service.GetDevicesFromUser(uid)
		fmt.Println("uid_uid_uid::::", uid)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

type AuthRequest struct {
	State        string `json:"state"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
}

func NewAuth() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientID := os.Getenv("TUYA_CLIENT_ID")
	clientSecret := os.Getenv("TUYA_SECRET_KEY")
	redirUr := "https://onviz-api.ru"
	domain := fmt.Sprintf("https://social.yandex.net/broker/redirect?response_type=code&client_id=%s&redirect_uri=%s", clientID, redirUr)

	// client memory store
	clientStore := store.NewClientStore()
	err := clientStore.Set(clientID, &models2.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: domain,
	})
	if err != nil {
		log.Println("Could not set client")
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
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		state := r.FormValue("state")
		redirectURI := r.FormValue("redirect_uri")
		responseType := r.FormValue("response_type")
		clientID := r.FormValue("client_id")
		scope := r.FormValue("scope")
		r.Header.Add("state", state)
		r.Header.Add("redirect_uri", redirectURI)
		r.Header.Add("response_type", responseType)
		r.Header.Add("client_id", clientID)
		r.Header.Add("scope", scope)
		log.Println("State is: ", state)
		log.Println("redirectURI is: ", redirectURI)
		log.Println("responseType is: ", responseType)
		log.Println("clientID is: ", clientID)
		log.Println("scope is: ", scope)
		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //here error
		}
	})

	http.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
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
