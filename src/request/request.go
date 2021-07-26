package request

import "net"

type Request struct {
	id     Method
	method int
	conn   net.Conn
}
