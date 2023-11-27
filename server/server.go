package server

import (
	"onviz/chat/models"
	"sync"
)

type Server struct {
	Ip        string
	Port      int
	Users     map[string]*models.User
	UserMutex sync.RWMutex
	msgC      chan string
}
