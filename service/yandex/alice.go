package yandex

import (
	"encoding/json"
	"fmt"
	"github.com/azzzak/alice"
	"io"
	"net/http"
)

var respNse *alice.Kit

func SimpleSkill() {
	updates := alice.ListenForWebhook("/hook")

	if updates != nil {
		respNse.Resp.Text("hello")
		updates.Loop(func(k alice.Kit) *alice.Response {
			req, resp := k.Init()
			resp.RandomText("one", "two", "three").
				Button("one", "", false).
				Button("device-open?", "", false)
			if req.IsNewSession() {
				return resp.Text("device-start?")
			}
			return resp.Text(req.OriginalUtterance())
		})
	}
}

func CheckConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	/*respNse.Resp.Text("hello")
	for {
		SimpleSkill()
	}*/
	var i interface{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body from server")
	}
	js := json.Unmarshal(body, i)
	fmt.Println("body request: ", js)
}
func UpdateDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var i interface{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body from server")
	}
	js := json.Unmarshal(body, i)
	fmt.Println("body request: ", js)

	if r.Method == "GET" {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	/*respNse.Resp.Text("hello")
	for {
		SimpleSkill()
	}*/

}
