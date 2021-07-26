package server

import (
	"fmt"
	"net"

	"github.com/IsmaeelAkram/ltp/src/request"
)

func handleConn(conn net.Conn, id int) {
	for {
		req, err := request.GetRequest(conn, id)
		if err != nil {
			break
		}
		conn.Write([]byte(fmt.Sprintf("%d\n", req.Method)))
	}
	defer fmt.Printf("Connection #%d closed\n", id)
	conn.Close()
}

func StartListening() {
	ls, err := net.Listen("tcp", ":2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started")

	i := 0
	for {
		conn, err := ls.Accept()
		fmt.Printf("Connection #%d established from %s\n", i, conn.RemoteAddr().String())
		if err != nil {
			panic(err)
		}
		go handleConn(conn, i)
		i++
	}
}
