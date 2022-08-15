package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/gorilla/websocket"
	"net/http"
	"webApi/live_ws/config"
	"webApi/live_ws/forms"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
	WsConn       *websocket.Conn
	WsUpgrader   = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ClientT map[*websocket.Conn]*forms.Client

var (
	Ping      = make(chan struct{}, 1024)
	PushMsg   = make(chan struct{}, 1024)
	WsClients = make(map[int]ClientT)
)
