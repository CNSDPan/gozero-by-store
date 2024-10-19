package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Connect struct{}

// NewConnect
// @Desc：初始结构体
// @return：*Connect
func NewConnect() *Connect {
	return &Connect{}
}

// Run
// @Desc：建立WS连接
// @param：w
// @param：r
// @return：*websocket.Conn
// @return：error
func (c *Connect) Run(w http.ResponseWriter, r *http.Request, maxMessageSize int64, PongPeriod time.Duration) (*websocket.Conn, error) {
	wsConn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	wsConn.SetReadLimit(maxMessageSize)
	wsConn.SetPongHandler(func(string) error {
		_ = wsConn.SetReadDeadline(time.Now().Add(PongPeriod))
		return nil
	})

	return wsConn, nil
}
