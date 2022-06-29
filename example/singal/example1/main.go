package main

import (
	"fmt"
	"os"
	"os/signal"
)

// 监听全部信号
func main() {
	c := make(chan os.Signal)
	// 监听所有信号
	signal.Notify(c)
	fmt.Println("启动了程序")
	s := <-c
	fmt.Println("收到信号:", s)
}
