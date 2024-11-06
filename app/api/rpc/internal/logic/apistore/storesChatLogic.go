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

type StoresChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoresChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoresChatLogic {
	return &StoresChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StoresChat
// @Desc：获取聊天记录
// @param：in
// @return：res
// @return：err
func (l *StoresChatLogic) StoresChat(in *api.StoreChatReq) (res *api.StoreChatRes, err error) {
	var (
		e    error
		code = xcode.RESPONSE_SUCCESS
	)
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取店铺聊天记录 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	res = &api.StoreChatRes{
		Result: &api.Response{},
		Data:   &apistore.StoresChatMap{},
	}
	items, e := l.svcCtx.StoreModel.ChatLogMgr.SelectPageChatLog(
		sqlsStore.NewPage(10, int64(in.Offset)),
		in.StoreId,
		in.Timestamp,
	)
	if e != nil {
		code = xcode.CHAT_LOG_GET_FAIL
		return res, nil
	} else {
		// 获取店铺会员列表
		userItems := l.svcCtx.StoreModel.StoresMemberMgr.MapKeyUserId(
			in.StoreId,
		)
		rows := make([]*api.StoreChatItem, len(items.GetRecords().([]sqlsStore.ChatLogApi)))
		for k, _ := range items.GetRecords().([]sqlsStore.ChatLogApi) {
			chat := items.GetRecords().([]sqlsStore.ChatLogApi)[k]
			userId := strconv.FormatInt(chat.UserId, 10)
			memberUser, ok := userItems[userId]

			parsedTimeStr := ""
			timestampStr := "0"
			if chat.CreatedAt != "" {
				parsedTime, _ := time.Parse(time.RFC3339, chat.CreatedAt)
				parsedTimeStr = parsedTime.Format("2006-01-02 15:04:05")
			}
			if chat.Timestamp != 0 {
				timestampStr = strconv.FormatInt(chat.Timestamp, 10)
			}
			if ok {
				rows[k] = &api.StoreChatItem{
					UserId:    userId,
					UserName:  &memberUser.User.Name,
					StoreId:   strconv.FormatInt(chat.StoreId, 10),
					StoreName: chat.StoreName,
					Message:   &chat.Message,
					Timestamp: &timestampStr,
					CreateAt:  &parsedTimeStr,
				}
			} else {
				userName := ""
				rows[k] = &api.StoreChatItem{
					UserId:    userId,
					UserName:  &userName,
					StoreId:   strconv.FormatInt(chat.StoreId, 10),
					StoreName: chat.StoreName,
					Message:   &chat.Message,
					Timestamp: &timestampStr,
					CreateAt:  &parsedTimeStr,
				}
			}
		}
		res.Data.Limit = items.GetSize()
		res.Data.Offset = items.GetCurrent()
		res.Data.Page = items.GetPages()
		res.Data.Current = items.GetCurrent()
		res.Data.Total = items.GetTotal()
		res.Data.Rows = rows
	}
	return res, nil
}
