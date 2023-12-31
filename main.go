package main

import (
	"context"
	"fmt"
	"github.com/azzzak/alice"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"onviz/DB"
	"onviz/chat/cache"
	"onviz/router"
	"onviz/service/tuya/service"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//queue.MqttInit()

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

	server := &http.Server{
		Addr:              ":9090",
		ReadHeaderTimeout: 3 * time.Second,
	}

	updates := alice.ListenForWebhook("/api/hook")

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server started with error")
		panic(err)
	}
	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("привет")
		}
		return resp.Text(req.OriginalUtterance())
	})
}
