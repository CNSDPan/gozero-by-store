package store

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store/app/api/client/internal/logic/store"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
	"store/pkg/xcode"
)

func StoreListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StoreListReq
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
		l := store.NewStoreListLogic(r.Context(), svcCtx)
		res, resp, err := l.StoreList(&req)
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
