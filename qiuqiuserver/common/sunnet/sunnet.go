package sunnet

import (
	"qiuqiu/common/service"
	"qiuqiu/common/worker"
	"sync"
)

const (
	WORK_NUM = 100
)

type Sunnet struct {
	workerNum     uint32
	workers       []worker.Worker
	workerThreads []worker.Thread
	//服务列表
	services sync.Map // serverid : service
	maxId    uint32

	//socket县城
	SocketWorker *worker.WorkerSocket
	socketThread worker.Thread
}

func (s *Sunnet) StartWorker() {
	for i := uint32(0); i < s.workerNum; i++ {
		w := worker.NewWorker(i, 2<<i)
		t := worker.NewThread(i, w)
		s.workerThreads = append(s.workerThreads, t)
		s.workers = append(s.workers, w)
	}
}

func (s *Sunnet) Start() {
	s.StartWorker()
}

func (s *Sunnet) Wait() {
	if len(s.workerThreads) > 0 {
		s.workerThreads[0].Join()
	}
}

func NewSunnet() *Sunnet {
	s := &Sunnet{
		workerNum:     WORK_NUM,
		workers:       make([]worker.Worker, 0, WORK_NUM),
		workerThreads: make([]worker.Thread, 0, WORK_NUM),
		services:      sync.Map{},
	}
	return s
}

func (s *Sunnet) AddServer(server service.Service) {
	s.services.Store(server.GetId(), server)
}

func (s *Sunnet) KillServer(serverId uint32) {
	server := s.LoadServer(serverId)
	if server == nil {
		return
	}
	server.OnExit()
	s.services.Delete(serverId)
}

func (s *Sunnet) LoadServer(serverId uint32) service.Service {
	data, ok := s.services.Load(serverId)
	if !ok {
		return nil
	}
	server, ok := data.(service.Service)
	if !ok {
		return nil
	}
	return server
}

func (s *Sunnet) StartSocket() {
	socket := worker.NewWorkerSocket()
	socket.OnInit()
	thread := worker.NewThread(i, w)
}
