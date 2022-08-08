package forms

import "github.com/gorilla/websocket"

// read msg
type Msg struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

// 用户client
type Client struct {
	Id   string
	Ws   *websocket.Conn
	UnSc chan []byte
}
