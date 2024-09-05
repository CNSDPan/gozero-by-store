package store

import (
	"context"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreUserListLogic {
	return &StoreUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreUserListLogic) StoreUserList(req *types.StoreUsersReq) (resp *types.StoreUsersRes, err error) {
	// todo: add your logic here and delete this line

	return
}
