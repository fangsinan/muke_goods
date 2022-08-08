package api

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"webApi/live_ws/global"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// 	uuid "github.com/satori/go.uuid"
// 	"go.uber.org/zap"
// )

// // read msg
// var Msg struct {
// 	Controller string      `json:"controller"`
// 	Action     string      `json:"action"`
// 	Data       interface{} `json:"data"`
// }

// type Client struct {
// 	id   string
// 	ws   *websocket.Conn
// 	unSc chan []byte
// }

// var (
// 	Ping      = make(chan struct{}, 1)
// 	PushMsg   = make(chan struct{}, 1)
// 	WsClients = make(map[*Client]bool)
// )

// // login
// func WsHandler(c *gin.Context) {
// 	// 升级为websocket
// 	var err error
// 	global.WsConn, err = global.WsUpgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	userC := &Client{id: fmt.Sprintf("%v", uuid.NewV4()), ws: global.WsConn, unSc: make(chan []byte)}
// 	defer func() {
// 		WsClients[userC] = false
// 		zap.S().Infof("用户池：%v", WsClients)
// 		global.WsConn.Close()
// 	}()
// 	// 写入连接广播
// 	WsClients[userC] = true

// 	// 监控写消息
// 	go func() {
// 		for {
// 			select {
// 			case <-Ping:
// 				WsPing()
// 			case <-PushMsg:
// 				userC.WsPushMsg()
// 			}
// 		}
// 	}()

// 	// 监控client发送的 json 消息
// 	for {
// 		// {controller: "Index", action: "Ping", data: {id: 778,com: "com"}}
// 		if err := global.WsConn.ReadJSON(&Msg); err != nil {
// 			if !websocket.IsCloseError(err,
// 				websocket.CloseGoingAway,
// 				websocket.CloseNormalClosure) {
// 				zap.S().Fatalf("unexpected read error%v", err)
// 			}
// 			break
// 		}
// 		// zap.S().Infof("msg：%v", Msg)
// 		// 调用
// 		WsToAction(Msg.Action)
// 	}
// }

// // 转发
// func WsToAction(action string) {
// 	// zap.S().Infof("转发 %s", action)
// 	switch action {
// 	case "Ping":
// 		Ping <- struct{}{}
// 	case "Comment":
// 		PushMsg <- struct{}{}
// 	}
// }

// func WsPing() {
// 	// zap.S().Infof("show WsPing")

// 	// <-Ping
// 	PingStruct := struct {
// 		Controller string         `json:"controller"`
// 		Action     string         `json:"action"`
// 		Data       map[string]int `json:"data"`
// 	}{
// 		Controller: "Index",
// 		Action:     "Ping",
// 		Data: map[string]int{
// 			"live_id": 19,
// 		},
// 	}

// 	wsErr := global.WsConn.WriteJSON(PingStruct)
// 	if wsErr != nil {
// 		log.Println(wsErr)
// 	}

// }

// func (uc *Client) WsPushMsg() {
// 	comment := struct {
// 		AccessUserToken string  `json:"accessUserToken"`
// 		Content         string  `json:"content"`
// 		Live_id         float64 `json:"live_id"`
// 		User_id         float64 `json:"user_id"`
// 	}{}
// 	// 解析评论结构
// 	b, err := json.Marshal(&Msg.Data)
// 	if err != nil {
// 		zap.S().Errorf("cannot json Marshal:%v", err)
// 	}

// 	err = json.Unmarshal(b, &comment)
// 	if err != nil {
// 		zap.S().Errorf("cannot UnMarshal comment:%v", err)
// 	}

// 	PingStruct := struct {
// 		Controller string                 `json:"controller"`
// 		Action     string                 `json:"action"`
// 		Data       map[string]interface{} `json:"data"`
// 	}{
// 		Controller: "WsPushMsg",
// 		Action:     "WsPushMsg",
// 		Data: map[string]interface{}{
// 			"status":  200,
// 			"lid":     comment.Live_id,
// 			"uid":     comment.User_id,
// 			"comment": comment.Content,
// 		},
// 	}
// 	zap.S().Infof("PingStruct:%v", PingStruct)
// 	zap.S().Infof("WsClients:%v", WsClients)

// 	// 消息广播
// 	for client, v := range WsClients {
// 		if v {
// 			if client.id == uc.id {
// 				break
// 			}
// 			wsErr := client.ws.WriteJSON(PingStruct)
// 			if wsErr != nil {
// 				log.Println(wsErr)
// 				break
// 			}
// 		}
// 	}

// }
