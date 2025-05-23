package apistorelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/clause"
	"store/app/rpc/api/internal/svc"
	"store/app/rpc/api/pb/api"
	mysqlModel "store/db/dao/model"
	mysqlQuery "store/db/dao/query"
	"store/pkg/xcode"
	"strconv"
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
		e             error
		code          = xcode.RESPONSE_SUCCESS
		storeIds      = make([]int64, 0)
		whereClause   interface{}
		storeIdsInter = make([]interface{}, 0)
	)
	res = &api.StoreListRes{
		Result: &api.Response{},
		Data:   &api.StoresMap{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 获取店铺列表 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	if in.UserId > 0 {
		storeIds = l.svcCtx.MysqlQuery.StoreMember.GetStoreIdsByUserId(in.UserId)
		for _, storeId := range storeIds {
			storeIdsInter = append(storeIdsInter, storeId)
		}
		whereClause = clause.Where{Exprs: []clause.Expression{
			clause.IN{
				Column: "store_id",
				Values: storeIdsInter,
			},
		}}
	}
	items, e := l.svcCtx.MysqlQuery.Store.SelectPageApi(
		mysqlQuery.NewPage(100, 0),
		whereClause,
		l.svcCtx.MysqlQuery.Store.WithStatusEnable(),
	)
	if e != nil {
		code = xcode.STORE_ITEM
		return res, nil
	} else {
		rows := make([]*api.StoreItem, len(items.GetRecords().([]mysqlModel.StoresApi)))
		for k, item := range items.GetRecords().([]mysqlModel.StoresApi) {
			rows[k] = &api.StoreItem{
				StoreId:  strconv.FormatInt(item.StoreId, 10),
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
