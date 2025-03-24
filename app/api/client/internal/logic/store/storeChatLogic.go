package store

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
	"store/app/rpc/api/client/apistore"
	"store/pkg/xcode"
)

type StoreChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreChatLogic {
	return &StoreChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreChatLogic) StoreChat(req *types.StoreChatReq) (res *types.Response, resp *types.StoreChatRes, err error) {
	res = &types.Response{}
	resp = &types.StoreChatRes{}
	rpcRes := &apistore.StoreChatRes{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func() {
		if err != nil {
			l.Logger.Errorf("%s 系统错误 fail:%v", l.svcCtx.Config.ServiceName, err.Error())
			res.Code, res.Message = xcode.GetCodeMessage(xcode.RESPONSE_FAIL)
		} else {
			res.ErrMsg = rpcRes.Result.ErrMsg
			res.Code = rpcRes.Result.Code
			res.Message = rpcRes.Result.Message

			resp.Limit = rpcRes.Data.Limit
			resp.Offset = rpcRes.Data.Offset
			resp.Page = rpcRes.Data.Page
			resp.Current = rpcRes.Data.Current
			resp.Total = rpcRes.Data.Total
			if len(rpcRes.Data.Rows) > 0 {
				resp.Rows = rpcRes.Data.Rows
			} else {
				resp.Rows = make([]interface{}, 0)
			}
		}
	}()
	rpcRes, err = l.svcCtx.ApiRpcCl.Store.StoresChat(ctx, &apistore.StoreChatReq{
		StoreId:   req.StoreId,
		Offset:    req.Offset,
		Timestamp: req.Timestamp,
	})
	return
}
