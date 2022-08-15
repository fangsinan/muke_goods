package forms

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Msg read msg
type Msg struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

// Client 用户client
type Client struct {
	Id      string
	Ws      *websocket.Conn
	IsValid bool
	Mux     *sync.Mutex
}
