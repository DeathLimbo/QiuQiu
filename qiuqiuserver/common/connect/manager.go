package connect

import (
	"sync"
)

type Manager struct {
	services sync.Map     // 服务id:service
	Process  chan Service //处理服务通道
	ExitChan chan bool
}

func (m *Manager) Run() {
	for {
		select {
		case s := <-m.Process:
			m.StoreService(s)
		case <-m.ExitChan:
			return
		}
	}
}

func (m *Manager) StoreService(service Service) {
	m.services.Store(service.GetId(), service)
}

func (m *Manager) DeleteService(id uint32) {
	service := m.GetService(id)
	if service == nil {
		return
	}
	service.OnExit()
	m.services.Delete(id)
}

func (m *Manager) GetService(id uint32) Service {
	data, ok := m.services.Load(id)
	if !ok {
		return nil
	}
	service, ok := data.(Service)
	if !ok {
		return nil
	}
	return service
}

func NewManager() *Manager {
	return &Manager{
		services: sync.Map{},
		Process:  make(chan Service, 10),
	}
}
