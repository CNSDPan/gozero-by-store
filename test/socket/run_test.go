package socket

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"log"
	"store/pkg/consts"
	"store/pkg/types"
	"testing"
	"time"
)

func TestUser1(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA2ODk1OTMsImlhdCI6MTczMDA4NDc5MywidXNlcklkIjoxODM3MDUzNzIwMjI4NjE4MjQwfQ.dLw_jwWyDgJF8sTzv2med_vVirWyFVIbZTJefjyBYRo"
	s := NewServer(token)
	s.RunSocket(1837056807609659392, 1837053720228618240)
}

func TestWrite(t *testing.T) {
	body := types.SocketMsgBody{
		Operate:      consts.OperatePublic,
		Method:       consts.MethodNormal,
		ResponseTime: time.Now().UTC().Format("2006-01-02 15:04:05"),
		Event: types.Event{
			Params: fmt.Sprintf("我来了"),
		},
	}
	b, err := jsonx.Marshal(body)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("b byte:%+v", string(b))

	broadcastBody := types.SocketMsgBody{
		Operate:      0,
		Method:       "",
		ResponseTime: "",
		Event: types.Event{
			Params: "",
			Data:   types.DataByNormal{},
		},
	}

	err = jsonx.Unmarshal(b, &broadcastBody)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("broadcastBody:%+v", broadcastBody)
}
