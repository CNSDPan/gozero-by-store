package store

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
	"store/app/rpc/api/client/apistore"
	"store/app/rpc/api/client/apiuser"
	"store/pkg/xcode"
)

type StoreInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreInfoLogic {
	return &StoreInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreInfoLogic) StoreInfo(req *types.StoreInfoReq) (res *types.Response, resp *types.StoreInfoRes, err error) {
	code := ""
	res = &types.Response{}
	resp = &types.StoreInfoRes{
		StoreUser: types.StoreUser{},
	}
	rpcRes := &apistore.StoreInfoRes{}
	userRpcRes := &apiuser.UserInfoRes{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func() {
		if err != nil {
			if code == "" {
				l.Logger.Errorf("%s 系统错误 fail:%v", l.svcCtx.Config.ServiceName, err.Error())
				res.Code, res.Message = xcode.GetCodeMessage(xcode.RESPONSE_FAIL)
			} else {
				res.ErrMsg = err.Error()
				res.Code, res.Message = xcode.GetCodeMessage(code)
			}
		} else {
			res.ErrMsg = rpcRes.Result.ErrMsg
			res.Code = rpcRes.Result.Code
			res.Message = rpcRes.Result.Message

			if res.Code == xcode.RESPONSE_SUCCESS {
				resp.StoreId = rpcRes.StoreId
				resp.Name = rpcRes.Name
				resp.Avatar = rpcRes.Avatar
				resp.Contacts = *rpcRes.Contacts
				resp.StoreUser = types.StoreUser{
					StoreUserId: rpcRes.StoreUserId,
					UserId:      rpcRes.UserId,
					Mobile:      userRpcRes.Mobile,
					Name:        userRpcRes.Name,
					Avatar:      userRpcRes.Avatar,
				}
			}

		}
	}()
	if rpcRes, err = l.svcCtx.ApiRpcCl.Store.Info(ctx, &apistore.StoreInfoReq{StoreId: req.StoreId, UserId: req.UserId}); err != nil {
		return
	}
	userRpcRes, err = l.svcCtx.ApiRpcCl.User.Info(ctx, &apiuser.UserInfoReq{
		UserId: rpcRes.UserId,
	})
	return
}
