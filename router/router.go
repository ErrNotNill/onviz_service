package router

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"io"
	"log"
	"net/http"
	"onviz/LEADS"
	"onviz/VK"
	"onviz/addons"
	"onviz/bot_bitrix"
	"onviz/chat"
	"onviz/tests"
	"onviz/tuya"
	"onviz/yandex"
)

func Router() {
	//http.HandleFunc("/authorize", tuya.GetDeviceNew)
	//http.HandleFunc("/token", tuya.GetDeviceNew)
	//http.Handle("/", http.FileServer(http.Dir("./chat/public")))

	http.HandleFunc("/devices/:device_id", tuya.GetDeviceNew)
	http.HandleFunc("/yandex", yandex.Alice)
	http.HandleFunc("/v1.0", yandex.CheckConnectionYandex)
	http.HandleFunc("/get_auth_token", GetAuthTokenYandex)

	//http.HandleFunc("/", LEADS.TestStatus)
	http.HandleFunc("/chat", chat.TestChat)
	//http.HandleFunc("/tilda", TildaWebHooks)
	//http.HandleFunc("/getListOfLines", GetListOfLines)
	http.HandleFunc("/callback", VK.CallBack)
	//http.HandleFunc("/parse", testHandleFunc)
	http.HandleFunc("/test", tests.NewTestHandleFunc)
	http.HandleFunc("/check", tests.NewTestHandleFunc)

	http.HandleFunc("/leads", LEADS.LeadsAdd)
	http.HandleFunc("/leads_list", LEADS.GetLeads)
	http.HandleFunc("/dealer_deal", LEADS.DealerDealAdded)
	http.HandleFunc("/leads_get", LEADS.GetLeadsAll)

	http.HandleFunc("/bot", bot_bitrix.BotBitrix)
	http.HandleFunc("/auth", bot_bitrix.CallbackHandler)
	http.HandleFunc("/redir", bot_bitrix.RedirectHandler)

	http.HandleFunc("/auth_tuya", tuya.AuthHandler)
	http.HandleFunc("/get_token", tuya.GetTokenHandler)
	http.HandleFunc("/refresh_token", tuya.RefreshTokenHandler)

	http.HandleFunc("/text_collect", chat.GetTextCollectHandler)
	http.HandleFunc("/iframe", addons.IframeHandler)

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
	//body := []byte(``)
	req, _ := http.NewRequest(method, `http://localhost:9090/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read`, nil)
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
