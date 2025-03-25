package server

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/app/rpc/im/client/socket"
	"store/pkg/types"
	"store/pkg/xcode"
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
// @param：userName
// @param：joinStoreIds 加入的店铺（我的店铺和加入的会员店铺）
// @return：*Client
func NewClient(wsConn *websocket.Conn, clientId int64, userId int64, userName string, joinStoreIds []int64) *Client {
	return &Client{
		ConnectTime:  uint64(time.Now().Unix()),
		WsConn:       wsConn,
		ClientId:     clientId,
		UserId:       userId,
		UserName:     userName,
		StoreIds:     joinStoreIds,
		IsRepeatConn: "",
		Extend:       "",
		HandleClose:  make(chan int, 10),
		Broadcast:    make(chan types.SocketMsg, 10000),
	}
}

// JoinMsg 入会消息推送
func (client *Client) JoinMsg() {

}

// SendPrivateMsg
// @Desc：私信
// @param：msg
func (client *Client) SendPrivateMsg(msg types.SocketMsg) {

}

// SendPublicMsg
// @Desc：公开消息推送
// @param：socketMsg
// @return：code
// @return：msg
// @return：errMsg
func (client *Client) SendPublicMsg(socketMsg types.SocketMsg) (code string, msg string, errMsg string, err error) {
	code, msg = xcode.GetCodeMessage(xcode.SOCKET_BROADCAST_MSG_FAIL)
	b, err := jsonx.Marshal(socketMsg.Body)
	if err != nil {
		errMsg = err.Error()
		return
	}
	res, e := WsServer.RpcMap.Socket.BroadcastMsg(context.Background(), &socket.BroadcastReq{
		Operate:       int32(socketMsg.Operate),
		Method:        socketMsg.Method,
		StoreId:       socketMsg.StoreId,
		SendUserId:    client.UserId,
		SendUserName:  client.UserName,
		ReceiveUserId: socketMsg.ReceiveUserId,
		Extend:        socketMsg.Extend,
		Body:          string(b),
	})
	if e != nil {
		errMsg = e.Error()
		return
	} else if res.Code != xcode.RESPONSE_SUCCESS {
		return res.Code, res.Message, res.ErrMsg, nil
	}
	return xcode.RESPONSE_SUCCESS, "", "", nil
}

// PushMsg
// @Desc：写入发送管道
// @param：msg
func (client *Client) PushMsg(msg types.SocketMsg) {
	select {
	case client.Broadcast <- msg:
	default:
	}
	return
}
