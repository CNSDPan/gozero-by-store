package apistorelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/rpc/api/internal/svc"
	"store/app/rpc/api/pb/api"
	"store/pkg/xcode"
)

type MyAllStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMyAllStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyAllStoreLogic {
	return &MyAllStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MyAllStore
// @Desc：获取我的店铺和会员店铺的所有店铺ID
// @param：in
// @return：res
// @return：err
func (l *MyAllStoreLogic) MyAllStore(in *api.MyAllStoreIdReq) (res *api.MyAllStoreIdRes, err error) {
	var (
		code = xcode.RESPONSE_SUCCESS
	)
	res = &api.MyAllStoreIdRes{
		Result:  &api.Response{},
		StoreId: make([]int64, 0),
	}
	res.StoreId = l.svcCtx.MysqlQuery.User.MyStoreIds(in.UserId)
	res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
	return res, nil
}
