package types

type JSONResponseCtx struct {
	ErrMsg  string `json:"errMsg"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
