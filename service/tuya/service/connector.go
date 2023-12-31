package service

import (
	"context"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/tuya/tuya-connector-go/connector"
	"github.com/tuya/tuya-connector-go/connector/constant"
	"github.com/tuya/tuya-connector-go/connector/env"
	"github.com/tuya/tuya-connector-go/connector/env/extension"
	"github.com/tuya/tuya-connector-go/connector/logger"
	"github.com/tuya/tuya-connector-go/example/messaging"
	"github.com/tuya/tuya-connector-go/example/router"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Custom configuration
func InitConnector() {
	// Custom configuration
	connector.InitWithOptions(env.WithApiHost("https://openapi.tuyaeu.com"),
		env.WithMsgHost("https://openapi.tuyaeu.com"),
		env.WithAccessID("9x8wfym7m5vyck7tdwwt"),
		env.WithAccessKey("d8205ed66f15471fa969aecab48ab495"))
	// Start the service
	go messaging.Listener()
	r := router.NewGinEngin()
	go func() {
		err := r.Run(":9090")
		if err != nil {
			log.Println("Error run connector", err.Error())
		}
	}()
	watitSignal()
}

func watitSignal() {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		select {
		case c := <-quitCh:
			extension.GetMessage(constant.TUYA_MESSAGE).Stop()
			logger.Log.Infof("receive sig:%v, shutdown the http server...", c.String())
			return
		}
	}
}

func SendAccessTokenForYandex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bs, _ := io.ReadAll(r.Body)
	log.Println("response:::", string(bs))
	reqId := fmt.Sprintf(`{
  "token": %v,
}`, AccessToken)
	w.Write([]byte(reqId))
}

func SendRefreshTokenForYandex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bs, _ := io.ReadAll(r.Body)
	log.Println("response:::", string(bs))
	reqId := fmt.Sprintf(`{
  "request_id": %v,
}`, RefreshTokenVal)
	w.Write([]byte(reqId))
}

func GetDevicesInfo(w http.ResponseWriter, r *http.Request) {

	urlQueries := r.URL.Query().Encode()
	fmt.Println("urlQueries:>", urlQueries)

	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing")
	}

	var filter interface{}
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		log.Println("Error decoding filter")
	}

	// Do something with filter
	fmt.Printf("%+v", filter)

	bd, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read yandex")
	}
	fmt.Println("GetDevicesInfo(bd):>", string(bd))

	var requestId int
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		bs, _ := io.ReadAll(r.Body)
		log.Println("response:::", string(bs))
		reqId := fmt.Sprintf(`{
  "request_id": %v,
}`, requestId)
		w.Write([]byte(reqId))
	}
}

func GetDevicesState(w http.ResponseWriter, r *http.Request) {
	bd, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read yandex")
	}
	fmt.Println("GetDevicesState(bd):>", string(bd))

	var requestId int
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bs, _ := io.ReadAll(r.Body)
		log.Println("response:::", string(bs))
		reqId := fmt.Sprintf(`{
  "request_id": %v,
}`, requestId)
		w.Write([]byte(reqId))
	}
}
func GetDevicesStatusChanged(w http.ResponseWriter, r *http.Request) {
	bd, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read yandex")
	}
	fmt.Println("GetDevicesStatusChanged(bd):>", string(bd))

	var requestId int
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bs, _ := io.ReadAll(r.Body)
		log.Println("response:::", string(bs))
		reqId := fmt.Sprintf(`{
  "request_id": %v,
}`, requestId)
		w.Write([]byte(reqId))
	}
}

func UnlinkUser(w http.ResponseWriter, r *http.Request) {
	bd, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read yandex")
	}
	fmt.Println("UnlinkUser(bd):>", string(bd))

	var requestId int
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bs, _ := io.ReadAll(r.Body)
		log.Println("response:::", string(bs))
		reqId := fmt.Sprintf(`{
  "request_id": %v,
}`, requestId)
		w.Write([]byte(reqId))
	}
}

func GetDeviceNew(w http.ResponseWriter, r *http.Request) {
	resp := &GetDeviceResponse{}
	// Initiate an API request
	err := connector.MakeGetRequest(
		context.Background(),
		connector.WithAPIUri(fmt.Sprintf("/v1.0/devices/%s", DeviceID)),
		connector.WithResp(resp))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	bs, _ := io.ReadAll(r.Body)
	log.Println("response:::", string(bs))
}

// Data structure returned by OpenAPI
type GetDeviceResponseNew struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	T       int64       `json:"t"`
}
