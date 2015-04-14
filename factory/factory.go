package factory

import (
	"gotwisted/session"
	"net"
)

type IFactory interface {
	BuildSession(conn net.Conn) session.ISession
}

type FactoryBase struct {
	IFactory
}
