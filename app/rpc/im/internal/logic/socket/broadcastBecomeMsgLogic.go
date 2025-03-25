package socketlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/rpc/im/internal/svc"
	"store/app/rpc/im/pb/im"
	"store/pkg/biz"
	"store/pkg/consts"
	"store/pkg/types"
	"store/pkg/xcode"
	"time"
)

type BroadcastBecomeMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastBecomeMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastBecomeMsgLogic {
	return &BroadcastBecomeMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BroadcastBecomeMsgLogic) BroadcastBecomeMsg(in *im.BroadcastReq) (res *im.Response, err error) {
	var (
		e             error
		b             []byte
		code          = xcode.RESPONSE_SUCCESS
		t             = time.Now()
		broadcastBody = types.SocketMsgBody{
			Operate:      consts.OperatePublic,
			Method:       consts.MethodBecome,
			ResponseTime: t.UTC().Format("2006-01-02 15:04:05"),
			Timestamp:    t.UnixMicro(),
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

	res = &im.Response{}
	defer func() {
		res.Code, res.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 广播消息业务 operate:%d method:%d fail:%s", l.svcCtx.Config.ServiceName, in.Operate, in.Method, e.Error())
			res.ErrMsg = e.Error()
		}
	}()
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
