package router

import (
	"github.com/rs/cors"
	"net/http"
	"onviz/addons"
	"onviz/bot_bitrix"
	"onviz/chat"
	"onviz/login"
	"onviz/repository/bitrix"
	"onviz/service/VK"
	tuya2 "onviz/service/tuya"
	yandex2 "onviz/service/yandex"
	"onviz/tests"
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
	http.Handle("/auth_page", c.Handler(http.HandlerFunc(login.AuthPage)))
	http.Handle("/login_page", c.Handler(http.HandlerFunc(login.LoginPage)))

	http.HandleFunc("/devices/:device_id", tuya2.GetDeviceNew)
	http.HandleFunc("/yandex", yandex2.Alice)
	http.HandleFunc("/v1.0", yandex2.CheckConnectionYandex)
	http.HandleFunc("/get_auth_token", login.GetAuthTokenYandex)
	http.HandleFunc("/refresh_token", tuya2.RefreshToken)

	//http.HandleFunc("/", LEADS.TestStatus)
	http.HandleFunc("/chat", chat.TestChat)
	//http.HandleFunc("/tilda", TildaWebHooks)
	//http.HandleFunc("/getListOfLines", GetListOfLines)
	http.HandleFunc("/callback", VK.CallBack)
	//http.HandleFunc("/parse", testHandleFunc)
	http.HandleFunc("/test", tests.NewTestHandleFunc)
	http.HandleFunc("/check", tests.NewTestHandleFunc)

	http.HandleFunc("/leads", bitrix.LeadsAdd)
	http.HandleFunc("/leads_list", bitrix.GetLeads)
	http.HandleFunc("/dealer_deal", bitrix.DealerDealAdded)
	http.HandleFunc("/leads_get", bitrix.GetLeadsAll)

	http.HandleFunc("/bot", bot_bitrix.BotBitrix)
	http.HandleFunc("/auth", bot_bitrix.CallbackHandler)
	http.HandleFunc("/redir", bot_bitrix.RedirectHandler)

	http.HandleFunc("/auth_tuya", tuya2.AuthHandler)
	http.HandleFunc("/get_token", tuya2.GetTokenHandler)
	//http.HandleFunc("/refresh_token", tuya.RefreshTokenHandler)

	http.HandleFunc("/text_collect", chat.GetTextCollectHandler)
	http.HandleFunc("/iframe", addons.IframeHandler)

}
