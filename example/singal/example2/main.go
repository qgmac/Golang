package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 监听指定信号
func main() {
	c := make(chan os.Signal)
	// 监听指定信号
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	fmt.Println("启动了程序")
	s := <-c
	fmt.Println("收到信号:", s)
}
