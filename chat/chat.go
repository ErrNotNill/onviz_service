package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"onviz/chat/cache"
)

var upgrader = websocket.Upgrader{}

func TestChat(w http.ResponseWriter, r *http.Request) {
	/*chatMessages, err := cache.RDB.LRange(context.Background(), "chat_messages", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", chatMessages)*/

	message := r.FormValue("message")
	//upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	data := make([]string, 0)
	Reader(message, ws)

	if err != nil {
		fmt.Print("i cant execute chat.html")
	}
	cache.Clients[ws] = true
	log.Println("client connected!")
	for {
		var msg cache.ChatMessage
		err = ws.ReadJSON(&msg)
		if err != nil {
			delete(cache.Clients, ws)
		}
		data = append(data, msg.Text)
		cache.Broadcaster <- msg
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
