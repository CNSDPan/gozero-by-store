package store

import (
	"context"
	"store/app/api/rpc/api/apiuser"
	"store/app/store/rpc/pb/store"
	"store/pkg/xcode"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreBecomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreBecomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreBecomeLogic {
	return &StoreBecomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreBecomeLogic) StoreBecome(req *types.StoreBecomeReq, token string) (res *types.Response, resp *types.StoreBecomeRes, err error) {
	code := ""
	res = &types.Response{}
	resp = &types.StoreBecomeRes{}
	apiRpcRes := &apiuser.UserInfoRes{}
	rpcRes := &store.CreateStoreRes{}
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
		}
	}()
	apiRpcRes, err = l.svcCtx.ApiRpcCl.User.Info(ctx, &apiuser.UserInfoReq{
		Token: token,
	})
	if err != nil || apiRpcRes.Result.Code != xcode.RESPONSE_SUCCESS {
		rpcRes.Result.Code = apiRpcRes.Result.Code
		rpcRes.Result.Message = apiRpcRes.Result.Message
		rpcRes.Result.ErrMsg = apiRpcRes.Result.ErrMsg
		return
	}
	rpcRes, err = l.svcCtx.StoreRpcCl.Become.CreateStore(ctx, &store.CreateStoreReq{
		UserId: apiRpcRes.UserId,
		Name:   req.Name,
		Desc:   req.Desc,
	})
	return
}
