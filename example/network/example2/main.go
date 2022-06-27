/*
深入Go UDP编程
一个简单的UDP的例子，这个例子演示了Go UDP通过Dial方式发送数据报的例子
原文链接 https://colobu.com/2016/10/19/Go-UDP-Programming/
*/

package main

import (
	"fmt"
	"net"
)

// linux 发送UPD数据包测试
//echo "udp_testaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" | socat 4-datagram:192.168.8.68:9981

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}
