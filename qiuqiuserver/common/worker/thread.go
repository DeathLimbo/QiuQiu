package worker

type Thread interface {
	GetId()
	Join()
}

type BaseThread struct {
	Thread
	id     uint32
	worker Worker
}

func NewThread(id uint32, worker Worker) Thread {
	return &BaseThread{
		id:     id,
		worker: worker,
	}
}

func (s *BaseThread) Join() {
	s.worker.Operator()
}
