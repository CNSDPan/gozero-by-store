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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ5NDEyMjcsImlhdCI6MTc0NDMzNjQyNywidXNlcklkIjoxODk4OTIwMTMxNjY1OTQ0NTc2fQ.yUmHcfbyjakX574IUnHp-UlT2q3flbC_uCMjvuRZhHM"
	s := NewServer(token)
	s.RunSocket(0, 1898920131665944576)
}

func TestUser2(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4ODUyMzcsImlhdCI6MTc0NDI4MDQzNywidXNlcklkIjoxOTA1MDgzMTQ3MTA5MzU5NjE2fQ.dneZ3d-7uqqQbZHcgrwEUs57xnEluApQUjVkIiVWQiM"
	s := NewServer(token)
	s.RunSocket(0, 1905083147109359616)
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
