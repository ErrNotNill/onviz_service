package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"onviz/chat/cache"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
}

var wsConn *websocket.Conn

type Message struct {
	Greeting string `json:"greeting"`
}

var clients = make(map[*websocket.Conn]*Client)

func TestChat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	//upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	client := &Client{conn: ws}
	clients[ws] = client
	//cache.Clients[ws] = true
	//ws.SetReadLimit(86400)
	chatMessages, err := cache.RDB.LRange(context.Background(), "chat_messages", -1, -1).Result()
	if err != nil {
		panic(err)
	}
	for _, message := range chatMessages {
		err := ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Failed to send chat history:", err)
			break
		}
	}

	for {
		var msg Message
		err = ws.ReadJSON(&msg)
		if err != nil {
			// Handle disconnection or errors
			delete(clients, ws)
			break
		}

		// Broadcast the incoming message to all connected clients
		broadcastMessage(msg)
	}

}

func broadcastMessage(msg Message) {
	// Store the message in the database or cache
	status := cache.RDB.RPush(context.Background(), "chat_messages", msg.Greeting)

	if status.Err() != nil {
		// Handle the error, such as logging it or returning an error response.
		log.Printf("Failed to push the message to Redis: %v", status.Err())
		// You may also want to inform the client that their message was not saved.
	} else {
		// The message was successfully pushed to the Redis list.
	}

	// Convert the message to JSON
	messages := []Message{msg}
	js, err := json.Marshal(&messages)
	if err != nil {
		fmt.Println("Error marshalling messages")
		return
	}

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, js)
		if err != nil {
			// Handle disconnection or errors
			delete(clients, client)
		}
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
