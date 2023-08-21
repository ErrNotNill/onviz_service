package chat

import (
	"fmt"
	"io"
	"net/http"
)

func GetTextCollectHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	reader, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("Error reading")
	}
	fmt.Println(string(reader))
}
