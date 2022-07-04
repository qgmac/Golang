package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":1080")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Listen failed: %v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("Connection from %s\n", remoteAddr)
	conn.Write([]byte("Hello world!\n"))
	conn.Close()
}
