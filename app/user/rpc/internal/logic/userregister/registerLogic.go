package userregisterlogic

import (
	"context"
	"store/app/user/model/sqls"
	"store/pkg/xcode"

	"store/app/user/rpc/internal/svc"
	"store/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	var (
		code   = "200"
		result = &user.RegisterRes{}
		info   = sqls.UsersApi{}
		err    error
	)
	defer func() {
		result.Result.Code, result.Result.Message = xcode.GetCodeMessage(code)
		if err != nil {
			l.Logger.Errorf("%s 注册用户失败 fail:%s", l.svcCtx.Config.ServiceName, err.Error())
			result.Result.ErrMsg = err.Error()
		}
	}()
	info, err = l.svcCtx.UserModel.GetUserApi(sqls.Users{
		Mobile: int(in.Mobile),
	})
	if err != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	if info.UserID > 0 {
		code = xcode.USER_CREAT_MOBILE_FAIL
		goto Result
	}
	err = l.svcCtx.UserModel.CreatUser(sqls.Users{
		UserID:   0,
		Mobile:   int(in.Mobile),
		Name:     in.Name,
		Password: in.Password,
		Avatar:   "",
	})
	if err != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}

Result:
	return result, nil
}
