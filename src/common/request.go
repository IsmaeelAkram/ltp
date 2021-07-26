package common

import "net"

type Request struct {
	id     int
	method int
	conn   net.Conn
}
