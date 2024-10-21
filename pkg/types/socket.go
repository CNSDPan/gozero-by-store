package types

// SocketMsg 广播消息结构体
type SocketMsg struct {
	Operate       int           `json:"operate"`
	Method        string        `json:"method"`
	StoreId       int64         `json:"storeId,string"`
	SendUserId    int64         `json:"sendUserId,string"`
	ReceiveUserId int64         `json:"ReceiveUserId,string"`
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

// DataByEnter 进入房间响应Data结构
type DataByEnter struct {
	RoomId   int64  `json:"roomId,string"`   // 房间
	ClientId int64  `json:"clientId,string"` // clientId
	UserId   int64  `json:"userId,string"`   // 用户id
	UserName string `json:"userName"`        // 发送人
}

// DataByNormal 进入普通消息响应Data结构
type DataByNormal struct {
	RoomId        int64  `json:"roomId,string"`     // 房间
	FromUserId    int64  `json:"fromUserId,string"` // 消息来源用户
	FrommUserName string `json:"fromUserName"`      // 消息来源用户
	Message       string `json:"message"`           // 消息内容
}

/******************Event 请求&响应结构*********************/
