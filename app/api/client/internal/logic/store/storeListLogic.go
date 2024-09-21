package store

import (
	"context"
	"store/app/api/rpc/api/apistore"
	"store/pkg/xcode"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreListLogic {
	return &StoreListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreListLogic) StoreList() (res *types.Response, resp *types.StoreListRes, err error) {
	code := ""
	res = &types.Response{}
	resp = &types.StoreListRes{}
	rpcRes := &apistore.StoreListRes{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func() {
		if err != nil {
			if code == "" {
				l.Logger.Errorf("%s 系统错误 fail:%v", l.svcCtx.Config.ServiceName, err.Error())
				res.Code, res.Message = xcode.GetCodeMessage(xcode.RESPONSE_FAIL)
			} else {
				res.ErrMsg = err.Error()
				res.Code, res.Message = xcode.GetCodeMessage(code)
			}
		} else {
			res.ErrMsg = rpcRes.Result.ErrMsg
			res.Code = rpcRes.Result.Code
			res.Message = rpcRes.Result.Message

			resp.Limit = rpcRes.Data.Limit
			resp.Offset = rpcRes.Data.Offset
			resp.Page = rpcRes.Data.Page
			resp.Current = rpcRes.Data.Current
			resp.Total = rpcRes.Data.Total
			resp.Rows = rpcRes.Data.Rows
		}
	}()
	rpcRes, err = l.svcCtx.ApiRpcCl.Store.List(ctx, &apistore.StoreListReq{})
	return
}
