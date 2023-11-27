package router

import (
	"github.com/rs/cors"
	"net/http"
	"onviz/addons"
	"onviz/chat"
	repository2 "onviz/internal/repository"
	"onviz/service/VK"
	bot_bitrix2 "onviz/service/bitrix/bot_bitrix"
	"onviz/service/bitrix/repository"
	tuya2 "onviz/service/tuya"
	yandex2 "onviz/service/yandex"
)

func Router() {
	//http.HandleFunc("/authorize", tuya.GetDeviceNew)
	//http.HandleFunc("/token", tuya.GetDeviceNew)
	//http.Handle("/", http.FileServer(http.Dir("./chat/public")))
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this to the specific origin of your Vue.js app in a production environment.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	http.Handle("/auth_page", c.Handler(http.HandlerFunc(repository2.AuthPage)))
	http.Handle("/login_page", c.Handler(http.HandlerFunc(repository2.LoginPage)))

	//http.HandleFunc("/devices/:device_id", tuya2.GetDeviceNew)
	http.HandleFunc("/yandex/v1.0", yandex2.CheckConnection)
	http.HandleFunc("/yandex/v1.0/user/devices", tuya2.GetDeviceNew)
	http.HandleFunc("/yandex/v1.0/user/unlink", tuya2.UnlinkUser)
	http.HandleFunc("/yandex/v1.0/user/devices/query", tuya2.GetDevicesStatus)
	http.HandleFunc("/yandex/v1.0/user/devices/action", tuya2.GetDevicesStatusChanged)

	http.HandleFunc("/v1.0", yandex2.CheckConnectionYandex)
	http.HandleFunc("/get_auth_token", repository2.GetAuthTokenYandex)
	http.HandleFunc("/refresh_token", tuya2.RefreshToken)

	//http.HandleFunc("/", LEADS.TestStatus)
	http.HandleFunc("/chat", chat.TestChat)
	//http.HandleFunc("/tilda", TildaWebHooks)
	//http.HandleFunc("/getListOfLines", GetListOfLines)
	http.HandleFunc("/callback", VK.CallBack)
	//http.HandleFunc("/parse", testHandleFunc)

	http.HandleFunc("/leads", repository.LeadsAdd)
	http.HandleFunc("/leads_list", repository.GetLeads)
	http.HandleFunc("/dealer_deal", repository.DealerDealAdded)
	http.HandleFunc("/leads_get", repository.GetLeadsAll)

	http.HandleFunc("/bot", bot_bitrix2.BotBitrix)
	//http.HandleFunc("/auth", bot_bitrix.CallbackHandler)
	http.HandleFunc("/redir", bot_bitrix2.RedirectHandler)

	http.HandleFunc("/auth_tuya", tuya2.AuthHandler)
	http.HandleFunc("/get_token", tuya2.GetTokenHandler)
	//http.HandleFunc("/refresh_token", tuya.RefreshTokenHandler)

	http.HandleFunc("/text_collect", chat.GetTextCollectHandler)
	http.HandleFunc("/iframe", addons.IframeHandler)

}
