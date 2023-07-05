package tests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CheckAvialable(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewTestHandleFunc(w http.ResponseWriter, r *http.Request) {
	url := "https://onviz.bitrix24.site/"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	fmt.Println(responseString)
}
