package session

import (
	"net"
)

type OnConnecter interface {
	OnConnect(addr net.Addr)
}

type OnDisconnecter interface {
	OnDisconnect()
}

type ISession interface {
	OnConnecter
	OnDisconnecter
	OnPacketParsed(data []byte)
	ParsePacket(conn net.Conn, buffer []byte) ([]byte, error)
}
