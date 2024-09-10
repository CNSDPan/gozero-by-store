package store

import (
	"context"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreBecomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreBecomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreBecomeLogic {
	return &StoreBecomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreBecomeLogic) StoreBecome(req *types.StoreBecomeReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
