package storebecomelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/store/model/sqls"
	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"
	"store/pkg/xcode"
	"time"
)

type CreateStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStoreLogic {
	return &CreateStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStoreLogic) CreateStore(in *store.CreateStoreReq) (res *store.CreateStoreRes, err error) {
	var (
		e         error
		code      = "200"
		storeInfo = sqls.StoreUsersApi{}
	)
	res = &store.CreateStoreRes{
		Result: &store.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 创建门店 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	storeInfo, e = l.svcCtx.StoreUserModel.GetFromUserIdApi(in.UserId)
	if storeInfo.StoreUserId > 0 {
		code = xcode.STORE_CREATED
		return res, err
	}
	if l.svcCtx.StoreModel.GetFromNameApi(in.Name).StoreId > 0 {
		code = xcode.STORE_CREATED_NAME
		return res, err
	}
	stores := sqls.Stores{
		StoreId: l.svcCtx.Node.Generate().Int64(),
		Name:    in.Name,
		Avatar:  "",
		Desc:    in.Desc,
	}
	time.Sleep(5 * time.Millisecond)
	storeUsers := sqls.StoreUsers{
		StoreUserId: l.svcCtx.Node.Generate().Int64(),
		StoreId:     stores.StoreId,
		UserId:      in.UserId,
	}
	time.Sleep(5 * time.Millisecond)
	storeMember := sqls.StoreMember{
		StoreMemberId: l.svcCtx.Node.Generate().Int64(),
		StoreId:       stores.StoreId,
		UserId:        in.UserId,
	}
	if e = l.svcCtx.StoreUserModel.CreateStoreUser(storeMember, storeUsers, stores); e != nil {
		code = xcode.STORE_CREAT
		return res, err
	}

	if e = l.svcCtx.CacheConnApi.Store.SetStoreAndStoreUser(stores.StoreId, map[string]interface{}{
		"storeUserId": storeUsers.StoreUserId,
		"storeId":     stores.StoreId,
		"name":        stores.Name,
		"avatar":      stores.Avatar,
		"desc":        stores.Desc,
		"contacts":    0,
	}, storeUsers.StoreUserId, map[string]interface{}{
		"storeUserId": storeUsers.StoreUserId,
		"storeId":     stores.StoreId,
		"userId":      in.UserId,
	}, l.svcCtx.Config.CacheSeconds); e != nil {
		l.Logger.Errorf("%s 存储门店和店长数据缓存 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
		e = nil
	}
	return res, err
}
