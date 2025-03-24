package server

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"store/app/chat/rpc/chat/socket"
	"store/pkg/consts"
	"store/pkg/types"
	"store/pkg/util"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	Node       *snowflake.Node
	Log        logx.Logger
	ServerName string
	ServerId   string
	ServerIp   string
	Buckets    []*Bucket
	LenBucket  uint32
	Option     Options
	RpcMap     RpcCl
}

type Options struct {
	BucketNum uint
	// MaxMessageSize 消息最大字节
	MaxMessageSize int64
	// PingPeriod 每次ping的间隔时长
	PingPeriod time.Duration
	// PongPeriod 每次pong的间隔时长，可以是PingPeriod的一倍|两倍
	PongPeriod time.Duration
	// WriteWait client的写入等待超时时长
	WriteWait time.Duration
	// ReadWait client的读取等待超时时长
	ReadWait time.Duration
}

type RpcCl struct {
	Socket socket.Socket
}

var WsServer *Server

// NewServer
// @Desc：初始化socket的服务
// @param：serverId
// @param：serverName
// @param：serverIp
// @param：serverPort
// @param：node
// @param：l
// @return：*Server
func NewServer(serverId string, serverName string, serverIp string, optionsConf types.SocketOptionsConf, node *snowflake.Node, l logx.Logger) *Server {
	buckets := NewBucket(optionsConf.BucketNum)
	options := Options{
		BucketNum:      optionsConf.BucketNum,
		MaxMessageSize: optionsConf.MaxMessageSize,
		PingPeriod:     time.Duration(optionsConf.PingPeriod) * time.Second,
		PongPeriod:     time.Duration(optionsConf.PongPeriod) * time.Second,
		WriteWait:      time.Duration(optionsConf.WriteWait) * time.Second,
		ReadWait:       time.Duration(optionsConf.ReadWait) * time.Second,
	}
	WsServer = &Server{
		Node:       node,
		Log:        l,
		ServerName: serverName,
		ServerId:   serverId,
		ServerIp:   serverIp,
		Buckets:    buckets,
		LenBucket:  uint32(len(buckets)),
		Option:     options,
	}
	return WsServer
}

// SetSocketRpc
// @Desc：将外层初始化好的RPC客户端传进来
// @param：socket
func (s *Server) SetSocketRpc(socket socket.Socket) {
	s.RpcMap.Socket = socket
}

// Run
// @Desc：协程允许每个client的I/O
// @param：client
// @param：storeIds
func (s *Server) Run(client *Client, storeIds []int64) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	ctx, cancel := context.WithCancel(context.Background())
	defer s.CloseClient(client)
	defer cancel()

	// 当client关闭时，要通知所以的子协程关闭且清理
	go func(wg *sync.WaitGroup, client *Client, cancel context.CancelFunc) {
		defer wg.Done()
		defer cancel()
		for range client.HandleClose {
			return
		}
	}(wg, client, cancel)

	// 协程I/O
	go s.ReadChannel(ctx, wg, client)
	go s.WriteChan(ctx, wg, client)

	// 消息加载我的店铺和入会店铺
	s.GetBucket(client.UserId).AddBucket(client, storeIds...)
	wg.Wait()
}

// GetBucket
// @Desc：获取连接桶
// @param：client
func (s *Server) GetBucket(userId int64) *Bucket {
	userIdStr := strconv.FormatInt(userId, 10)
	// 通过cityHash算法 % 池子数量进行取模,得出需要放入哪个连接池里
	idx := util.CityHash32([]byte(userIdStr), uint32(len(userIdStr))) % s.LenBucket
	return s.Buckets[idx]
}

// CloseClient
// @Desc：关闭ws连接
// @param：client
func (s *Server) CloseClient(client *Client) {
	_ = client.WsConn.Close()
	close(client.HandleClose)
	s.GetBucket(client.UserId).UnBucket(client)
	s.Log.Infof("%s client 连接关闭;user:[%d,%s]", s.ServerName, client.UserId, client.UserName)
}

func (s *Server) WriteChan(ctx context.Context, wg *sync.WaitGroup, client *Client) {
	var (
		w           io.WriteCloser
		err         error
		ticker      = time.NewTicker(s.Option.PingPeriod)
		b           []byte
		clientClose = false
	)
	defer func() {
		if r := recover(); r != nil {
			s.Log.Errorf("%s WriteChan 严重异常捕抓 fail:%v", s.ServerName, r)
		}
	}()
	defer func() {
		if !clientClose {
			client.HandleClose <- 1
		}
		wg.Done()
	}()
	for {
		select {
		case message, ok := <-client.Broadcast:
			s.Log.Errorf("消息消息:%+v", message)
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = client.WsConn.SetWriteDeadline(time.Now().Add(s.Option.WriteWait))
			if !ok {
				s.Log.Errorf("%s 写消息协程无法正常获取消息", s.ServerName)
				_ = client.WsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err = client.WsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				s.Log.Errorf("%s 写消息 fail:%s", s.ServerName, err.Error())
				return
			}

			b, err = jsonx.Marshal(message.Body)
			if err != nil {
				s.Log.Errorf("%s 写消息 jsonx.Marshal() :%s", s.ServerName, err.Error())
				continue
			}
			_, _ = w.Write(b)
			if err = w.Close(); err != nil {
				s.Log.Errorf("%s 写消息 w.Close() :%s", s.ServerName, err.Error())
				return
			}
		case <-ticker.C:
			// 心跳检测
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = client.WsConn.SetWriteDeadline(time.Now().Add(s.Option.PingPeriod))
			if err = client.WsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				s.Log.Errorf("%s ping pong fail:%s", s.ServerName, err.Error())
				return
			}
		case <-ctx.Done():
			clientClose = true
			return
		}
	}
}

func (s *Server) ReadChannel(ctx context.Context, wg *sync.WaitGroup, client *Client) {
	var (
		e           error
		ticker      = time.NewTicker(2 * time.Millisecond)
		clientClose = false
		broadcast   = types.SocketMsg{
			Body: types.SocketMsgBody{
				Operate:      0,
				Method:       "",
				ResponseTime: "",
				Event:        types.Event{},
			},
		}
	)
	defer func() {
		if r := recover(); r != nil {
			s.Log.Errorf("%s ReadChannel 严重异常捕抓 fail:%v", s.ServerName, r)
		}
	}()
	defer func() {
		if !clientClose {
			client.HandleClose <- 1
		}
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			clientClose = true
			return
		case <-ticker.C:
			messageType, message, err := client.WsConn.ReadMessage()
			if err != nil || (message == nil && messageType == -1) {
				s.Log.Errorf("%s read client 连接关闭;user:[%d,%s]; messageType:%d; fail:%v", s.ServerName, client.UserId, client.UserName, messageType, err)
				return
			}
			// 每次需设置读超时时间，否则接收不到
			client.WsConn.SetReadLimit(s.Option.MaxMessageSize)
			_ = client.WsConn.SetReadDeadline(time.Now().Add(s.Option.ReadWait))
			client.WsConn.SetPongHandler(func(string) error {
				_ = client.WsConn.SetReadDeadline(time.Now().Add(s.Option.PongPeriod))
				return nil
			})
			err = jsonx.Unmarshal(message, &broadcast)
			if err != nil {
				s.Log.Errorf("%s 消息转换指定结构 fail:%s", s.ServerName, err.Error())
				continue
			}
			t := time.Now()
			broadcast.Body.ResponseTime = t.UTC().Format("2006-01-02 15:04:05")
			broadcast.Body.Timestamp = t.UnixMicro()

			switch broadcast.Operate {
			case consts.OperatePrivate:

			case consts.OperatePublic:
				if broadcast.Method == consts.MethodNormal {
					_, _, _, e = client.SendPublicMsg(broadcast)
				}

				if e != nil {
					s.Log.Errorf("%s OperatePublic 广播 消息 fail:%v", e)
				}
			}
		}
	}
}
