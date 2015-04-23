package factory

import (
	"github.com/yuankan20081/gotwisted/session"
	"net"
)

type IFactory interface {
	BuildSession(conn net.Conn) session.ISession
}
