package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"onviz/tuya"
)

// var linkToRemoteServerUsage = "http://45.141.79.120/getListOfLines"

func main() {
	/*
		const (
			clientID  = "9x8wfym7m5vyck7tdwwt&"
			secretKey = "d8205ed66f15471fa969aecab48ab495"
			baseURL   = "https://openapi.tuyaeu.com"
			endpoint  = "/v1.0/token?"
			grantType = "grant_type=1"
		)
	*/

	//VK.StartVkBridge()

	// Make the "users.get" API call and handle the response here..

	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}

	tuya.GetToken()

	//tuya.PolicyAction()

	//tuya.GetUsers()
	devices, err := tuya.GetDevicesList()
	if err != nil {
		log.Println("No devices")
	}
	for _, device := range devices {
		fmt.Printf("ID: %v, Name: %v, Online: %v\n", device.Result, device.Success, device.T)
	}
	tuya.RefreshToken(tuya.ClientID, tuya.RefreshTokenVal)

	//tuya.GetDeviceList()
	//tuya.GetDevicesWithToken()

	//todo get device from ID's list
	//tuya.GetDevice(tuya.DeviceID)
	//tuya.GetDevice("bf85de23e4cf1c10fb6bsn")

	//VK.StartVkBridge()
	/*fmt.Println("Starting")

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
	//messageType := "access"
	//cache.RDB.RPush(context.Background(), "chat_messages", messageType)

	fmt.Println("Server started")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Server started with error")
		return
	}
	//go chat.WsStart()*/
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
