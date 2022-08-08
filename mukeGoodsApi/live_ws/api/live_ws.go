package api

import (
	"fmt"
	"log"
	"webApi/live_ws/api/handler"
	"webApi/live_ws/forms"
	"webApi/live_ws/global"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

var Msg = forms.Msg{}

// login
func WsHandler(c *gin.Context) {
	// 升级为websocket
	var err error
	global.WsConn, err = global.WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// userC := &handler.Clients{Fc: &forms.Client{Id: fmt.Sprintf("%v", uuid.NewV4()), Ws: global.WsConn, UnSc: make(chan []byte)}}

	// wsIn := WsPushService{}
	// wsIn.WsPush = &handler.Clients{Fc: &forms.Client{Id: fmt.Sprintf("%v", uuid.NewV4()), Ws: global.WsConn, UnSc: make(chan []byte)}}

	userC := &forms.Client{Id: fmt.Sprintf("%v", uuid.NewV4()), Ws: global.WsConn, UnSc: make(chan []byte)}
	WsPush := handler.NewMsPush(&forms.Client{Id: fmt.Sprintf("%v", uuid.NewV4()), Ws: global.WsConn, UnSc: make(chan []byte)})

	defer func() {
		global.WsClients[userC] = false
		zap.S().Infof("用户池：%v", global.WsClients)
		global.WsConn.Close()
	}()
	// 写入连接广播
	global.WsClients[userC] = true

	// 监控写消息
	go func() {
		for {
			select {
			case <-global.Ping:
				WsPush.WsPing()
			case <-global.PushMsg:
				WsPush.WsPushMsg(Msg)
			}
		}
	}()

	// 监控client发送的 json 消息
	// Ping消息体 { action: "Ping", data: {id: 123}}
	// Ping消息体 controller: "Push", action: "Comment", data: {content: "你好",}

	for {
		if err := global.WsConn.ReadJSON(&Msg); err != nil {
			if !websocket.IsCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseNormalClosure) {
				zap.S().Fatalf("unexpected read error%v", err)
			}
			break
		}
		// 调用
		WsToAction(Msg.Action)
	}
}

// 转发
func WsToAction(action string) {
	// zap.S().Infof("转发 %s", action)
	switch action {
	case "Ping":
		global.Ping <- struct{}{}
	case "Comment":
		global.PushMsg <- struct{}{}
	}
}
