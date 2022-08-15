package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"sync"
	"time"
	"webApi/ws_test/push"
)

var Up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UserClient struct {
	Conn    *websocket.Conn
	IsClose bool
	CID     string
}

var (
	conn *websocket.Conn
	Ping = make(chan *websocket.Conn, 1024)
)

type Push interface {
	Ping(*websocket.Conn)
}
type PushService struct {
	Push Push
	conn *websocket.Conn
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("发起socket请求")
	// 升级为websocket
	var err error
	conn, err = Up.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	uc := &UserClient{
		Conn:    conn,
		IsClose: true,
		CID:     fmt.Sprintf("%v", uuid.NewV4()),
	}
	defer func() {
		uc.IsClose = false
		conn.Close()
		log.Println("关闭链接")
	}()
	fmt.Println(uc)
	// 接受 消息
	go Read(uc)
	// 推送消息
	go Writer(1)
	for {

	}
}

func Writer(liveID int) {

	p := &PushService{
		Push: &push.Push{
			Conn: conn,
			Live: liveID,
			Mux:  &sync.Mutex{},
		},
		conn: conn,
	}

	for {
		select {
		case ping := <-Ping:
			log.Println("chan 接收 ")
			p.Push.Ping(ping)
		case <-time.After(10 * time.Second):
			log.Println("10秒后关闭socket链接")
			conn.Close()

		}
	}

}

func Read(u *UserClient) {
	Msg := struct {
		Action string      `json:"action"`
		Data   interface{} `json:"data"`
	}{}
	for {
		if !u.IsClose { // 客户端断开链接后无需继续任务
			log.Println("校验关闭...")
			break
		}
		_, b, err := u.Conn.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseNormalClosure) {
				log.Printf("unexpected read error %v", err)

			}
			break
		}
		err = json.Unmarshal(b, &Msg)
		if err != nil {
			log.Fatalf("unexpected read unmarshal error %v", err)
			return
		}
		// 调用
		log.Println("chan 发送 ")
		Ping <- conn
	}
}

func main() {
	addr := ":8025"
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
