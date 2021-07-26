package server

import "net"

func StartListening() {
	ls, err := net.Listen("tcp", ":2")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ls.Accept()
		if err != nil {
			panic(err)
		}
	}
}
