package tilda

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TildaWebHooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	w.WriteHeader(http.StatusOK)
	webhookData := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("got webhook payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s : %v\n", k, v)
	}
	switch name := webhookData["username"]; name {
	case "test":
		fmt.Println("Invoked")
	default:
		fmt.Println("Not invoked")
	}
}
