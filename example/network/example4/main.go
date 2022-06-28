/*
等价的客户端和服务器
下面这个是两个服务器通信的例子，互为客户端和服务器，在发送数据报的时候，我们可以将发送的一方称之为源地址，发送的目的地一方称之为目标地址。
原文链接 https://colobu.com/2016/10/19/Go-UDP-Programming/
*/

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func read(conn *net.UDPConn) {
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}

func main() {
	addr1 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981}
	addr2 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9982}
	go func() {
		listener1, err := net.ListenUDP("udp", addr1)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listener1)
		time.Sleep(5 * time.Second)
		listener1.WriteToUDP([]byte("ping to #2: "+addr2.String()), addr2)
	}()
	go func() {
		listener1, err := net.ListenUDP("udp", addr2)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listener1)
		time.Sleep(5 * time.Second)
		listener1.WriteToUDP([]byte("ping to #1: "+addr1.String()), addr1)
	}()
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
