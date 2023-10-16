package tuya

import (
	"context"
	"fmt"
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
	go r.Run(":9090")
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

func GetDevicesStatusChanged(w http.ResponseWriter, r *http.Request) {
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

func GetDevicesStatus(w http.ResponseWriter, r *http.Request) {
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
