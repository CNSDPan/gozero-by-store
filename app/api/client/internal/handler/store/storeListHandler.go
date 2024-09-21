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
		l := store.NewStoreListLogic(r.Context(), svcCtx)
		res, resp, err := l.StoreList()
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
