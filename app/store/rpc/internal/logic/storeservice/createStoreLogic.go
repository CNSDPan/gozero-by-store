package storeservicelogic

import (
	"context"

	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CreateStoreLogic) CreateStore(in *store.CreateStoreReq) (*store.CreateStoreRes, error) {
	// todo: add your logic here and delete this line

	return &store.CreateStoreRes{}, nil
}
