package apistorelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/rpc/api/internal/svc"
	"store/app/rpc/api/pb/api"
	mysqlModel "store/db/dao/model"
	mysqlQuery "store/db/dao/query"
	"store/pkg/xcode"
	"strconv"
)

type MemberUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUserListLogic {
	return &MemberUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberUserList
// @Desc：会员列表
// @param：in
// @return：res
// @return：err
func (l *MemberUserListLogic) MemberUserList(in *api.MemberUsersItemReq) (res *api.MemberUsersItemRes, err error) {
	var (
		e    error
		code = xcode.RESPONSE_SUCCESS
	)
	res = &api.MemberUsersItemRes{
		Result: &api.Response{},
		Data:   &api.UsersMap{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取店铺列表 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	items, e := l.svcCtx.MysqlQuery.StoreMember.SelectPageApi(
		mysqlQuery.NewPage(int64(in.Limit), int64(in.Offset)),
		l.svcCtx.MysqlQuery.StoreMember.WithStoreId(in.StoreId),
	)
	if e != nil {
		code = xcode.STORE_MEMBER_ITEM_FAIL
		return
	} else {
		rows := make([]*api.UserItem, len(items.GetRecords().([]mysqlModel.MemberUserItem)))
		for k, item := range items.GetRecords().([]mysqlModel.MemberUserItem) {
			rows[k] = &api.UserItem{
				UserId: strconv.FormatInt(item.UserId, 10),
				Name:   item.User.Name,
				Avatar: item.User.Avatar,
				Mobile: item.User.Mobile,
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
