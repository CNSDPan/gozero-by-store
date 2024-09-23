package apiuserlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/app/api/rpc/api/apiuser"
	"store/app/user/model/sqls"
	"store/pkg/biz"
	"store/pkg/xcode"
	"strconv"

	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Info
// @Desc：用户详情
// @param：in
// @return：res
// @return：err
func (l *InfoLogic) Info(in *api.UserInfoReq) (res *api.UserInfoRes, err error) {
	var (
		e         error
		code      = "200"
		userId    int64
		userIdStr string
		info      = sqls.UsersApi{}
		infoM     = make(map[string]string)
		b         []byte
	)
	res = &apiuser.UserInfoRes{
		Result: &apiuser.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取用户信息 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		} else {
			res.UserId = info.UserId
			res.Mobile = info.Mobile
			res.Name = info.Name
			res.Avatar = info.Avatar
		}
	}()
	if in.UserId != 0 {
		userIdStr = strconv.FormatInt(in.UserId, 10)
		goto GetCache
	}
	userIdStr, e = l.svcCtx.BizConn.Get(l.ctx, fmt.Sprintf("%s%s", biz.Biz_Key_USER_TOKEN, in.Token)).Result()
	if e != nil && e != redis.Nil {
		code = xcode.USER_TOKEN_GET
		goto Result
	} else if userIdStr == "" {
		err = errors.New("")
		code = xcode.USER_TOKEN_FAIL
		goto Result
	}
GetCache:
	userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	infoM, e = l.svcCtx.CacheConnApi.User.GetInfo(userId)
	if e != nil || len(infoM) == 0 {
		info, e = l.svcCtx.UserModel.GetFromUserIDApi(userId)
		if e != nil {
			code = xcode.USER_INFO_FAIL
			goto Result
		} else if info.UserId == 0 {
			code = xcode.USER_INFO_FAIL
			goto Result
		}
	} else {
		b, e = jsonx.Marshal(infoM)
		if e != nil {
			code = xcode.USER_INFO_ERR
			goto Result
		}
		e = jsonx.Unmarshal(b, &info)
		if e != nil {
			code = xcode.USER_INFO_ERR
			goto Result
		}
	}
Result:
	return res, nil
}
