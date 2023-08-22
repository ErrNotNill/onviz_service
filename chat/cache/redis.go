package cache

import (
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
)

type ChatMessage struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

var Clients = make(map[*websocket.Conn]bool)
var Broadcaster = make(chan ChatMessage)
