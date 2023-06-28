package network

import (
	"errors"
	"io"
	"net"
	"qiuqiu/common/msg"
	"sync"
	"sync/atomic"
	"time"
)

const (
	HeadLenSize    = 1
	SETLINGER_SECS = 5
)

type ConnSet map[net.Conn]struct{}

type TCPConn struct {
	conn      net.Conn
	writeChan chan msg.Msg
	Data      interface{} //自定义数据
	mutex     sync.Mutex
	isClosed  bool
	stat      *NetStat
}

func newTCPConn(conn net.Conn, pendingWriteNum int) *TCPConn {
	tcpConn := new(TCPConn)
	tcpConn.conn = conn
	tcpConn.writeChan = make(chan msg.Msg, pendingWriteNum)
	tcpConn.stat = &NetStat{}
	return tcpConn
}

func (tcpConn *TCPConn) Destroy() {
	tcpConn.conn.(*net.TCPConn).SetLinger(SETLINGER_SECS)
	time.Sleep(time.Second * 3)
	//writeStat := ""
	//if tcpConn.stat.WriteStat != nil {
	//	writeStat = tcpConn.stat.WriteStat.Print()
	//}
	tcpConn.conn.Close()
}

func (tcpConn *TCPConn) Close() {
	tcpConn.mutex.Lock()
	defer tcpConn.mutex.Unlock()

	if tcpConn.isClosed {
		return
	}
	close(tcpConn.writeChan)
	tcpConn.isClosed = true
}

// b must not be modified by the others goroutines
func (tcpConn *TCPConn) Write(b []byte, ops ...OptionFunc) error {
	if b == nil {
		return nil
	}

	tcpConn.mutex.Lock()
	defer tcpConn.mutex.Unlock()

	//if tcpConn.stat.WriteStat != nil {
	//	tcpConn.stat.WriteStat.pktAddTask()
	//}

	if tcpConn.isClosed {
		return errors.New("write channel is closed")
	}

	if len(tcpConn.writeChan) >= cap(tcpConn.writeChan) {
		return errors.New("write channel full")
	}
	pck := &msg.ServiceMSG{
		BaseMsg: msg.BaseMsg{
			Byte: b,
		},
	}
	op := &Option{}
	for _, f := range ops {
		f(op)
	}
	//if op.ExpireTime > 0 {
	//	pck.expirets = time.Now().Unix() + op.ExpireTime
	//}
	tcpConn.writeChan <- pck

	return nil
}

func (tcpConn *TCPConn) Read(b []byte) (int, error) {
	return tcpConn.conn.Read(b)
}

func (tcpConn *TCPConn) ReadFull(b []byte) error {
	_, err := io.ReadFull(tcpConn.conn, b)
	return err
}

func (tcpConn *TCPConn) LocalAddr() net.Addr {
	return tcpConn.conn.LocalAddr()
}

func (tcpConn *TCPConn) RemoteAddr() net.Addr {
	return tcpConn.conn.RemoteAddr()
}

func (tcpConn *TCPConn) WriteTask() {
	for pck := range tcpConn.writeChan {
		//if pck.expirets > 0 {
		//	nowTime := time.Now().Unix()
		//	if nowTime > pck.expirets {
		//		continue
		//	}
		//}
		_, err := tcpConn.conn.Write(pck.GetBody())
		if err != nil {
			continue
		}

		//if tcpConn.stat.WriteStat != nil {
		//	tcpConn.stat.WriteStat.pktDeal()
		//}
		atomic.AddUint32(&tcpConn.stat.SendPktNum, 1)
		//log.Release("tcp write to %s [%d-%d] size data", tcpConn.conn.RemoteAddr(), len(b), n)
	}

	// cleanup
	tcpConn.Destroy()
}

func (tcpConn *TCPConn) Stat() *NetStat {
	return tcpConn.stat
}

func (tcpConn *TCPConn) IsClosed() bool {
	tcpConn.mutex.Lock()
	defer tcpConn.mutex.Unlock()

	return tcpConn.isClosed
}
