package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	sip := net.ParseIP("127.0.0.network")
	port := 9981
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: port}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	write, err := conn.Write([]byte("hello000"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(write)
	}

	fmt.Printf("<%s>\n", conn.RemoteAddr())
	for {
		time.Sleep(time.Second)
		_, err := conn.Write([]byte("hello000"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(write)
		}
	}
}
