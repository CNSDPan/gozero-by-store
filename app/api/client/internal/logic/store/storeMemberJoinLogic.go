package store

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
	"store/app/rpc/api/client/apiuser"
	"store/app/rpc/store/pb/store"
	"store/pkg/xcode"
)

type StoreMemberJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreMemberJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreMemberJoinLogic {
	return &StoreMemberJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreMemberJoinLogic) StoreMemberJoin(req *types.StoreMemberJoinReq, token string) (res *types.Response, resp *types.StoreMemberJoinRes, err error) {
	code := ""
	res = &types.Response{}
	rpcRes := &store.JoinStoreMemberRes{}
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
		}
	}()
	if userRpcRes, err = l.svcCtx.ApiRpcCl.User.Info(ctx, &apiuser.UserInfoReq{
		Token: token,
	}); err != nil {

		rpcRes.Result = &store.Response{
			ErrMsg:  userRpcRes.Result.ErrMsg,
			Code:    userRpcRes.Result.Code,
			Message: userRpcRes.Result.Message,
		}
		return
	}
	rpcRes, err = l.svcCtx.StoreRpcCl.Become.JoinStoreMember(ctx, &store.JoinStoreMemberReq{
		StoreId: req.StoreId,
		UserId:  userRpcRes.UserId,
	})
	return
}
