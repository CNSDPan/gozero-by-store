package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"store/app/rpc/user/pb/user"
	"store/pkg/xcode"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterReq) (res *types.Response, resp *types.RegisterRes, err error) {
	code := ""
	register := &user.RegisterRes{}
	res = &types.Response{}
	resp = &types.RegisterRes{}
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
			res.ErrMsg = register.Result.ErrMsg
			res.Code = register.Result.Code
			res.Message = register.Result.Message
			resp.UserId = register.UserId
			resp.Token = register.Token
		}
	}()
	if !regexp.MustCompile("^1[345789]{1}\\d{9}$").MatchString(fmt.Sprintf("%d", req.Mobile)) {
		code = xcode.USER_CREAT_MOBILE_RULER
		err = errors.New("mobile is invalid")
		return
	}
	register, err = l.svcCtx.UserRpcCl.Register.Register(ctx, &user.RegisterReq{
		Mobile:    req.Mobile,
		Name:      req.Name,
		Password:  req.Password,
		JwtSecret: l.svcCtx.Config.Auth.AccessSecret,
		Seconds:   l.svcCtx.Config.Auth.AccessExpire,
	})
	return
}
