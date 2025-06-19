package client

import (
	"chatbackendapp/internal/common"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10 // Must be less than pongWait
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ReadPump(c *common.Client) {
	defer func() {
		c.Hub.Unregister <- c
		if ws, ok := c.Conn.(*websocket.Conn); ok {
			ws.Close()
		}
	}()
}
