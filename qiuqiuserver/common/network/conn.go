package network

import "net"

type OptionFunc func(op *Option)
type Option struct {
	ExpireTime int64
}

type Conn interface {
	Read(data []byte) (int, error)
	ReadFull(data []byte) error
	Write(args []byte, ops ...OptionFunc) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()

	WriteTask()
	Stat() *NetStat
	IsClosed() bool
}

type NetStat struct {
	SendPktNum uint32
	RecvPktNum uint32
	RemoteAddr string
	//WriteStat  *NetPktStat
}
