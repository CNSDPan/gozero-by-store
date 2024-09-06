package user

import (
	"context"
	"store/app/user/rpc/pb/user"
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
	var (
		register = &user.RegisterRes{}
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func() {
		if err != nil {
			res.ErrMsg = err.Error()
			res.Code, res.Message = xcode.GetCodeMessage(xcode.RESPONSE_FAIL)
		} else {
			res.ErrMsg = register.Result.ErrMsg
			res.Code = register.Result.Code
			res.Message = register.Result.Message
			resp.UserId = register.UserId
			resp.Token = register.Token
		}
	}()
	register, err = l.svcCtx.UserRpcCl.Register.Register(ctx, &user.RegisterReq{
		Mobile:   req.Mobile,
		Name:     req.Name,
		Password: req.Password,
	})
	return
}
