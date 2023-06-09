package connect

import (
	"fmt"
	"time"
)

type Service interface {
	GetId() int64
	GetType() string
	OnInit() //注意是否线程安全
	OnExit() //注意是否线程安全
	SendMsg(msg *Msg)
	ProcessMsg()
}

type serviceBase struct {
	id      int64     //服务唯一id
	typ     string    // 服务类型
	exiting bool      // 是否正在推出
	msg     chan Msg  //消息通道
	destory chan bool //销毁通道
}

func (s *serviceBase) GetId() int64 {
	return s.id
}

func (s *serviceBase) GetType() string {
	return s.typ
}

func (s *serviceBase) OnInit() {
	fmt.Println(fmt.Sprintf("server:%v 初始化", s.GetId()))
}

func (s *serviceBase) OnExit() {
	fmt.Println(fmt.Sprintf("server:%v 退出", s.GetId()))
}

func (s *serviceBase) SendMsg(msg *Msg) {
	fmt.Println(fmt.Sprintf("server:%v 发送消息", s.GetId()))
}

func (s *serviceBase) ProcessMsg() {
	for {
		fmt.Println(fmt.Sprintf("server:%v 处理消息", s.GetId()))
		time.Sleep(2 * time.Second)
	}
}

func NewService() Service {
	return &serviceBase{
		id:      time.Now().Unix(),
		msg:     make(chan Msg, 1000),
		destory: make(chan bool, 1),
	}
}
