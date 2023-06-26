package worker

import "fmt"

type WorkerSocket struct {
	WorkerBase
}

func (s *WorkerSocket) Operator() {
	s.isWorking = true
	for {
		select {
		case msg := <-s.msg:
			fmt.Println(fmt.Sprintf("server:%v 处理消息 msg:%v", s.GetId(), msg))
		default:
		}
	}
}

func NewWorkerSocket() *WorkerSocket {
	return &WorkerSocket{}
}
