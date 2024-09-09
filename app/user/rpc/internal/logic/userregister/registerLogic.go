package userregisterlogic

import (
	"context"
	"store/app/user/model/sqls"
	"store/pkg/jwt"
	"store/pkg/xcode"
	"time"

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

func (l *RegisterLogic) Register(in *user.RegisterReq) (res *user.RegisterRes, err error) {
	var (
		e      error
		userId int64
	)
	code := "200"
	info := sqls.UsersApi{}

	token := ""
	res = &user.RegisterRes{
		Result: &user.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 注册用户失败 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		} else {
			res.UserId = userId
			res.Token = token
		}
	}()
	info, e = l.svcCtx.UserModel.GetFromMobile(int32(in.Mobile))
	if info.UserID > 0 {
		code = xcode.USER_CREAT_MOBILE_FAIL
		goto Result
	}
	if e != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	userId = l.svcCtx.Node.Generate().Int64()
	e = l.svcCtx.UserModel.CreatUser(sqls.Users{
		UserID:   userId,
		Mobile:   int32(in.Mobile),
		Name:     in.Name,
		Password: in.Password,
		Avatar:   "",
	})
	if e != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	token, e = jwt.GetJwtToken(in.JwtSecret, time.Now().Unix(), in.Seconds, map[string]interface{}{
		"userId": userId,
	})

Result:
	return res, nil
}
