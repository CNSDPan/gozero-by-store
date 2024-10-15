package apitokenlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"store/pkg/biz"
	"store/pkg/xcode"

	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAuthLogic) CheckAuth(in *api.AuthReq) (res *api.AuthRes, err error) {
	var (
		e         error
		code      = xcode.RESPONSE_SUCCESS
		userIdStr string
	)
	res = &api.AuthRes{
		Result: &api.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取用户token fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	userIdStr, e = l.svcCtx.BizConn.Get(l.ctx, fmt.Sprintf("%s%s", biz.Biz_Key_USER_TOKEN, in.Token)).Result()
	if e != nil && e != redis.Nil {
		code = xcode.USER_TOKEN_GET
	} else if userIdStr == "" {
		e = errors.New("")
		code = xcode.USER_TOKEN_FAIL
	}
	return res, nil
}
