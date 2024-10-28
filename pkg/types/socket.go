package types

// SocketMsg 广播消息结构体
type SocketMsg struct {
	Operate       int           `json:"operate"`
	Method        string        `json:"method"`
	StoreId       int64         `json:"storeId,string"`
	SendUserId    int64         `json:"sendUserId,string"`
	ReceiveUserId int64         `json:"receiveUserId,string"`
	Extend        string        `json:"extend"`
	Body          SocketMsgBody `json:"body"`
}
type SocketMsgBody struct {
	Operate      int    `json:"operate"`
	Method       string `json:"method"`
	ResponseTime string `json:"responseTime"`
	Event        Event  `json:"event"`
}

/******************Event 请求&响应结构*********************/

// Event 请求&响应结构
type Event struct {
	Params interface{} `json:"params"` // 请求参数
	Data   interface{} `json:"data"`   // 响应参数
}

// DataByNormal 进入普通消息响应Data结构
type DataByNormal struct {
	StoreId       int64  `json:"storeId,string"`       // 店铺
	SendUserId    int64  `json:"sendUserId,string"`    // 消息来源用户
	ReceiveUserId int64  `json:"receiveUserId,string"` // 消息指定推送用户
	Message       string `json:"message"`              // 消息内容
}

/******************Event 请求&响应结构*********************/
