package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
	"store/app/rpc/api/client/apiuser"
	"store/pkg/xcode"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (res *types.Response, resp *types.LoginRes, err error) {
	code := ""
	res = &types.Response{}
	resp = &types.LoginRes{}
	login := &apiuser.UserLoginRes{}
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
			res.ErrMsg = login.Result.ErrMsg
			res.Code = login.Result.Code
			res.Message = login.Result.Message
			resp.UserId = login.UserId
			resp.Token = login.Token
		}
	}()
	if !regexp.MustCompile("^1[345789]{1}\\d{9}$").MatchString(fmt.Sprintf("%d", req.Mobile)) {
		code = xcode.USER_CREAT_MOBILE_RULER
		err = errors.New("mobile is invalid")
		return
	}
	login, err = l.svcCtx.ApiRpcCl.User.Login(ctx, &apiuser.UserLoginReq{
		Mobile:    req.Mobile,
		Password:  req.Password,
		JwtSecret: l.svcCtx.Config.Auth.AccessSecret,
		Seconds:   l.svcCtx.Config.Auth.AccessExpire,
	})
	return
}
