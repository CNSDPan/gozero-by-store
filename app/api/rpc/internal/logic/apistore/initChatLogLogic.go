package apistorelogic

import (
	"context"
	"store/app/api/rpc/api/apistore"
	sqlsStore "store/app/store/model/sqls"
	"store/pkg/xcode"
	"strconv"
	"time"

	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitChatLogLogic {
	return &InitChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InitChatLog
// @Desc：每次获取10个店铺的最新一条聊天记录
// @param：in
// @return：res
// @return：err
func (l *InitChatLogLogic) InitChatLog(in *api.InitChatLogReq) (res *api.InitChatLogRes, err error) {
	var (
		e    error
		code = xcode.RESPONSE_SUCCESS
	)
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 初始化店铺聊天记录 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	res = &api.InitChatLogRes{
		Result: &apistore.Response{},
		Data:   &apistore.StoresChatMap{},
	}
	items, e := l.svcCtx.StoreModel.ChatLogMgr.InitChatLog(
		sqlsStore.NewPage(10, 0),
		in.UserId,
	)
	if e != nil {
		code = xcode.CHAT_LOG_INIT_FAIL
		return res, nil
	} else {
		rows := make([]*api.StoreChatItem, len(items.GetRecords().([]sqlsStore.ChatLogApi)))
		for k, item := range items.GetRecords().([]sqlsStore.ChatLogApi) {
			parsedTime, _ := time.Parse(time.RFC3339, item.CreatedAt)
			rows[k] = &api.StoreChatItem{
				UserId:    strconv.FormatInt(item.UserId, 10),
				StoreId:   strconv.FormatInt(item.StoreId, 10),
				StoreName: item.StoreName,
				Message:   item.Message,
				Timestamp: strconv.FormatInt(item.Timestamp, 10),
				CreateAt:  parsedTime.Format("2006-01-02 15:04:05"),
			}
		}
		res.Data.Limit = items.GetSize()
		res.Data.Offset = items.GetCurrent()
		res.Data.Page = items.GetPages()
		res.Data.Current = items.GetCurrent()
		res.Data.Total = items.GetTotal()
		res.Data.Rows = rows
	}
	return
}
