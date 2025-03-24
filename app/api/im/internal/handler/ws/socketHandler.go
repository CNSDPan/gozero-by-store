package ws

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store/app/api/im/internal/logic/ws"
	"store/app/api/im/internal/svc"
	"store/app/api/im/internal/types"
)

// socket 连接
func SocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConnectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ws.NewSocketLogic(r.Context(), svcCtx)
		resp, err := l.Socket(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
