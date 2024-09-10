package apiuserlogic

import (
	"context"
	"fmt"
	"store/app/api/rpc/api/apiuser"
	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"
	"store/app/user/model/sqls"
	"store/pkg/biz"
	"store/pkg/jwt"
	"store/pkg/util"
	"store/pkg/xcode"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *api.UserLoginReq) (res *api.UserLoginRes, err error) {
	var (
		e error
	)
	code := "200"
	info := sqls.Users{}

	token := ""
	res = &apiuser.UserLoginRes{
		Result: &apiuser.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 用户登录失败 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		} else {
			res.UserId = info.UserID
			res.Token = token
		}
	}()
	info, e = l.svcCtx.UserModel.GetFromMobile(int32(in.Mobile))
	if info.UserID == 0 {
		code = xcode.USER_INFO_FAIL
		goto Result
	}
	if ok := util.CheckPasswordHash(in.Password, info.Password); !ok {
		code = xcode.USER_INFO_FAIL
		goto Result
	}
	token, e = jwt.GetJwtToken(in.JwtSecret, time.Now().Unix(), in.Seconds, map[string]interface{}{
		"userId": info.UserID,
	})
	if e != nil {
		code = xcode.USER_TOKEN_CREATE
		goto Result
	}
	e = l.svcCtx.CacheConnApi.User.SetInfo(info.UserID, map[string]interface{}{
		"userId": info.UserID,
		"mobile": info.Mobile,
		"name":   info.Name,
		"avatar": info.Avatar,
	}, in.Seconds)
	if e != nil {
		code = xcode.USER_SET_INFOCACHE_FAIL
		l.Logger.Errorf("%s 存储用户缓存 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
		goto Result
	}
	e = l.svcCtx.BizConn.Set(l.ctx, fmt.Sprintf("%s%s", biz.Biz_Key_USER_TOKEN, token), info.UserID, time.Duration(in.Seconds)*time.Second).Err()
	if e != nil {
		code = xcode.USER_SET_INFOCACHE_FAIL
		l.Logger.Errorf("%s 存储用户token fail:%s", l.svcCtx.Config.ServiceName, e.Error())
	}
Result:
	return res, nil
}
