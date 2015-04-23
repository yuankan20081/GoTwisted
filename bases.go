package gotwisted

import (
	"net"
)

type FactoryBase struct {
}

type SessionBase struct {
	Transport net.Conn
}
