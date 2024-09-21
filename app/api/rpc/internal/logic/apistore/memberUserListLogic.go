package apistorelogic

import (
	"context"

	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUserListLogic {
	return &MemberUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberUserListLogic) MemberUserList(in *api.MemberUsersItemReq) (*api.MemberUsersItemRes, error) {
	// todo: add your logic here and delete this line

	return &api.MemberUsersItemRes{}, nil
}
