package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"onviz/DB"
	"onviz/chat/cache"
	"onviz/router"
)

var linkToRemoteServerUsage = "http://45.141.79.120/getListOfLines"

func main() {

	err := DB.InitDB()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db init accepted")
	}

	router.Router()

	cache.RDB = redis.NewClient(&redis.Options{
		Addr:     "45.141.79.120:6379",
		Password: "redis",
		DB:       0,
	})
	ping := cache.RDB.Ping(context.Background())
	fmt.Println("redis started", ping)
	//messageType := "access"
	//cache.RDB.RPush(context.Background(), "chat_messages", messageType)

	fmt.Println("Server started")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Server started with error")
		return
	}
	//go chat.WsStart()
}

func taskAdd() {

}

func GetListOfLines(w http.ResponseWriter, r *http.Request) {
	getList, err := http.Get("https://onviz.bitrix24.ru/rest/13938/6rh8x17zqjx2sb9x/imopenlines.config.list.get")
	if err != nil {
		log.Println(err.Error(), "Cant get list of OpenLines in bitrix")
	}
	fmt.Println(getList.Body)
}
