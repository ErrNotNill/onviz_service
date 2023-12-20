package repository

import (
	"io"
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
	req, err := http.NewRequest(http.MethodGet, uriBitrixWebHook, nil)
	if err != nil {
		log.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error default client:", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing:", err.Error())
		}
	}(res.Body)
}
