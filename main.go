package main

import (
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
