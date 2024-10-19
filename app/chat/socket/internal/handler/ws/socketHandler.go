package ws

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store/app/chat/socket/internal/logic/ws"
	"store/app/chat/socket/internal/svc"
	"store/app/chat/socket/internal/types"
	"store/pkg/response"
	"store/pkg/xcode"
)

func SocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConnectReq
		var (
			code = xcode.RESPONSE_SUCCESS
			msg  = ""
		)
		if err := httpx.Parse(r, &req); err != nil {
			code, msg = xcode.GetCodeMessage(xcode.RESPONSE_NOT_FOUND)
			httpx.ErrorCtx(r.Context(), w, err)
			response.Response(w, r, code, msg, map[string]interface{}{}, err.Error())
			return
		}
		l := ws.NewSocketLogic(r.Context(), svcCtx)
		res, _, _ := l.Socket(&req, w, r)
		if res.Code != xcode.RESPONSE_SUCCESS {
			response.Response(w, r, res.Code, res.Message, map[string]interface{}{}, res.ErrMsg)
			return
		}
	}
}
