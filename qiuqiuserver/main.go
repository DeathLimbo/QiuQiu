package main

import (
	"qiuqiu/common/service"
	"qiuqiu/common/sunnet"
)

var Sun *sunnet.Sunnet

func main() {
	Sun = sunnet.NewSunnet()
	Sun.Start()
	server := service.NewService()
	Sun.AddServer(server)
	Sun.Wait()
}
