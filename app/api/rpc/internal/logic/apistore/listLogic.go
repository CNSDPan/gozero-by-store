package apistorelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/api/rpc/api/apistore"
	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"
	sqlsStore "store/app/store/model/sqls"
	"store/pkg/xcode"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// List
// @Desc：店铺列表
// @param：in
// @return：res
// @return：err
func (l *ListLogic) List(in *api.StoreListReq) (res *api.StoreListRes, err error) {
	var (
		e    error
		code = xcode.RESPONSE_SUCCESS
	)
	res = &api.StoreListRes{
		Result: &apistore.Response{},
		Data:   &apistore.StoresMap{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取店铺列表 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	items, e := l.svcCtx.StoreModel.StoresMgr.SelectPageApi(
		sqlsStore.NewPage(100, 0),
		l.svcCtx.StoreModel.StoresMgr.WithStatusEnable(),
	)
	if e != nil {
		code = xcode.STORE_ITEM
		return res, nil
	} else {
		rows := make([]*api.StoreItem, len(items.GetRecords().([]sqlsStore.StoresApi)))
		for k, item := range items.GetRecords().([]sqlsStore.StoresApi) {
			rows[k] = &api.StoreItem{
				StoreId:  item.StoreId,
				Name:     item.Name,
				Avatar:   item.Avatar,
				Desc:     item.Desc,
				Contacts: &item.Contacts,
			}
		}
		res.Data.Limit = items.GetSize()
		res.Data.Offset = items.GetCurrent()
		res.Data.Page = items.GetPages()
		res.Data.Current = items.GetCurrent()
		res.Data.Total = items.GetTotal()
		res.Data.Rows = rows
	}

	return res, nil
}
