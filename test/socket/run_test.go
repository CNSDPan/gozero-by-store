package socket

import (
	"testing"
)

func TestUser1(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAwODM5NzcsImlhdCI6MTcyOTQ3OTE3NywidXNlcklkIjoxODM3MDUzNzIwMjI4NjE4MjQwfQ.KZi1sSF0QtzJcsaEXys-LTsIBTvuo_Zcj6nchLLtv-A"
	s := NewServer(token)
	s.RunSocket(1837056807609659392, 1837053720228618240)
}
