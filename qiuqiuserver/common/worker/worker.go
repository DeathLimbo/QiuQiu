package worker

import (
	"fmt"
	"time"
)

type Worker interface {
	GetId() int64
	GetType() string
	OnInit() //注意是否线程安全
	OnExit() //注意是否线程安全
	SendMsg(msg *Msg)
	ProcessMsg()
	IsFree() bool
	GetProMsg() chan Msg
}

type WorkerBase struct {
	id      int64     //服务唯一id
	typ     string    // 服务类型
	exiting bool      // 是否正在推出
	msg     chan Msg  //消息通道
	destory chan bool //销毁通道
}

func (s *WorkerBase) GetId() int64 {
	return s.id
}

func (s *WorkerBase) GetType() string {
	return s.typ
}

func (s *WorkerBase) GetProMsg() chan Msg {
	return s.msg
}

func (s *WorkerBase) OnInit() {
	fmt.Println(fmt.Sprintf("server:%v 初始化", s.GetId()))
}

func (s *WorkerBase) OnExit() {
	fmt.Println(fmt.Sprintf("server:%v 退出", s.GetId()))
}

func (s *WorkerBase) SendMsg(msg *Msg) {
	fmt.Println(fmt.Sprintf("server:%v 发送消息", s.GetId()))
}

func (s *WorkerBase) ProcessMsg() {
	for msg := range s.msg {
		fmt.Println(fmt.Sprintf("server:%v 处理消息 msg:%v", s.GetId(), msg))
	}
}

func (s *WorkerBase) IsFree() bool {
	if len(s.msg) < cap(s.msg) {
		return true
	}
	return false
}

func NewWorker() Worker {
	return &WorkerBase{
		id:      time.Now().Unix(),
		msg:     make(chan Msg, 1000),
		destory: make(chan bool, 1),
	}
}
