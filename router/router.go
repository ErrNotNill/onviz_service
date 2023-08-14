package router

import (
	"net/http"
	"onviz/LEADS"
	"onviz/VK"
	"onviz/bot_bitrix"
	"onviz/chat"
	"onviz/tests"
	"onviz/tuya"
	"onviz/yandex"
)

func Router() {

	//http.Handle("/", http.FileServer(http.Dir("./chat/public")))

	http.HandleFunc("/yandex", yandex.Alice)
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
}
