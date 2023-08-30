package addons

import (
	"fmt"
	"html/template"
	"net/http"
)

func IframeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Iframe Handler")
	ts, err := template.ParseFiles("./addons/iframe.html")
	if err != nil {
		fmt.Print("cant parse html")
	}
	fmt.Println("before executing")
	err = ts.Execute(w, r)
	if err != nil {
		fmt.Println("error executing")
		return
	}
}
