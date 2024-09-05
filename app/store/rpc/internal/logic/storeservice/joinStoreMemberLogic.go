package storeservicelogic

import (
	"context"

	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinStoreMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinStoreMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinStoreMemberLogic {
	return &JoinStoreMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinStoreMemberLogic) JoinStoreMember(in *store.JoinStoreMemberReq) (*store.JoinStoreMemberRes, error) {
	// todo: add your logic here and delete this line

	return &store.JoinStoreMemberRes{}, nil
}
