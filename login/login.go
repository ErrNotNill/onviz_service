package login

import (
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"io"
	"log"
	"net/http"
	"onviz/DB"
)

func CreateAccount(login, password string) {
	fmt.Println("db connected")
	result, err := DB.Db.Exec(`insert into Users (Login, Password ) values (?, ?)`,
		login, password)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var userData UserData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	// Process user registration data (userData) as needed
	fmt.Printf("Received registration data: %+v\n", userData)
	// You can now handle the registration logic, such as storing the data in a database
	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Registration successful"})
}

func Auth() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
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
	clientID := r.FormValue("client_id")
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
