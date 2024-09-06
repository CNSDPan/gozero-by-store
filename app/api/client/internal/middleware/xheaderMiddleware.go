package middleware

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

type XHeaderMiddleware struct {
}

func NewXHeaderMiddleware() *XHeaderMiddleware {
	return &XHeaderMiddleware{}
}

func (m *XHeaderMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SetConfCORS(w) // 跨域
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		xApiV := r.Header.Get("X-API-Version")
		xSource := r.Header.Get("X-Source")
		xReqTime := r.Header.Get("X-Request-Time")
		if xApiV == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("X-API-Version header is required"))
			return
		}
		if xSource == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("X-Source header is required"))
			return
		}
		if xReqTime == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("X-Request-Time header is required"))
			return
		}
		if _, err := time.Parse(time.RFC3339, xReqTime); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("X-Request-Time header is invalid"))
			return
		}
		next(w, r)
	}
}

func SetConfCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
