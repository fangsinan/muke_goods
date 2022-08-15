package push

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Push struct {
	Conn *websocket.Conn
	Live int
	Mux  sync.Locker
}
type PingT struct {
	Conn *websocket.Conn
	Ping bool
}

func (p *Push) Ping(c *websocket.Conn) {
	PingStruct := struct {
		Action string `json:"action"`
	}{
		Action: "Ping",
	}
	p.Mux.Lock()
	defer p.Mux.Unlock()
	wsErr := c.WriteJSON(PingStruct)
	if wsErr != nil {
		if !websocket.IsCloseError(wsErr,
			websocket.CloseGoingAway,
			websocket.CloseNormalClosure) {
			log.Printf("[WsPing] unexpected read error %v", wsErr)
		}
	}
	log.Printf("[WsPing] run .....")
}
