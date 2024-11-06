package socketlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/pkg/biz"
	"store/pkg/types"
	"store/pkg/xcode"

	"store/app/chat/rpc/internal/svc"
	"store/app/chat/rpc/pb/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type BroadcastMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastMsgLogic {
	return &BroadcastMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BroadcastMsg
// @Desc：广播消息到 消费者队列上
// @param：in
// @return：res
// @return：err
func (l *BroadcastMsgLogic) BroadcastMsg(in *chat.BroadcastReq) (res *chat.Response, err error) {
	var (
		e             error
		b             []byte
		code          = xcode.RESPONSE_SUCCESS
		broadcastBody = types.SocketMsgBody{
			Operate:      0,
			Method:       "",
			ResponseTime: "",
			Timestamp:    0,
			Event: types.Event{
				Params: "",
				Data:   types.DataByNormal{},
			},
		}
		dataByNormal = types.DataByNormal{
			StoreId:       in.StoreId,
			SendUserId:    in.SendUserId,
			SendUserName:  in.SendUserName,
			ReceiveUserId: in.ReceiveUserId,
			Message:       "",
		}
	)
	res = &chat.Response{}
	defer func() {
		res.Code, res.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 广播消息业务 operate:%d method:%d fail:%s", l.svcCtx.Config.ServiceName, in.Operate, in.Method, e.Error())
			res.ErrMsg = e.Error()
		}
	}()

	e = jsonx.Unmarshal([]byte(in.Body), &broadcastBody)
	if e != nil {
		code = xcode.SOCKET_BROADCAST_MSG_FAIL
		return res, nil
	}
	message, ok := broadcastBody.Event.Params.(string)
	if !ok {
		code = xcode.SOCKET_BROADCAST_MSG_STRING
		return res, nil
	}
	dataByNormal.Message = message
	broadcastBody.Operate = int(in.Operate)
	broadcastBody.Method = in.Method
	broadcastBody.Event.Data = dataByNormal
	b, e = jsonx.Marshal(broadcastBody)
	if e != nil {
		code = xcode.SOCKET_BROADCAST_MSG_FAIL
		return res, nil
	}
	// 发布消息，将消息都分发给订阅了的消费者,群聊|私聊都是同一个发布者，这里不做区分
	e = l.svcCtx.BizConn.Publish(l.ctx, biz.SOCKET_PUB_SUB_BROADCAST_NORMAL_CHAN_KEY, string(b)).Err()
	if e != nil {
		code = xcode.SOCKET_BROADCAST_MSG_PUB
		return res, nil
	}
	return res, nil
}
