package xcode

import "fmt"

var codeMessage = map[string]string{
	RESPONSE_SUCCESS:           "success",
	RESPONSE_FAIL:              "服务器内部错误",
	RESPONSE_NOT_FOUND:         "请求资源不存在",
	RESPONSE_UNAUTHORIZED:      "缺少身份认证",
	RESPONSE_NOT_CODE:          "无定义code码",
	RESPONSE_REQUEST_TIME_FAIL: "缺少请求日期",
	RESPONSE_TOKEN_FAIL:        "无效token",
	RESPONSE_APPID_FAIL:        "无效APPID",
	RESPONSE_SECRET_FAIL:       "无效secret",
	RESPONSE_SIGN_FAIL:         "无效sign",

	SOCKET_UPGRADER_FAIL: "ws无效协议",

	/*****************用户模块 [10100 ~ 10199]**********************/
	USER_LOGIN_FAIL:         "用户账号密码不正确",
	USER_LOGINED:            "用户已登录",
	USER_INFO_FAIL:          "用户信息不存在",
	USER_ID_FAIL:            "用户ID不存在|错误",
	USER_TOKEN_FAIL:         "用户Token不存在|错误",
	USER_TOKEN_CREATE:       "用户Token创建错误",
	USER_TOKEN_GET:          "用户Token获取错误",
	USER_TOKEN_DEL:          "用户Token删除错误",
	USER_CREAT_FAIL:         "用户创建失败",
	USER_CREAT_MOBILE_FAIL:  "用户手机号已创建",
	USER_CREAT_MOBILE_RULER: "手机号无效",
	USER_SET_INFOCACHE_FAIL: "用户信息存储缓存错误",
	USER_INFO_ERR:           "用户信息获取异常",
	/***************************************/

	/*****************店铺模块 [10200 ~ 10299]**********************/
	STORE_CREAT:        "店铺创建失败",
	STORE_CREATED:      "已创建过店铺",
	STORE_INFO:         "店铺信息不存在",
	STORE_INFO_FAIL:    "店铺信息获取异常",
	STORE_CREATED_NAME: "店铺名已存在",
	STORE_ITEM:         "店铺列表获取异常",

	STORE_USER_INFO:         "店长信息不存在",
	STORE_USER_TOKEN_CREATE: "店长Token创建错误",
	STORE_USER_TOKEN_GET:    "店长Token获取错误",
	STORE_USER_TOKEN_DEL:    "店长Token删除错误",
	STORE_USER_TOKEN_RULER:  "店长Token不存在|错误",

	STORE_MEMBER_ITEM_FAIL: "店铺会员列表获取异常",
	STORE_MEMBER_JOIN_FAIL: "加入会员异常",
	STORE_MEMBER_JOINED:    "已是会员",
	STORE_MEMBER_JOIN_FULL: "店铺会员已满额",

	CHAT_LOG_INIT_FAIL: "店铺聊天记录列表获取异常",
	CHAT_LOG_GET_FAIL:  "聊天记录获取异常",
	/***************************************/

	/*****************socket消息模块 [20000 ~ 2099]**********************/
	SOCKET_BROADCAST_MSG_FAIL:   "广播消息异常",
	SOCKET_BROADCAST_MSG_STRING: "广播消息非文字",
	SOCKET_BROADCAST_MSG_PUB:    "广播消息发布异常",
	/***************************************/
}

// GetCodeMessage
// @Desc：获取code码对应message内容
// @param：code
// @return：string
// @return：string
func GetCodeMessage(code string) (string, string) {
	var (
		message string
		ok      bool
		c       string
	)
	if message, ok = codeMessage[code]; !ok {
		message = fmt.Sprintf("原始Code:%s ,更改Code:%s ,message:%s", code, RESPONSE_NOT_CODE, codeMessage[RESPONSE_NOT_CODE])
		c = RESPONSE_NOT_CODE
	} else {
		c = code
	}
	return c, message
}
