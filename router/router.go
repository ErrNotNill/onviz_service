package router

import (
	"net/http"
	"onviz/LEADS"
	"onviz/VK"
	"onviz/bot_bitrix"
	"onviz/chat"
	"onviz/tests"
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
	http.HandleFunc("/backend/leads_get", LEADS.GetLeadsAll)

	http.HandleFunc("/backend/bot", bot_bitrix.BotBitrix)
}
