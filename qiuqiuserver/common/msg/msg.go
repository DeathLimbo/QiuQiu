package msg

type MsgTyp int

const (
	SG_TYP_BASE   MsgTyp = 1
	SG_TYP_ACCEPT MsgTyp = 1 //新客户端链接
	SG_TYP_RW     MsgTyp = 1 // 客户端链接可读可写
)

type Msg interface {
	GetBody() []byte
}

type BaseMsg struct {
	Msg
	MsgTyp MsgTyp
}

type ServiceMSG struct {
	BaseMsg
	Source uint32
	Byte   []byte
	Lne    uint32
}

type SocketAcceptMsg struct {
	listenFd int
	clientFd int
	BaseMsg
}

type SocketRWMsg struct {
	fd      uint32
	isRead  bool
	isWrite bool
	BaseMsg
}
