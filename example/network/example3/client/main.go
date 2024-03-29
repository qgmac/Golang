/*
更复杂的例子
我们还可以将上面的例子演化一下，实现双向的读写。

服务器端代码不用修改，因为它已经实现了读写，读是通过listener.ReadFromUDP,写通过listener.WriteToUDP
原文链接 https://colobu.com/2016/10/19/Go-UDP-Programming/
*/

package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("127.0.0.network")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	_, err = conn.Write([]byte("hello xxxx"))
	if err != nil {
		return
	}
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())
}
