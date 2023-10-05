package router

import (
	"net/http"
	"onviz/LEADS"
	"onviz/VK"
	"onviz/addons"
	"onviz/bot_bitrix"
	"onviz/chat"
	"onviz/login"
	"onviz/tests"
	"onviz/tuya"
	"onviz/yandex"
)

func Router() {
	//http.HandleFunc("/authorize", tuya.GetDeviceNew)
	//http.HandleFunc("/token", tuya.GetDeviceNew)
	//http.Handle("/", http.FileServer(http.Dir("./chat/public")))

	http.HandleFunc("/auth_page", login.AuthPage)

	http.HandleFunc("/devices/:device_id", tuya.GetDeviceNew)
	http.HandleFunc("/yandex", yandex.Alice)
	http.HandleFunc("/v1.0", yandex.CheckConnectionYandex)
	http.HandleFunc("/get_auth_token", login.GetAuthTokenYandex)
	http.HandleFunc("/refresh_token", tuya.RefreshToken)

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
	//http.HandleFunc("/refresh_token", tuya.RefreshTokenHandler)

	http.HandleFunc("/text_collect", chat.GetTextCollectHandler)
	http.HandleFunc("/iframe", addons.IframeHandler)

}
