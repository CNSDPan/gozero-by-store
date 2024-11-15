package storebecomelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/store/model/sqls"
	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"
	"store/pkg/xcode"
	"strconv"
)

type SaveChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveChatMessageLogic {
	return &SaveChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveChatMessage
// @Desc：记录聊天消息
// @param：in
// @return：res
// @return：err
func (l *SaveChatMessageLogic) SaveChatMessage(in *store.SaveChatReq) (res *store.Response, err error) {
	res = &store.Response{}
	res.Code, res.Message = xcode.GetCodeMessage(xcode.RESPONSE_SUCCESS)
	item := []sqls.ChatLog{}
	for _, chat := range in.List {
		timestamp, e := strconv.ParseInt(chat.SaveTime, 10, 64)
		if e != nil {
			continue
		}
		item = append(item, sqls.ChatLog{
			UserId:    chat.UserId,
			StoreId:   chat.StoreId,
			Message:   chat.Message,
			Timestamp: timestamp,
		})
	}

	e := l.svcCtx.ChatLogModel.InsertChatLogs(item)
	if e != nil {
		l.Logger.Error("批量写入聊天记录 fail:%s", e.Error())
	}

	return res, nil
}
