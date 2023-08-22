package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"onviz/chat/cache"
)

var upgrader = websocket.Upgrader{}

var wsConn *websocket.Conn

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
	cache.Clients[ws] = true
	for {
		var msg cache.ChatMessage
		err = ws.ReadJSON(&msg)
		if err != nil {
			delete(cache.Clients, ws)
			continue
		}
		cache.Broadcaster <- msg
		fmt.Println("msg: ", msg)
		fmt.Println(<-cache.Broadcaster)
		sendMessage("Hello, client")
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
