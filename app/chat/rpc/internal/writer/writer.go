package writer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/store/rpc/store/storebecome"
	"store/pkg/types"
	"strconv"
	"time"
)

type BroadcastWriter struct {
	Body chan *types.SocketMsgBody
	Log  logx.Logger
}

func NewBroadcastWriter() *BroadcastWriter {
	return &BroadcastWriter{
		Body: make(chan *types.SocketMsgBody, 5000),
		Log:  logx.WithContext(context.Background()),
	}
}

func (w *BroadcastWriter) SaveChan(broadcast *types.SocketMsgBody) {
	go func() {
		w.Body <- broadcast
	}()
}

func (w *BroadcastWriter) WriteChat(writeRpc storebecome.StoreBecome) {
	var (
		err      error
		chat     = &types.SocketMsgBody{}
		chatData []*storebecome.SaveChatItem
	)
Label:
	err = nil
	chatData = []*storebecome.SaveChatItem{}
	for {
		select {
		case chat = <-w.Body:
			dataByNormal, _ := chat.Event.Data.(types.DataByNormal)
			chatData = append(chatData, &storebecome.SaveChatItem{
				UserId:   dataByNormal.SendUserId,
				StoreId:  dataByNormal.StoreId,
				Message:  dataByNormal.Message,
				SaveTime: strconv.FormatInt(chat.Timestamp, 10),
			})
			if len(chatData) > 5000 {
				_, err = writeRpc.SaveChatMessage(context.Background(), &storebecome.SaveChatReq{List: chatData})
				if err != nil {
					w.Log.Errorf("写入记录管道请求rpc fail：%s", err.Error())
				}
				goto Label
			}
		case <-time.After(3 * time.Second):
			if len(chatData) > 0 {
				_, err = writeRpc.SaveChatMessage(context.Background(), &storebecome.SaveChatReq{List: chatData})
				if err != nil {
					w.Log.Errorf("写入记录管道请求rpc fail：%s", err.Error())
				}
				goto Label
			}
		}
	}
}
