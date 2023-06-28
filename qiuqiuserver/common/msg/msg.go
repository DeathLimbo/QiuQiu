package msg

type MsgTyp int

const (
	MSG_TYP_BASE   MsgTyp = 1
	MSG_TYP_ACCEPT MsgTyp = 2 //新客户端链接
	MSG_TYP_RW     MsgTyp = 3 // 客户端链接可读可写
)

type Msg interface {
	GetBody() []byte
}

type BaseMsg struct {
	Msg
	MsgTyp MsgTyp
	Byte   []byte
}

func (m *BaseMsg) GetBody() []byte {
	return m.Byte
}

type ServiceMSG struct {
	BaseMsg
	Source uint32

	Lne uint32
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
