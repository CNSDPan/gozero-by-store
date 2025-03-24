package storebecomelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/rpc/store/internal/svc"
	"store/app/rpc/store/pb/store"
	mysqlModel "store/db/dao/model"
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

func (l *SaveChatMessageLogic) SaveChatMessage(in *store.SaveChatReq) (res *store.Response, err error) {
	res = &store.Response{}
	res.Code, res.Message = xcode.GetCodeMessage(xcode.RESPONSE_SUCCESS)
	item := make([]mysqlModel.ChatLog, 0)
	for _, chat := range in.List {
		timestamp, e := strconv.ParseInt(chat.SaveTime, 10, 64)
		if e != nil {
			continue
		}
		item = append(item, mysqlModel.ChatLog{
			UserId:    chat.UserId,
			StoreId:   chat.StoreId,
			Message:   chat.Message,
			Timestamp: timestamp,
		})
	}

	e := l.svcCtx.MysqlQuery.ChatLog.InsertChatLogs(item)
	if e != nil {
		l.Logger.Error("批量写入聊天记录 fail:%s", e.Error())
	}

	return res, nil
}
