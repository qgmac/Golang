package main

import (
	"fmt"
	"golang.org/x/net/ipv4"
	"net"
)

func main() {
	//1. 得到一个interface
	enp0s3, err := net.InterfaceByName("enp0s3")
	if err != nil {
		fmt.Println("卧槽这里报错了 InterfaceByName ->", err)
	}
	group := net.IPv4(224, 0, 0, 250)
	//2. bind一个本地地址
	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	if err != nil {
		fmt.Println("卧槽这里报错了 ListenPacket ->", err)
	}
	defer c.Close()
	//3.
	p := ipv4.NewPacketConn(c)
	if err := p.JoinGroup(enp0s3, &net.UDPAddr{IP: group}); err != nil {
		fmt.Println("卧槽这里报错了 NewPacketConn ->", err)
	}
	//4.更多的控制
	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
		fmt.Println(err)
	}
	//5.接收消息
	b := make([]byte, 1500)
	for {
		fmt.Println("---> 1")
		n, cm, src, err := p.ReadFrom(b)
		fmt.Println("---> 2")
		if err != nil {
			fmt.Println(err)
		}
		if cm.Dst.IsMulticast() {
			fmt.Println("Unknown group")
			if cm.Dst.Equal(group) {
				fmt.Printf("received: %s from <%s>\n", b[:n], src)
				n, err = p.WriteTo([]byte("world"), cm, src)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown group")
				continue
			}
		}
	}
}
