package apistorelogic

import (
	"context"
	"store/app/api/rpc/api/apistore"
	sqlsStore "store/app/store/model/sqls"
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

func (l *InfoLogic) Info(in *api.StoreInfoReq) (res *api.StoreInfoRes, err error) {
	var (
		e          error
		code       = xcode.RESPONSE_SUCCESS
		storeCache = map[string]string{}
		storeInfo  = sqlsStore.StoresInfoApi{}
	)
	res = &api.StoreInfoRes{
		Result: &apistore.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取店铺信息 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	if storeCache, e = l.svcCtx.CacheConnApi.Store.GetInfo(in.StoreId); e != nil {
		code = xcode.STORE_INFO_FAIL
		return res, nil
	}
	if len(storeCache) == 0 {
		storeInfo, e = l.svcCtx.StoreModel.StoresMgr.GetFromStoreIDApi(in.StoreId)
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
	}
	return res, nil
}
