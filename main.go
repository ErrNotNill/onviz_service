package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"onviz/DB"
	"onviz/chat/cache"
	"onviz/router"
	"os"
)

var linkToRemoteServerUsage = "http://45.141.79.120/getListOfLines"

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}
	/*
		const (
			clientID  = "9x8wfym7m5vyck7tdwwt&"
			secretKey = "d8205ed66f15471fa969aecab48ab495"
			baseURL   = "https://openapi.tuyaeu.com"
			endpoint  = "/v1.0/token?"
			grantType = "grant_type=1"
		)
	*/

	//tuya.TuyaStart() //blocked tuya
	//VK.StartVkBridge()

	// Make the "users.get" API call and handle the response here..

	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}

	//VK.StartVkBridge()
	fmt.Println("Starting")

	urlDb := os.Getenv("URL_MYSQL")
	err := DB.InitDB(urlDb)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db init accepted")
	}

	router.Router()

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")
	cache.RDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
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
