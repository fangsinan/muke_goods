package global

import (
	"net/http"
	"webApi/live_ws/config"
	"webApi/live_ws/forms"

	ut "github.com/go-playground/universal-translator"
	"github.com/gorilla/websocket"
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

var (
	Ping      = make(chan struct{}, 1024)
	PushMsg   = make(chan struct{}, 1024)
	WsClients = make(map[*forms.Client]bool)
)
