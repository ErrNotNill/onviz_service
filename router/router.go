package router

import (
	"net/http"
	"onviz/LEADS"
	"onviz/VK"
	"onviz/tests"
)

func Router() {

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
}
