package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"sync"
	"time"
	"webApi/live_ws/api/handler"
	"webApi/live_ws/forms"
	"webApi/live_ws/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Msg = forms.Msg{}

// WsPushInter 定义ws接口
type WsPushInter interface {
	WsPing()
	WsPushMsg(forms.Msg)
}

// WsService 定义ws接口
type WsService struct {
	WsPush WsPushInter
	RC     *forms.Client
	//Mux    sync.RWMutex
}

// WsHandler websocket应用
func WsHandler(c *gin.Context) {

	// 处理参数
	liveIdStr := c.DefaultQuery("live_id", "0")
	liveId, _ := strconv.Atoi(liveIdStr)
	if liveId <= 0 {
		zap.S().Errorf("直播间不存在")
		return
	}
	//初始化liveId
	if _, ok := global.WsClients[liveId]; !ok {
		global.WsClients[liveId] = make(global.ClientT)
	}

	// 升级net 为websocket
	var (
		conn *websocket.Conn
		err  error
	)
	conn, err = global.WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.S().Errorf("ws upgrader error: %v", err)
		return
	}

	// 用户链接基础信息
	userC := &forms.Client{
		Id:      fmt.Sprintf("%s_%v", liveIdStr, uuid.NewV4()),
		Ws:      conn,
		IsValid: true,
		Mux:     &sync.Mutex{},
	}

	// 配置推送服务的接口
	Ws := WsService{
		WsPush: &handler.Clients{
			Fc:     userC,
			LiveId: liveId,
		},
		RC: userC,
	}
	defer func() {
		userC.IsValid = false
		userC.Ws.Close()
		zap.S().Infof("退出 =====关闭连接")
	}()

	// 写入连接池进行广播
	global.WsClients[liveId][conn] = userC

	zap.S().Infof("[用户池] ：%v", global.WsClients)
	for k, v := range global.WsClients {
		zap.S().Infof("直播间%d有 %d链接", k, len(v))
		for _, client := range v {
			zap.S().Infof("直播间id:  %d ", client.Id)
		}
	}
	// 监控发送「
	go Ws.Send()

	// 监控接收消息
	Ws.Receive()
}

// Receive 循环接收消息
func (s *WsService) Receive() {
	for {
		//校验当前链接是否有效
		if s.RC.IsValid == true {

			s.RC.Mux.Lock()
			_, b, err := s.RC.Ws.ReadMessage()
			s.RC.Mux.Unlock()
			if err != nil {
				//zap.S().Infof("info read error %v", err)
				if !websocket.IsCloseError(err,
					websocket.CloseGoingAway,
					websocket.CloseNormalClosure,
					websocket.CloseNoStatusReceived,
				) {
					zap.S().Infof("unexpected read error %v", err)
				}
				break
			}
			err = json.Unmarshal(b, &Msg)
			if err != nil {
				zap.S().Infof("unexpected read unmarshal error %v", err)
				return
			}

			// 转发调用不同动作
			switch Msg.Action {
			case "Ping":
				global.Ping <- struct{}{}
			case "Comment":
				global.PushMsg <- struct{}{}
			}
		}

	}
}

func (s *WsService) Send() {
	// 监控client发送的 json 消息
	// Ping消息体 { action: "Ping", data: {id: 123}}
	// Ping消息体 controller: "Push", action: "Comment", data: {content: "你好",}
	for {
		select {
		case <-global.Ping: //每5秒发一次ping
			s.WsPush.WsPing()
		case <-global.PushMsg: //每收到一次评论 广播一次
			s.WsPush.WsPushMsg(Msg)
		case <-time.After(3600 * time.Second): // 设置10秒超时机制 未收到任何消息 关闭当前链接
			zap.S().Infof("==10===关闭连接")
			s.RC.IsValid = false
			s.RC.Ws.Close()
		default:
		}
	}
}
