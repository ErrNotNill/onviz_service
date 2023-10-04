package yandex

import (
	"net/http"
)

func CheckConnectionYandex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
