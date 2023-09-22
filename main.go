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
	"onviz/tuya"
	"os"
)

func main() {

	//queue.MqttInit()
	//VK.StartVkBridge()

	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}

	TheTuyaAllFunctions()

	//VK.StartVkBridge()
	fmt.Println("Starting")

	urlDb := os.Getenv("URL_MYSQL")
	err := DB.InitDB(urlDb)
	if err != nil {
		fmt.Println("cant' connect to mysql")
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
	//cache.RDB.RPush(context.Background(), "chat_messages", messageType)

	fmt.Println("Server started")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Server started with error")
		return
	}
	//go chat.WsStart()
}

func TheTuyaAllFunctions() {
	//tuya.Cfg()

	tuya.GetToken()
	tuya.RefreshToken()
	tuya.GetDevicesInProject()
	//tuya.PolicyAction()

	//tuya.GetUsers()
	/*devices, err := tuya.GetDevicesList()
	if err != nil {
		log.Println("No devices")
	}
	for _, device := range devices {
		fmt.Printf("ID: %v, Name: %v, Online: %v\n", device.Result, device.Success, device.T)
	}*/
	tuya.GetDevice("bf85de23e4cf1c10fb6bsn")
}
