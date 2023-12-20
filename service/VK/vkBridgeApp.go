package VK

import (
	"fmt"
	"github.com/go-vk-api/vk"
	"log"
	"os"
)

func StartVkBridge() {
	accessToken := os.Getenv("VK_BRIDGE_ACCESS_TOKEN")
	client, err := vk.NewClientWithOptions(
		vk.WithToken(accessToken),
	)
	if err != nil {
		fmt.Println("panic after client auth")
		panic(err)
	}
	// Make the "users.get" API call
	methodName := "users.get"
	params := map[string]interface{}{
		"user_ids": "258098783",            // Replace "1" with the VK user ID you want to retrieve.
		"fields":   "first_name,last_name", // Replace with the desired fields.
	}
	response := client.CallMethod(methodName, params, nil)
	if err != nil {
		fmt.Println("err response")
		response.Error()
	}
	fmt.Println("OK")

	err = printMe(client)
	if err != nil {
		fmt.Println("printMe error: ", err)
	}
}

func printMe(api *vk.Client) error {
	var user interface{}
	err := api.CallMethod("users.get", vk.RequestParams{}, &user)
	if err != nil {
		fmt.Println("Error getting users: ", err)
	}
	me := user
	log.Println(me)
	return nil
}
