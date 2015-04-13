package session

import (
	"net"
)

type ISession interface {
	OnConnect(addr net.Addr)
	OnDisconnect()
	OnPacketParsed(data []byte)
	ParsePacket(conn net.Conn, buffer []byte) ([]byte, error)
}

type SessionBase struct {
	ISession
	Transport net.Conn
}
