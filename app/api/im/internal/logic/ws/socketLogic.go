package ws

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"store/app/api/im/internal/svc"
	"store/app/api/im/internal/types"
	"store/app/api/im/server"
	"store/app/rpc/api/client/apistore"
	"store/app/rpc/api/client/apiuser"
	"store/pkg/xcode"
)

type SocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// socket 连接
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
	storeIdsRes := &apiuser.MyAllStoreIdRes{}
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
	storeIdsRes, err = l.svcCtx.ApiRpcCl.Store.MyAllStore(ctx, &apistore.MyAllStoreIdReq{UserId: rpcRes.UserId})
	if err != nil || storeIdsRes.Result.Code != xcode.RESPONSE_SUCCESS {
		rpcRes.Result.Code = storeIdsRes.Result.Code
		rpcRes.Result.Message = storeIdsRes.Result.Message
		rpcRes.Result.ErrMsg = storeIdsRes.Result.ErrMsg
		return
	}
	wsConnect := server.NewConnect()
	wsConn, err := wsConnect.Run(w, r, l.svcCtx.WsServer.Option.MaxMessageSize, l.svcCtx.WsServer.Option.PongPeriod)
	if err != nil {
		rpcRes.Result.Code, rpcRes.Result.Message = xcode.GetCodeMessage(xcode.SOCKET_UPGRADER_FAIL)
		return
	}

	wsClient := server.NewClient(wsConn, l.svcCtx.Node.Generate().Int64(), rpcRes.UserId, rpcRes.Name, storeIdsRes.StoreId)
	// 这里会阻塞,直到socket断开
	l.svcCtx.WsServer.Run(wsClient, storeIdsRes.StoreId)
	return
}
