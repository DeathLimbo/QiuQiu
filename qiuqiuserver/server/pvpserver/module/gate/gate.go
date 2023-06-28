package gate

import "qiuqiu/common/network"

var (
	Module = new(module)
)

type module struct {
	*network.Gate
}

func (m *module) OnInit() {
	//netGate
	m.Gate = &network.Gate{
		Name: "gate",
		//NewAgent: game.HallAgentMgr.NewServerAgent,
		//FuncMaxConnNum: func() int {
		//	return conf.Server.MaxConnNum
		//},
		//PendingWriteNum: conf.PendingWriteNum,
		//TCPAddr:         conf.Core.TCPAddr,
	}
}
