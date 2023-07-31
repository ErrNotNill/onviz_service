package yandex

import (
	"github.com/azzzak/alice"
	"net/http"
)

var respNse *alice.Kit

func SimpleSkill() {
	updates := alice.ListenForWebhook("/hook")
	go http.ListenAndServe(":3000", nil)

	respNse.Resp.Text("hello")
	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		resp.RandomText("one", "two", "three").
			Button("one", "", false).
			Button("отстань", "", false)
		if req.IsNewSession() {
			return resp.Text("привет")
		}
		return resp.Text(req.OriginalUtterance())
	})

}

func Alice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	respNse.Resp.Text("hello")
	for {
		SimpleSkill()
	}
}
