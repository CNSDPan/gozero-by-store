package ws

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store/app/chat/socket/internal/logic/ws"
	"store/app/chat/socket/internal/svc"
	"store/app/chat/socket/internal/types"
	"store/pkg/xcode"
)

func SocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		req := types.ConnectReq{
			Token: auth,
		}
		l := ws.NewSocketLogic(r.Context(), svcCtx)
		res, _, _ := l.Socket(&req, w, r)
		if res.Code != xcode.RESPONSE_SUCCESS {
			httpx.ErrorCtx(r.Context(), w, errors.New(res.Message))
			return
		}
	}
}
