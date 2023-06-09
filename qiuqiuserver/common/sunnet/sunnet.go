package sunnet

import (
	"os"
	signal2 "os/signal"
	"qiuqiu/common/signal"
	"qiuqiu/common/worker"
)

type Sunnet struct {
	MsgChan chan worker.Msg
	Mgr     *worker.Manager
	Exit    chan bool
}

func NewSunnet() *Sunnet {
	mgr := worker.NewManager()
	s := &Sunnet{
		Mgr:     mgr,
		MsgChan: make(chan worker.Msg),
		Exit:    make(chan bool),
	}
	return s
}

func (s *Sunnet) Run() {
	c := make(chan os.Signal, 1)

	sigs := []os.Signal{os.Interrupt, os.Kill}
	if len(signal.Signals) > 0 {
		sigs = append(sigs, signal.Signals...)
	}

	signal2.Notify(c, sigs...)
	for {
		select {
		case sig := <-c:
			switch sig {

			}
		case <-s.Exit:
			break
		case msg := <-s.MsgChan:
			s.SendMsg2Worker(msg)
		}
	}

}

func (s *Sunnet) SendMsg2Worker(msg worker.Msg) {
	w := s.Mgr.FreeWorker()
	if w == nil {
		s.MsgChan <- msg
		return
	}
	w.GetProMsg() <- msg
}
