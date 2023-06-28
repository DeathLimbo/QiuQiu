package main

import (
	"qiuqiu/common/module"
	"qiuqiu/common/service"
	"qiuqiu/common/sunnet"
	"qiuqiu/server/pvpserver/module/gate"
)

var Sun *sunnet.Sunnet

func main() {
	Sun = sunnet.NewSunnet()
	modules := []module.Module{gate.Module}
	Sun.Start(modules...)
	server := service.NewService()
	Sun.AddServer(server)
	Sun.Wait()
}
