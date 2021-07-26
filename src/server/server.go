package server

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/IsmaeelAkram/ltp/src/request"
)

func handleConn(conn net.Conn, id int) {
	for {
		bMethod := make([]byte, 12) // 12 bytes is max request size
		_, err := conn.Read(bMethod)
		if err != nil {
			break
		}
		bMethod = bytes.Trim(bMethod, "\x00") // Remove null bytes

		sMethod := strings.ReplaceAll(string(bMethod), "\n", "")
		method := request.GetMethod(sMethod)
		conn.Write([]byte(fmt.Sprintf("%d\n", method)))
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
