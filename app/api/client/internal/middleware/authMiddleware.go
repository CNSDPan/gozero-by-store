package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"store/app/api/client/internal/types"
	"store/pkg/xcode"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			code = ""
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
		r.Header.Set("Authorization", auth)
		next(w, r)
	}
}
