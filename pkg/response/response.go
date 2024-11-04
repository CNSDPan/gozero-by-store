package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	ErrMsg  string      `json:"errMsg"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response
// @Desc： 统一格式的 body 响应
// @param：w
// @param：r
// @param：code
// @param：msg
// @param：resp
// @param：errMsg
// @param：err
func Response(w http.ResponseWriter, r *http.Request, code string, msg string, resp interface{}, errMsg string) {
	var body Body
	body.ErrMsg = errMsg
	body.Code = code
	body.Message = msg
	if resp == nil {
		body.Data = map[string]interface{}{}
	} else {
		body.Data = resp
	}
	httpx.OkJsonCtx(r.Context(), w, body)
}
