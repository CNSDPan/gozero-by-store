package store

import (
	"net/http"
	"store/pkg/response"
	"store/pkg/xcode"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store/app/api/client/internal/logic/store"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
)

func InitChatLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitChatLogReq
		if err := httpx.Parse(r, &req); err != nil {
			code, msg := xcode.GetCodeMessage(xcode.RESPONSE_NOT_FOUND)
			response.Response(w, r, code, msg, nil, err.Error())
			return
		}

		l := store.NewInitChatLogLogic(r.Context(), svcCtx)
		res, resp, _ := l.InitChatLog(&req)
		response.Response(w, r, res.Code, res.Message, resp, res.ErrMsg)
	}
}
