package server

import (
	"github.com/gorilla/websocket"
	"store/pkg/types"
	"time"
)

type Client struct {
	ConnectTime  uint64
	WsConn       *websocket.Conn
	ClientId     int64
	UserId       int64
	UserName     string
	StoreIds     []int64
	IsRepeatConn string
	Extend       string
	HandleClose  chan int
	Broadcast    chan types.SocketMsg
}

// NewClient
// @Desc：初始化client
// @param：wsConn
// @param：clientId
// @param：userId
// @return：*Client
func NewClient(wsConn *websocket.Conn, clientId int64, userId int64, userName string) *Client {
	return &Client{
		ConnectTime:  uint64(time.Now().Unix()),
		WsConn:       wsConn,
		ClientId:     clientId,
		UserId:       userId,
		UserName:     userName,
		IsRepeatConn: "",
		Extend:       "",
		HandleClose:  make(chan int, 10),
		Broadcast:    make(chan types.SocketMsg, 10000),
	}
}

// JoinMsg 入会消息推送
func (client *Client) JoinMsg() {

}

// SendMsg 普通消息推送
func (client *Client) SendMsg(msg types.SocketMsg) {

}
