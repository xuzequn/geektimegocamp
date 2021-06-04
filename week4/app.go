package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//  创建一个signal channel
	sigs := make(chan os.Signal, 1)
	// 创建一个bool channel
	done := make(chan bool, 1)

	// 注册要收到的信号， syscall.sigint:接收ctrl+c、 sigterm：程序退出
	// 信号没有信号参数表述接受所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 此goroutine为执行阻塞接收信号，一旦有了它，它就会打印出来
	// 然后通知程序可以完成
	go func() {
		sig := <-sigs
		fmt.Println(1, sig)
		done <- true
	}()

	// 不允许继续往sigs 中存入内容
	signal.Stop(sigs)
	//程序在此处等待，直到它预期信号
	// 在 “done” 上发送一个值，然后退出
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
