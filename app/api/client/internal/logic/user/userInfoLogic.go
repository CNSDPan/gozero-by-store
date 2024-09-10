package user

import (
	"context"
	"store/app/api/rpc/api/apiuser"
	"store/pkg/xcode"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (res *types.Response, resp *types.UserInfoRes, err error) {
	code := ""
	res = &types.Response{}
	resp = &types.UserInfoRes{}
	rpcRes := &apiuser.UserInfoRes{}
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
			resp.UserId = rpcRes.UserId
			resp.Mobile = int64(rpcRes.Mobile)
			resp.Name = rpcRes.Name
			resp.Avatar = rpcRes.Avatar
		}
	}()
	rpcRes, err = l.svcCtx.ApiRpcCl.User.Info(ctx, &apiuser.UserInfoReq{
		UserId: req.UserId,
		Token:  req.Token,
	})
	return
}
