package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conn *websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	// 升级为websocket
	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// json 消息
	for {
		Msg := make(map[string]interface{})
		conn.ReadJSON(&Msg)
		log.Println(Msg)
	}
	// message 消息
	// for {
	// 	messageType, p, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Println(err)
	// 		break
	// 	}
	// 	log.Println(string(p))
	// 	if err := conn.WriteMessage(messageType, []byte("send msg")); err != nil {
	// 		log.Println(err)
	// 		break
	// 	}
	// }
}

func main() {
	// http.Handle("/foo", handler)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8051", nil))

}
