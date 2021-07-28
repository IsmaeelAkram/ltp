package request

import (
	"bytes"
	"net"
	"strings"
)

type Request struct {
	ID       int
	Method   MethodType
	Conn     net.Conn
	BRequest []byte
	SRequest string
}

var EMPTY_REQUEST = Request{}

func GetRequest(conn net.Conn, id int) (Request, error) {
	bRequest := make([]byte, 12) // 12 bytes is max request size
	_, err := conn.Read(bRequest)
	if err != nil {
		return Request{}, err
	}
	bRequest = bytes.Trim(bRequest, "\x00") // Remove null bytes
	sRequest := strings.ReplaceAll(string(bRequest), "\n", "")
	method := GetMethodType(strings.ToUpper(sRequest))

	return Request{id, method, conn, bRequest, sRequest}, nil
}
