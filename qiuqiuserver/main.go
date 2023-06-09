package main

import (
	"qiuqiu/common/sunnet"
)

var Sun *sunnet.Sunnet

func main() {
	Sun = sunnet.NewSunnet()

	Sun.Run()
}
