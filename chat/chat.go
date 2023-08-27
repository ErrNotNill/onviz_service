package chat

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"onviz/chat/cache"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var wsConn *websocket.Conn

type Message struct {
	Greeting string `json:"greeting"`
}

func TestChat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	/*chatMessages, err := cache.RDB.LRange(context.Background(), "chat_messages", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", chatMessages)*/

	//upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	//cache.Clients[ws] = true
	//ws.SetReadLimit(86400)
	for {
		//fmt.Println(r.Body)
		//var msg cache.ChatMessage
		var msg Message

		err = ws.ReadJSON(&msg)
		if err != nil {
			cache.RDB.RPush(context.Background(), "chat_messages", msg)
			//	cache.ChatMessage{Message: msg.Greeting}
			fmt.Println("i can't read message")
			//delete(cache.Clients, ws)
			break
		}
		fmt.Println("not errors,", "message:", msg)
		//fmt.Println(msg.Greeting)
		//cache.Broadcaster <- msg
		//fmt.Println("msg: ", msg)
		//fmt.Println(<-cache.Broadcaster)
		messages := make([]Message, 0)
		messages = append(messages, msg)
		err = ws.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(rand.Intn(12313123123))))
		if err != nil {
			//todo correct redis adding
			cache.RDB.RPush(context.Background(), "chat_messages", msg)
			log.Println("write:", err)
			break
		}

	}
}

func sendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Println("msg: ", err)
	}
}

/*func HandleMessages() {
	for {

		// grab any next message from channel
		msg := <-cache.Broadcaster
		for client := range cache.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(cache.Clients, client)
			}
			// send previous messages
			for _, chatMessage := range cache.ChatMessage {
				var msg ChatMessage
				json.Unmarshal([]byte(chatMessage), &msg)
				err := client.WriteJSON(msg)
				if err != nil && unsafeError(err) {
					log.Printf("error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}

			cache.RDB.RPush(context.Background(), "chat_messages", json)
		}
	}
	// send previous messages

}*/
