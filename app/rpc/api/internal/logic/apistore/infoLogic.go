package apistorelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/rpc/api/internal/svc"
	"store/app/rpc/api/pb/api"
	"store/db/dao/model"
	"store/pkg/xcode"
	"strconv"
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
// @Desc：店铺详情
// @param：in
// @return：res
// @return：err
func (l *InfoLogic) Info(in *api.StoreInfoReq) (res *api.StoreInfoRes, err error) {
	var (
		e          error
		code       = xcode.RESPONSE_SUCCESS
		storeCache = map[string]string{}
		storeInfo  = model.StoresInfoApi{}
	)
	res = &api.StoreInfoRes{
		Result: &api.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取店铺信息 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	if in.UserId != 0 {
		in.StoreId = l.svcCtx.MysqlQuery.StoreUser.GetStoreIdByUserId(in.UserId)
	}

	if storeCache, e = l.svcCtx.CacheConnApi.Store.GetInfo(in.StoreId); e != nil {
		code = xcode.STORE_INFO_FAIL
		return res, nil
	}
	if len(storeCache) == 0 {
		storeInfo, e = l.svcCtx.MysqlQuery.Store.GetFromStoreIDApi(in.StoreId)
		if storeInfo.StoreId == 0 {
			code = xcode.STORE_INFO
			return res, nil
		}
		if e != nil {
			code = xcode.STORE_INFO_FAIL
			return res, nil
		}
		if e = l.svcCtx.CacheConnApi.Store.SetStoreAndStoreUser(storeInfo.StoreId, map[string]interface{}{
			"storeUserId": storeInfo.StoreUserId,
			"storeId":     storeInfo.StoreId,
			"name":        storeInfo.Name,
			"avatar":      storeInfo.Avatar,
			"desc":        storeInfo.Desc,
			"contacts":    storeInfo.Contacts,
		}, storeInfo.StoreUserId, map[string]interface{}{
			"storeUserId": storeInfo.StoreUserId,
			"storeId":     storeInfo.StoreId,
			"userId":      storeInfo.UserId,
		}, l.svcCtx.Config.CacheSeconds); e != nil {
			l.Logger.Errorf("%s 存储门店和店长数据缓存 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			e = nil
		}
		res.StoreId = storeInfo.StoreId
		res.Name = storeInfo.Name
		res.Avatar = storeInfo.Avatar
		res.Desc = storeInfo.Desc
		res.Contacts = &storeInfo.Contacts
		res.StoreUserId = storeInfo.StoreUserId
		res.UserId = storeInfo.UserId
	} else {
		res.StoreId, _ = strconv.ParseInt(storeCache["storeId"], 10, 64)
		res.Name = storeCache["name"]
		res.Avatar = storeCache["avatar"]
		res.Desc = storeCache["desc"]
		contacts, _ := strconv.ParseInt(storeCache["contacts"], 10, 64)
		res.Contacts = &contacts
		res.StoreUserId, _ = strconv.ParseInt(storeCache["storeUserId"], 10, 64)
		res.UserId, _ = strconv.ParseInt(storeCache["userId"], 10, 64)
		l.Logger.Errorf("用户ID：%+v %+v", res.UserId, storeCache["userId"])
	}
	return res, nil
}
