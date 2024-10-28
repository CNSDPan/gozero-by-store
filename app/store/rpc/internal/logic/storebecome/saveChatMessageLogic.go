package storebecomelogic

import (
	"context"
	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"
	"store/pkg/xcode"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
	if in.StoreId == 0 || in.UserId == 0 || in.Message == "" {
		return res, nil
	}
	t, e := time.Parse("2006-01-02 15:04:05", in.SaveTime)
	if e != nil {
		l.Logger.Errorf("%s time format err:%s", l.svcCtx.Config.ServiceName, e.Error())
		return res, nil
	}
	l.svcCtx.ChatLogModel.CreateChatLog(in.StoreId, in.UserId, in.Message, t)

	return res, nil
}
