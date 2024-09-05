package store

import (
	"context"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreInfoLogic {
	return &StoreInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreInfoLogic) StoreInfo(req *types.StoreInfoReq) (resp *types.StoreInfoRes, err error) {
	// todo: add your logic here and delete this line

	return
}
