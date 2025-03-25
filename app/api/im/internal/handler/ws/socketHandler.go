package ws

import (
	"errors"
	"net/http"
	"store/pkg/xcode"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store/app/api/im/internal/logic/ws"
	"store/app/api/im/internal/svc"
	"store/app/api/im/internal/types"
)

// SocketHandler 连接
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
