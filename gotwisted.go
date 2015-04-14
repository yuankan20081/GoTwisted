// gotwisted project gotwisted.go
package gotwisted

import (
	"fmt"
	"github.com/yuankan20081/gotwisted/factory"
	"net"
)

type Reactor struct {
	sessionFactory factory.IFactory
}

func NewReactor(factoryImpl factory.IFactory) *Reactor {
	if factoryImpl == nil {
		panic("factoryImpl is nil!")
	}
	r := &Reactor{
		sessionFactory: factoryImpl,
	}

	return r
}

func (r *Reactor) StartListen(ip string, port uint32) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go r.handleConn(conn)
	}
}

func (r *Reactor) handleConn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[PANIC]", err)
		}
	}()
	defer conn.Close()

	clientSession := r.sessionFactory.BuildSession(conn)
	if clientSession.OnConnect != nil {
		clientSession.OnConnect(conn.RemoteAddr())
	}

	buf := make([]byte, 4096)

	for {
		packet, err := clientSession.ParsePacket(conn, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		clientSession.OnPacketParsed(packet)
	}

	if clientSession.OnDisconnect != nil {
		clientSession.OnDisconnect()
	}
}
