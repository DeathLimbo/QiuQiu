package network

type Gate struct {
	Name string
	// tcpserver
	TCPAddr            string
	NewAgent           func(*TCPConn) Agent
	tcpServer          *TCPServer
	FuncMaxConnNum     func() int
	PendingWriteNum    int
	NumberOfConnPreSec int32 // 每s限定的连接数

}

func (gate *Gate) Run(closeSig chan bool) {
	if gate.TCPAddr != "" {
		gate.tcpServer = new(TCPServer)
		gate.tcpServer.Addr = gate.TCPAddr
		gate.tcpServer.FuncMaxConnNum = gate.FuncMaxConnNum
		gate.tcpServer.PendingWriteNum = gate.PendingWriteNum
		gate.tcpServer.NewAgent = gate.NewAgent
		gate.tcpServer.NumberOfConnPreSec = gate.NumberOfConnPreSec
		gate.tcpServer.Start()
	}

	<-closeSig
}

func (gate *Gate) OnDestroy() {
	if gate.tcpServer != nil {
		gate.tcpServer.Close()
	}
}

func (gate *Gate) CloseTcpServer() {
	if gate.tcpServer != nil {
		gate.tcpServer.Close()
	}
}
