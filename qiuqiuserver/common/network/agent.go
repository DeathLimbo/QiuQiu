package network

type Agent interface {
	OnClose(code uint) //Agent关闭通知
	Addr() string
}
