package worker

import (
	"fmt"
	"qiuqiu/common/msg"
)

var (
	WORKER = 10000
)

type Worker interface {
	GetId() uint32
	GetType() string
	OnInit() //注意是否线程安全
	OnExit() //注意是否线程安全
	SendMsg(msg msg.Msg)
	Operator()
	IsFree() bool
	IsWorking() bool
}

type WorkerBase struct {
	id        uint32       //服务唯一id
	eachNum   uint32       //服务唯一id
	typ       string       // 服务类型
	exiting   bool         // 是否正在推出
	msg       chan msg.Msg //消息通道
	destory   chan bool    //销毁通道
	isWorking bool         // 是否开始工作
}

func (s *WorkerBase) GetId() uint32 {
	return s.id
}

func (s *WorkerBase) GetType() string {
	return s.typ
}

func (s *WorkerBase) OnInit() {
	fmt.Println(fmt.Sprintf("server:%v 初始化", s.GetId()))
}

func (s *WorkerBase) OnExit() {
	s.isWorking = false
	fmt.Println(fmt.Sprintf("server:%v 退出", s.GetId()))
}

func (s *WorkerBase) SendMsg(msg msg.Msg) {
	s.msg <- msg
	fmt.Println(fmt.Sprintf("server:%v 发送消息", s.GetId()))
}

func (s *WorkerBase) Operator() {
	s.isWorking = true
	for {
		select {
		case msg := <-s.msg:
			fmt.Println(fmt.Sprintf("server:%v 处理消息 msg:%v", s.GetId(), msg))
		default:
		}
	}
}

func (s *WorkerBase) IsFree() bool {
	if !s.isWorking {
		return false
	}
	if len(s.msg) < cap(s.msg) {
		return true
	}
	return false
}

func (s *WorkerBase) IsWorking() bool {

	return s.isWorking
}

func NewWorker(id, eachNum uint32) Worker {
	return &WorkerBase{
		id:      id,
		eachNum: eachNum,
		msg:     make(chan msg.Msg, WORKER),
		destory: make(chan bool, 1),
	}
}
