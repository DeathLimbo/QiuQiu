package network

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type TCPServer struct {
	Addr     string
	NewAgent func(*TCPConn) Agent

	ln net.Listener

	offChan            chan bool // 是否退出了
	NumberOfConn       int32     // 本次统计内的连接数
	NumberOfConnPreSec int32     // 1s允许的连接数

	PendingWriteNum int

	conns ConnSet

	wgLn sync.WaitGroup

	FuncMaxConnNum func() int
	mutexConns     sync.Mutex

	wgConns sync.WaitGroup
}

func (server *TCPServer) Start() {
	err := server.init()
	if err != nil {
		return
	}

	if server.NumberOfConnPreSec > 0 { // 做了连接限制
		go server.tick()
	}

	go server.run()
}

const RepeatCnt = 3

func (server *TCPServer) init() error {
	i := 0
	var ln net.Listener
	var err error
	for {
		ln, err = net.Listen("tcp4", server.Addr)
		if err != nil || ln == nil {
			if i < RepeatCnt {
				i++
				time.Sleep(time.Second)
				continue
			}
			return err
		} else {
			//成功
			break
		}
	}

	/*if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		log.Release("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}*/
	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = 100
	}
	if server.NewAgent == nil {
		return errors.New("NewAgent must not be nil")
	}

	server.ln = ln
	server.conns = make(ConnSet)
	return nil
}

/*
 * 开协程的目的主要是想借用系统的time来处理 而不愿采用加减时间来判定
 * 开协程只涉及到原子操作，对于开销上可不计
 */
func (server *TCPServer) tick() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()

	server.offChan = make(chan bool)
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-server.offChan:
			return
		case <-t.C:
			atomic.StoreInt32(&server.NumberOfConn, 0)
		}
	}
}

func (server *TCPServer) run() {
	server.wgLn.Add(1)
	defer func() {
		if server.NumberOfConnPreSec > 0 {
			server.offChan <- true
		}
		server.wgLn.Done()
	}()

	var tempDelay time.Duration
	for {
		conn, err := server.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0

		if server.NumberOfConnPreSec > 0 { // 开启了频率检测
			num := atomic.AddInt32(&server.NumberOfConn, 1)
			if num >= server.NumberOfConnPreSec { // 超过每s的连接数 直接断开
				conn.Close()
				continue
			}
		}

		maxConnNum := 0
		if server.FuncMaxConnNum != nil {
			maxConnNum = server.FuncMaxConnNum()
		}
		if maxConnNum == 0 {
			maxConnNum = 100
		}

		server.mutexConns.Lock()
		if len(server.conns) >= maxConnNum {
			server.mutexConns.Unlock()
			conn.Close()
			continue
		}
		server.conns[conn] = struct{}{}
		server.mutexConns.Unlock()

		server.wgConns.Add(1)

		tcpConn := newTCPConn(conn, server.PendingWriteNum)
		tcpConn.stat.RemoteAddr = conn.RemoteAddr().String()
		go server.newAgent(tcpConn)
	}
}

func (server *TCPServer) newAgent(tcpConn *TCPConn) {
	server.NewAgent(tcpConn)

	// cleanup
	server.mutexConns.Lock()
	delete(server.conns, tcpConn.conn)
	server.mutexConns.Unlock()
	server.wgConns.Done()
}

func (server *TCPServer) Close() {
	server.ln.Close()
	server.wgLn.Wait()

	server.mutexConns.Lock()
	for conn := range server.conns {
		conn.Close()
	}
	server.conns = nil
	server.mutexConns.Unlock()
	server.wgConns.Wait()
}
