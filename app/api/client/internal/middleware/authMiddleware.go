package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store/app/api/client/internal/types"
	"store/app/api/rpc/api/apitoken"
	"store/app/api/rpc/api/apiuser"
	"store/pkg/xcode"
)

type AuthMiddleware struct {
	userClient apiuser.ApiUser
	authClient apitoken.ApiToken
}

func NewAuthMiddleware(u apiuser.ApiUser, a apitoken.ApiToken) *AuthMiddleware {
	return &AuthMiddleware{userClient: u, authClient: a}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			code = xcode.RESPONSE_SUCCESS
			msg  = ""
		)
		auth := r.Header.Get("Authorization")
		if auth == "" {
			code, msg = xcode.GetCodeMessage(xcode.RESPONSE_UNAUTHORIZED)
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  "authorization header is required",
				Code:    code,
				Message: msg,
				Data:    map[string]interface{}{},
			})
			return
		}

		auth = auth[len("Bearer "):]
		ctx := context.Background()
		apiRpcRes, err := m.authClient.CheckAuth(ctx, &apitoken.AuthReq{
			Token: auth,
		})
		if err != nil {
			code, msg = xcode.GetCodeMessage(xcode.RESPONSE_FAIL)
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  "authorization fail",
				Code:    code,
				Message: msg,
				Data:    map[string]interface{}{},
			})
			return
		} else if apiRpcRes.Result.Code != xcode.RESPONSE_SUCCESS {
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  apiRpcRes.Result.ErrMsg,
				Code:    apiRpcRes.Result.Code,
				Message: apiRpcRes.Result.Message,
				Data:    map[string]interface{}{},
			})
			return
		}
		r.Header.Set("Authorization", auth)
		next(w, r)
	}
}
