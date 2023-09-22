package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type NewTask struct {
	Title         string `json:"title"`
	CreatedBy     string `json:"createdBy"`
	ResponsibleId string `json:"responsible_id"`
}

func GetListOfLines(w http.ResponseWriter, r *http.Request) {
	uriBitrixWebHook := os.Getenv("BITRIX_WEBHOOK_OL")
	getList, err := http.Get(uriBitrixWebHook)
	if err != nil {
		log.Println(err.Error(), "Cant get list of OpenLines in bitrix")
	}
	fmt.Println(getList.Body)
}
