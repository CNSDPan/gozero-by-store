package store

import (
	"net/http"
	"store/pkg/xcode"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store/app/api/client/internal/logic/store"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
)

func StoreInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StoreInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			code, msg := xcode.GetCodeMessage(xcode.RESPONSE_NOT_FOUND)
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  err.Error(),
				Code:    code,
				Message: msg,
				Data:    map[string]interface{}{},
			})
			return
		}

		l := store.NewStoreInfoLogic(r.Context(), svcCtx)
		res, resp, err := l.StoreInfo(&req)
		if err != nil || res.Code != xcode.RESPONSE_SUCCESS {
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  res.ErrMsg,
				Code:    res.Code,
				Message: res.Message,
				Data:    map[string]interface{}{},
			})
		} else {
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  res.ErrMsg,
				Code:    res.Code,
				Message: res.Message,
				Data:    resp,
			})
		}
	}
}
