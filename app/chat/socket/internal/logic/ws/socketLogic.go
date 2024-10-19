package ws

import (
	"context"
	"net/http"
	"store/app/api/rpc/api/apistore"
	"store/app/api/rpc/api/apiuser"
	"store/app/chat/socket/internal/svc"
	"store/app/chat/socket/internal/types"
	"store/app/chat/socket/server"
	"store/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type SocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SocketLogic {
	return &SocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SocketLogic) Socket(req *types.ConnectReq, w http.ResponseWriter, r *http.Request) (res *types.Response, resp *types.UserInfoRes, err error) {
	code := ""
	res = &types.Response{}
	resp = &types.UserInfoRes{}
	rpcRes := &apiuser.UserInfoRes{}
	storeRes := &apiuser.StoreListRes{}
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
		}
	}()
	rpcRes, err = l.svcCtx.ApiRpcCl.User.Info(ctx, &apiuser.UserInfoReq{
		Token: req.Token,
	})
	if err != nil || rpcRes.Result.Code != xcode.RESPONSE_SUCCESS {
		return
	}
	storeRes, err = l.svcCtx.ApiRpcCl.Store.List(ctx, &apiuser.StoreListReq{
		UserId: rpcRes.UserId,
	})
	if err != nil || storeRes.Result.Code != xcode.RESPONSE_SUCCESS {
		rpcRes.Result.Code = storeRes.Result.Code
		rpcRes.Result.Message = storeRes.Result.Message
		rpcRes.Result.ErrMsg = storeRes.Result.ErrMsg
		return
	}
	storeIdsRes, err := l.svcCtx.ApiRpcCl.Store.MyAllStore(ctx, &apistore.MyAllStoreIdReq{UserId: rpcRes.UserId})
	if err != nil || storeRes.Result.Code != xcode.RESPONSE_SUCCESS {
		rpcRes.Result.Code = storeRes.Result.Code
		rpcRes.Result.Message = storeRes.Result.Message
		rpcRes.Result.ErrMsg = storeRes.Result.ErrMsg
		return
	}

	wsConnect := server.NewConnect()
	wsConn, err := wsConnect.Run(w, r, l.svcCtx.WsServer.Option.MaxMessageSize, l.svcCtx.WsServer.Option.PongPeriod)
	if err != nil {
		rpcRes.Result.Code, rpcRes.Result.Message = xcode.GetCodeMessage(xcode.SOCKET_UPGRADER_FAIL)
		return
	}

	wsClient := server.NewClient(wsConn, l.svcCtx.Node.Generate().Int64(), rpcRes.UserId, resp.Name)

	// 这里会阻塞,直到socket断开
	l.svcCtx.WsServer.Run(wsClient, storeIdsRes.StoreId)
	return
}
