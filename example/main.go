// example project main.go
package main

import (
	"fmt"
	"github.com/yuankan20081/gotwisted"
	"github.com/yuankan20081/gotwisted/session"
	"net"
)

type MySession struct {
	gotwisted.SessionBase // this is not necessary
}

func (ms *MySession) ParsePacket(conn net.Conn, buffer []byte) ([]byte, error) {
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:bytesRead], nil
}

func (ms *MySession) OnPacketParsed(packet []byte) {
	// echo
	ms.Transport.Write(packet)
}

func (ms *MySession) OnConnect(addr net.Addr) {
	fmt.Println(addr.String())
}

func (ms *MySession) OnDisconnect() {
	fmt.Println("ondisconnect")
}

type MyFactory struct {
	gotwisted.FactoryBase // this is not necessary
}

func (mf *MyFactory) BuildSession(conn net.Conn) session.ISession {
	s := new(MySession)
	s.Transport = conn
	return s
}

func main() {
	fmt.Println("Hello World!")
	reactor := gotwisted.NewReactor(&MyFactory{})
	reactor.StartListen("127.0.0.1", 8888)
}
