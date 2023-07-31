package chat

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"onviz/chat/cache"
)

func Reader(message string, conn *websocket.Conn) {

	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println("P: ", string(p))

		cache.RDB.RPush(context.Background(), "chat_messages", messageType)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte(message)); err != nil {
			cache.RDB.RPush(context.Background(), "chat_messages", messageType)
			log.Println(err)
			return
		}
	}
}
