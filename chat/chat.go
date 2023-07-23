package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func TestChat(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func TestChatOld(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(message))
		log.Printf("Received: %s", message)
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}

	// get client ip address
	clientIp := r.Header.Get("X-Real-Ip")
	if clientIp != "" {
		fmt.Fprintf(w, "Hi, clientIp %s!\n", clientIp)
	}
	fmt.Fprintf(w, "Hello, remoteaddr %s!\n", r.RemoteAddr)

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	fmt.Fprintf(w, "Hello 2nd remoteaddr %s!\n", r.RemoteAddr)
	// print out the ip address
	fmt.Fprintf(w, ip+"\n\n")

	// sometimes, the user acccess the web server via a proxy or load balancer.
	// The above IP address will be the IP address of the proxy or load balancer and not the user's machine.

	// let's get the request HTTP header "X-Forwarded-For (XFF)"
	// if the value returned is not null, then this is the real IP address of the user.
	fmt.Fprintf(w, "X-Forwarded-For :"+r.Header.Get("X-FORWARDED-FOR"))
}

/*func TestChat(w http.ResponseWriter, r *http.Request) {
	ip := getClientIpAddr(r)
	w.Write([]byte(ip))
}

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-Real-Ip")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}*/

/*func ReadUserIP(r *http.Request) string {
	fmt.Println("remote addr: ", r.RemoteAddr)
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-FORWARDED-FOR")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}*/
