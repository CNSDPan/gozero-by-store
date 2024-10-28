package server

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/pkg/biz"
	"store/pkg/types"
	"time"
)

var PubSub *redis.PubSub

// NewRedisMq
// @Desc：初始化redis的发布订阅服务
// @param：bizCon
// @return：error
func NewRedisMq(bizCon *redis.Client) error {
	ctx := context.Background()
	PubSub = bizCon.Subscribe(ctx, biz.SOCKET_PUB_SUB_BROADCAST_NORMAL_CHAN_KEY)
	if _, err := PubSub.ReceiveTimeout(ctx, 1*time.Second); err != nil {
		WsServer.Log.Errorf("socket 订阅 %s 接收消息异常，尝试 ping...", biz.SOCKET_PUB_SUB_BROADCAST_NORMAL_CHAN_KEY)
		if err = PubSub.Ping(ctx, ""); err != nil {
			return err
		}
	}
	go func() {
		SubReceive()
	}()
	return nil
}

func SubReceive() {
	var err error
	defer PubSub.Close()
	pubSubCh := PubSub.Channel()
	for msg := range pubSubCh {
		writeMsg := types.SocketMsgBody{
			Operate:      0,
			Method:       "",
			ResponseTime: "",
			Event: types.Event{
				Params: "",
				Data:   types.DataByNormal{},
			},
		}

		b := []byte(msg.Payload)
		if err = jsonx.Unmarshal(b, &writeMsg); err != nil {
			WsServer.Log.Errorf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", msg.Channel, err.Error())
		} else {
			// 群消息
			for _, bucket := range WsServer.Buckets {
				bucket.BroadcastMsg(writeMsg)
			}
		}
	}
	return
}
