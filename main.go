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
	"onviz/internal/repository"
	"onviz/router"
	"onviz/service/tuya/service"
	"os"
	"path/filepath"
)

func main() {
	//queue.MqttInit()
	repository.Auth() //not used now

	fmt.Println("good luck ^_^")

	//load .env file
	if err := godotenv.Load(filepath.Join(".env")); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}

	//init router
	router.Router()
	fmt.Println("Starting")

	//init mysql
	urlMysql := os.Getenv("URL_MYSQL")
	err := DB.InitDB(urlMysql)
	if err != nil {
		fmt.Println("cant' connect to mysql")
		log.Fatal(err)
	} else {
		fmt.Println("db init accepted")
	}

	//init Tuya api
	service.TheTuyaAllFunctions()

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	cache.RDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

	ping := cache.RDB.Ping(context.Background())
	fmt.Println("redis started", ping)

	fmt.Println("Server started")
	fmt.Println("http://localhost:9090")
	err = http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println("Server started with error")
		return
	}

}
