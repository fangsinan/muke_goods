package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"webApi/live_ws/forms"
	"webApi/live_ws/global"

	"go.uber.org/zap"
)

type Clients struct {
	Fc     *forms.Client
	LiveId int
	//Mux    sync.RWMutex
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
			"live_id": wsC.LiveId,
		},
	}
	wsC.Fc.Mux.Lock()
	// 返回当前链接
	wsErr := wsC.Fc.Ws.WriteJSON(PingStruct)
	wsC.Fc.Mux.Unlock()
	if wsErr != nil {
		if !websocket.IsCloseError(wsErr,
			websocket.CloseGoingAway,
			websocket.CloseNormalClosure) {
			zap.S().Errorf("[WsPing] unexpected read error %v", wsErr)
		}
	}

}

func (wsC *Clients) WsPushMsg(msg forms.Msg) {

	//校验当前goroutine的链接 是否有效  如果无效  无需广播
	if wsC.Fc.IsValid == false {
		return
	}
	//comment := struct {
	//	AccessUserToken string  `json:"accessUserToken"`
	//	Content         string  `json:"content"`
	//	LiveId          float64 `json:"live_id"`
	//	UserId          float64 `json:"user_id"`
	//}{}
	//// 解析评论结构
	b, err := json.Marshal(&msg.Data)
	if err != nil {
		zap.S().Errorf("cannot json Marshal:%v", err)
	}

	//err = json.Unmarshal(b, &comment)
	//if err != nil {
	//	zap.S().Errorf("cannot UnMarshal comment:%v", err)
	//}
	//
	//PingStruct := struct {
	//	Status int                    `json:"status"`
	//	Action string                 `json:"action"`
	//	Data   map[string]interface{} `json:"data"`
	//}{
	//	Action: "WsPushMsg",
	//	Status: 200,
	//	Data: map[string]interface{}{
	//		"type":    "comment",
	//		"lid":     comment.LiveId,
	//		"uid":     comment.UserId,
	//		"comment": comment.Content,
	//	},
	//}

	PingStruct := struct {
		Status int    `json:"status"`
		Action string `json:"action"`
		Data   string `json:"data"`
	}{
		Action: "WsPushMsg",
		Status: 200,
		Data:   "[我是服务器]: " + string(b),
	}
	// 开辟新内存 接收uid
	uid := wsC.Fc.Id

	// 消息广播
	i := 0
	ids := ""
	for _, v := range global.WsClients[wsC.LiveId] {
		if v.IsValid == true {
			ids += "------- " + v.Id
			if v.Id == uid { // id 相同
				continue
			}

			if v.Ws == wsC.Fc.Ws { // conn 相同
				continue
			}
			i++

			wsC.Fc.Mux.Lock()
			wsErr := v.Ws.WriteJSON(PingStruct)
			wsC.Fc.Mux.Unlock()
			if wsErr != nil {
				zap.S().Errorf("[WsPushMsg] Write  Marshal:%v", wsErr)
				break
			}
		}
	}

	zap.S().Infof("当前id为%v   \n 执行的live_id:%v  循环了%d次: \n    内容是：%v", uid, wsC.LiveId, i, ids)

}
