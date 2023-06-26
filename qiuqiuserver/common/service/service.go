package service

import "qiuqiu/common/msg"

type Service interface {
	GetId() uint32
	OnInit()
	OnExit()
	ProcessMsg()
	ProcessMsgs(max int)
	OnMsg(msg2 msg.Msg) //收到消息时候触发
	PushMsg(msg2 msg.Msg)
}

type ServiceBase struct {
	id        uint32
	typ       string
	isExiting bool
	msgQueue  chan msg.Msg
}

func NewService() Service {
	return &ServiceBase{}
}

func (s *ServiceBase) OnInit() {

}

func (s *ServiceBase) GetId() uint32 {
	return s.id
}

func (s *ServiceBase) OnExit() {

}

func (s *ServiceBase) ProcessMsg() {

}
func (s *ServiceBase) ProcessMsgs(num int) {

}

func (s *ServiceBase) OnMsg(msg2 msg.Msg) {

}
func (s *ServiceBase) PushMsg(msg2 msg.Msg) {

}
