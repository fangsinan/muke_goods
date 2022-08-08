package handler

import (
	"encoding/json"
	"log"
	"webApi/live_ws/forms"
	"webApi/live_ws/global"

	"go.uber.org/zap"
)

type Clients struct {
	Fc *forms.Client
}

func (wsC *Clients) WsPing() {
	PingStruct := struct {
		Controller string         `json:"controller"`
		Action     string         `json:"action"`
		Data       map[string]int `json:"data"`
	}{
		Controller: "Index",
		Action:     "Ping",
		Data: map[string]int{
			"live_id": 19,
		},
	}

	wsErr := global.WsConn.WriteJSON(PingStruct)
	if wsErr != nil {
		log.Println(wsErr)
	}

}

func (wsC *Clients) WsPushMsg(msg forms.Msg) {
	comment := struct {
		AccessUserToken string  `json:"accessUserToken"`
		Content         string  `json:"content"`
		LiveId          float64 `json:"live_id"`
		UserId          float64 `json:"user_id"`
	}{}
	// 解析评论结构
	b, err := json.Marshal(&msg.Data)
	if err != nil {
		zap.S().Errorf("cannot json Marshal:%v", err)
	}

	err = json.Unmarshal(b, &comment)
	if err != nil {
		zap.S().Errorf("cannot UnMarshal comment:%v", err)
	}

	PingStruct := struct {
		Status int                    `json:"status"`
		Action string                 `json:"action"`
		Data   map[string]interface{} `json:"data"`
	}{
		Action: "WsPushMsg",
		Status: 200,
		Data: map[string]interface{}{
			"type":    "comment",
			"lid":     comment.LiveId,
			"uid":     comment.UserId,
			"comment": comment.Content,
		},
	}

	// 消息广播
	for client, v := range global.WsClients {
		if v {
			if client.Id != wsC.Fc.Id {
				wsErr := client.Ws.WriteJSON(PingStruct)
				if wsErr != nil {
					log.Println(wsErr)
					break
				}
			}
		}
	}

}
