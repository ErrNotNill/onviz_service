package yandex

import (
	"encoding/json"
	"fmt"
	"github.com/azzzak/alice"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var respNse *alice.Kit

func SimpleSkill() {
	updates := alice.ListenForWebhook("/hook")

	if updates != nil {
		respNse.Resp.Text("hello")
		updates.Loop(func(k alice.Kit) *alice.Response {
			req, resp := k.Init()
			resp.RandomText("one", "two", "three").
				Button("one", "", false).
				Button("device-open?", "", false)
			if req.IsNewSession() {
				return resp.Text("device-start?")
			}
			return resp.Text(req.OriginalUtterance())
		})
	}
}

func AuthUserFromYandexToken(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body from server:", err)
		// Handle the error appropriately
		return
	}

	// Unmarshal JSON
	var requestData map[string]interface{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		fmt.Println("Error unmarshalling JSONz:", err)
		// Handle the error appropriately
		return
	}
	fmt.Println("Body request AuthUserFromYandexToken :", requestData)

	// Your custom client ID and secret
	clientId := os.Getenv("TUYA_CLIENT_ID")
	clientSecret := os.Getenv("TUYA_SECRET_KEY")

	// Create a client store and set your custom client
	clientStore := store.NewClientStore()
	err = clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		// Set other client properties as needed
	})
	if err != nil {
		log.Println("Error setting client", err.Error())
		return
	}

	// Create a token manager
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapClientStorage(clientStore)

	// Create an OAuth2 server
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	// Set internal error handler
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	// Set response error handler
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// Perform other OAuth2 server configurations as needed

	// Use the OAuth2 server for authentication
	// ...

	// Handle other logic or return responses as needed
}

func AuthUserFomYandex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()

	clientID := os.Getenv("TUYA_CLIENT_ID")
	clientSecret := os.Getenv("TUYA_SECRET_KEY")

	idGen := strconv.Itoa(rand.Intn(999999)) //nolint
	//todo but fix this in continuous, and lint
	err := clientStore.Set(idGen, &models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: "http://localhost:9090/yandex/authorize",
	})
	if err != nil {
		log.Println("Error setting client")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
		http.Error(w, re.Error.Error(), re.StatusCode)
		return
	})

	err = srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CheckConnection(w http.ResponseWriter, r *http.Request) {
	client_id := r.URL.Query().Get("client_id")
	fmt.Println("client_id", client_id)
	bd, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read yandex")
	}
	fmt.Println("CheckConnection(bd):>", string(bd))

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	/*respNse.Resp.Text("hello")
	for {
		SimpleSkill()
	}*/
	var i interface{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body from server")
	}
	fmt.Println("body i: ", i)
	js := json.Unmarshal(body, i)
	fmt.Println("body request CheckConnection: ", js)
}
func UpdateDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var i interface{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body from server")
	}
	js := json.Unmarshal(body, i)
	fmt.Println("body request CheckConnection: ", js)

	if r.Method == "GET" {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	/*respNse.Resp.Text("hello")
	for {
		SimpleSkill()
	}*/

}
