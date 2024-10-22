package middleware

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store/app/api/rpc/api/apitoken"
	"store/pkg/xcode"
)

type AuthMiddleware struct {
	authClient apitoken.ApiToken
}

func NewAuthMiddleware(a apitoken.ApiToken) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: a,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			_, msg := xcode.GetCodeMessage(xcode.RESPONSE_UNAUTHORIZED)
			httpx.ErrorCtx(r.Context(), w, errors.New(msg))

			return
		}
		auth := r.Header.Get("Authorization")[len("Bearer "):]
		ctx := context.Background()
		apiRpcRes, err := m.authClient.CheckAuth(ctx, &apitoken.AuthReq{
			Token: auth,
		})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} else if apiRpcRes.Result.Code != xcode.RESPONSE_SUCCESS {
			httpx.ErrorCtx(r.Context(), w, errors.New(apiRpcRes.Result.ErrMsg))
			return
		}
		r.Header.Set("Authorization", auth)
		next(w, r)
	}
}
