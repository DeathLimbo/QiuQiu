package main

import (
	"qiuqiu/common/connect"
	"time"
)

var Mgr *connect.Manager

func main() {
	Mgr = connect.NewManager()
	ser := func() {
		s := connect.NewService()
		s.OnInit()
		Mgr.StoreService(s)
	}
	for i := 0; i < 3; i++ {
		go func() {
			ser()
		}()
	}
	time.Sleep(100 * time.Second)
}
