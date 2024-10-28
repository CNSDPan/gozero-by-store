package biz

type Biz_Key string

// 用户redis业务处理 key
const (
	// Biz_Key_USER_TOKEN 用户token
	Biz_Key_USER_TOKEN Biz_Key = "userToken:"
)

const (
	// SOCKET_PUB_SUB_BROADCAST_NORMAL_CHAN_KEY 通讯消息的发布订阅的key
	SOCKET_PUB_SUB_BROADCAST_NORMAL_CHAN_KEY = "socket:broadcast:normal"
)
