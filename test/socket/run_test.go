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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM0OTIxNDgsImlhdCI6MTc0Mjg4NzM0OCwidXNlcklkIjoxODk4OTIwMTMxNjY1OTQ0NTc2fQ.gCCdRQSqDY8O3bxjx6JlxJZy-rJ6NN4z_f3-I2LZMwc"
	s := NewServer(token)
	s.RunSocket(1837056807609659392, 1898920131665944576)
}

func TestUser2(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM0OTI5MTksImlhdCI6MTc0Mjg4ODExOSwidXNlcklkIjoxODk4OTE5MjA2Nzc5OTY5NTM2fQ.pntwHSj83lz5x4OxCirpm5pmNELygFaoiEb5lHR1EhE"
	s := NewServer(token)
	s.RunSocket(1837056807609659392, 1898919206779969536)
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
