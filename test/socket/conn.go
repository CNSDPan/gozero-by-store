package socket

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io"
	"log"
	"net/http"
	"store/app/api/im/server"
	"store/pkg/consts"
	"store/pkg/types"
	"sync"
	"time"
)

var ws = "ws://0.0.0.0:6999/v1/ws/socket"

type Server struct {
	Token  string
	Client *server.Client
}

func NewServer(token string) *Server {
	return &Server{
		Token: token,
	}
}

func (s *Server) RunSocket(clientId int64, userId int64) {
	if s.Token == "" {
		panic("token为空，不进行InitSocket\n")
	}
	wsHead := http.Header{}

	wsHead.Set("X-API-Version", "v1")
	wsHead.Set("X-Source", "ws")
	wsHead.Set("X-Request-Time", time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"))
	wsHead.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))
	ws = fmt.Sprintf("%s?Authorization=%s", ws, s.Token)
	conn, res, err := websocket.DefaultDialer.Dial(ws, wsHead)

	if err != nil {
		log.Printf("拨号失败:%v fail:%s", res, err.Error())
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	ctx, cancel := context.WithCancel(context.Background())

	s.Client = server.NewClient(conn, clientId, userId, "测试用户", []int64{})
	defer s.CloseClient()
	defer cancel()

	// 当client关闭时，要通知所以的子协程关闭且清理
	go func(wg *sync.WaitGroup, client *server.Client, cancel context.CancelFunc) {
		defer wg.Done()
		defer cancel()
		for i := range client.HandleClose {
			log.Printf("%d :HandleClose signal received", i)
			return
		}
	}(wg, s.Client, cancel)

	go s.Writ(ctx, wg)
	go s.Read(ctx, wg)

	log.Println("握手成功")

	go func() {
		select {
		case <-time.After(time.Second * 5):
			s.Client.Broadcast <- types.SocketMsg{
				Operate:       consts.OperatePublic,
				Method:        consts.MethodNormal,
				StoreId:       1837056807609659392,
				SendUserId:    1837053720228618240,
				ReceiveUserId: 0,
				Extend:        "",
				Body: types.SocketMsgBody{
					Operate:      consts.OperatePublic,
					Method:       consts.MethodNormal,
					ResponseTime: time.Now().UTC().Format("2006-01-02 15:04:05"),
					Event: types.Event{
						Params: fmt.Sprintf("我来了"),
					},
				},
			}
			log.Println("发送消息了")
		}
	}()

	wg.Wait()
}

func (s *Server) CloseClient() {
	_ = s.Client.WsConn.Close()
	close(s.Client.HandleClose)
	log.Println("client 连接关闭")
}

func (s *Server) Writ(ctx context.Context, wg *sync.WaitGroup) {
	var (
		w      io.WriteCloser
		err    error
		ticker = time.NewTicker(time.Duration(30) * time.Second)
		b      []byte
	)
	defer func() {
		s.Client.HandleClose <- 1
		log.Printf("Writ defer")
		wg.Done()
	}()
	for {
		select {
		case message, ok := <-s.Client.Broadcast:
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = s.Client.WsConn.SetWriteDeadline(time.Now().Add(time.Duration(5) * time.Second))
			if !ok {
				log.Println(" 写消息协程无法正常获取消息")
				_ = s.Client.WsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err = s.Client.WsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf(" 写消息 fail:%s", err.Error())
				return
			}
			b, err = jsonx.Marshal(message)
			if err != nil {
				log.Printf(" 写消息 jsonx.Marshal() fail:%s", err.Error())
				continue
			}
			_, _ = w.Write(b)
			if err = w.Close(); err != nil {
				log.Printf(" 写消息  w.Close() fail:%s", err.Error())
				return
			}
		case <-ticker.C:
			// 心跳检测
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = s.Client.WsConn.SetWriteDeadline(time.Now().Add(time.Duration(30) * time.Second))
			if err = s.Client.WsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf(" writ ping pong fail:%v", err.Error())
				return
			}
		case <-ctx.Done():
			log.Printf("Writ 接到了关闭通知")
			return
		}
	}
}

func (s *Server) Read(ctx context.Context, wg *sync.WaitGroup) {
	var (
		ticker = time.NewTicker(2 * time.Millisecond)
	)
	defer func() {
		s.Client.HandleClose <- 1
		log.Printf("Read defer")
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Read 接到了关闭通知")
			return
		case <-ticker.C:
			messageType, message, err := s.Client.WsConn.ReadMessage()
			if err != nil || (message == nil && messageType == -1) {
				log.Printf("断开连接 messageType:%d; fail:%v", messageType, err)
				return
			}
			// 每次需设置读超时时间，否则接收不到
			s.Client.WsConn.SetReadLimit(8192)
			_ = s.Client.WsConn.SetReadDeadline(time.Now().Add(time.Duration(60) * time.Second))
			s.Client.WsConn.SetPongHandler(func(string) error {
				_ = s.Client.WsConn.SetReadDeadline(time.Now().Add(time.Duration(54) * time.Second))
				log.Printf("接收不到消息")
				return nil
			})
			log.Printf("输出通讯消息:%v", string(message))
		}
	}
}
