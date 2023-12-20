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
	"onviz/service/tuya/service"
	yandex2 "onviz/service/yandex"
)

//todo endpoints for yandex

func Router() {

	//http.HandleFunc("/authorize", tuya.GetDeviceNew)
	//http.HandleFunc("/token", tuya.GetDeviceNew)
	//http.Handle("/", http.FileServer(http.Dir("./chat/public")))
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this to the specific origin of your Vue.js app in a production environment.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	http.Handle("/api/auth_page", c.Handler(http.HandlerFunc(repository2.AuthPage)))
	http.Handle("/api/login_page", c.Handler(http.HandlerFunc(repository2.LoginPage)))
	http.HandleFunc("/api/redirect", repository2.RedirectPage) //here user redirects from login page

	http.HandleFunc("/api/callback/code", repository2.CallbackHandler)

	//http.HandleFunc("/devices/:device_id", tuya2.GetDeviceNew)
	http.HandleFunc("/api/v1.0/", yandex2.CheckConnection)
	http.HandleFunc("/api/v1.0/user/devices", yandex2.CheckConnection) //todo get user devices GET
	//http.HandleFunc("/v1.0/user/devices/query", yandex2.InfoAboutDevicesState) //todo info about state devices POST
	//http.HandleFunc("/v1.0/user/devices/action", yandex2.ChangeDevicesState)   //todo change state of devices POST

	http.HandleFunc("/api/yandex/token", yandex2.AuthUserFromYandexToken)

	http.HandleFunc("/api/yandex/v1.0/user/devices", service.GetDeviceNew)
	http.HandleFunc("/api/yandex/v1.0/user/unlink", service.UnlinkUser)
	http.HandleFunc("/api/yandex/v1.0/user/devices/query", service.GetDevicesStatus)
	http.HandleFunc("/api/yandex/v1.0/user/devices/action", service.GetDevicesStatusChanged)
	http.HandleFunc("/api/yandex/v1.0/get_token", service.SendAccessTokenForYandex)
	http.HandleFunc("/api/yandex/v1.0/refresh_token", service.SendRefreshTokenForYandex)

	http.HandleFunc("/api/v1.0", yandex2.CheckConnectionYandex)
	http.HandleFunc("/api/get_auth_token", repository2.GetAuthTokenYandex)
	http.HandleFunc("/api/refresh_token", service.RefreshToken)

	//http.HandleFunc("/", LEADS.TestStatus)
	http.HandleFunc("/api/chat", chat.TestChat)
	//http.HandleFunc("/tilda", TildaWebHooks)
	//http.HandleFunc("/getListOfLines", GetListOfLines)
	http.HandleFunc("/api/callback", VK.CallBack)
	//http.HandleFunc("/parse", testHandleFunc)

	http.HandleFunc("/api/leads", repository.LeadsAdd)
	http.HandleFunc("/api/leads_list", repository.GetLeads)
	http.HandleFunc("/api/dealer_deal", repository.DealerDealAdded)
	http.HandleFunc("/api/leads_get", repository.GetLeadsAll)

	http.HandleFunc("/api/bot", bot_bitrix2.BotBitrix)
	//http.HandleFunc("/auth", bot_bitrix.CallbackHandler)
	http.HandleFunc("/api/redir", bot_bitrix2.RedirectHandler)

	http.HandleFunc("/api/auth_tuya", service.AuthHandler)
	http.HandleFunc("/api/get_token", service.GetTokenHandler)
	//http.HandleFunc("/refresh_token", tuya.RefreshTokenHandler)

	http.HandleFunc("/api/text_collect", chat.GetTextCollectHandler)
	http.HandleFunc("/api/iframe", addons.IframeHandler)

}
