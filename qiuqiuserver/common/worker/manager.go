package worker

import (
	"sync"
)

var (
	WORK_NUM = 10
)

type Manager struct {
	services sync.Map // 服务id:worker
}

func (m *Manager) StoreWorker(worker Worker) {
	m.services.Store(worker.GetId(), worker)
}

func (m *Manager) DeleteWorker(id uint32) {
	worker := m.GetWorker(id)
	if worker == nil {
		return
	}
	worker.OnExit()
	m.services.Delete(id)
}

func (m *Manager) GetWorker(id uint32) Worker {
	data, ok := m.services.Load(id)
	if !ok {
		return nil
	}
	worker, ok := data.(Worker)
	if !ok {
		return nil
	}
	return worker
}

func NewManager() *Manager {
	mgr := &Manager{
		services: sync.Map{},
	}
	for i := 0; i < WORK_NUM; i++ {
		w := NewWorker()
		w.OnInit()
		mgr.StoreWorker(w)
	}
	return mgr
}

func (m *Manager) FreeWorker() Worker {
	var w Worker
	m.services.Range(func(key, value any) bool {
		data, ok := value.(Worker)
		if !ok {
			return true
		}
		if !data.IsFree() {
			return true
		}
		w = data
		return false
	})
	return w
}
