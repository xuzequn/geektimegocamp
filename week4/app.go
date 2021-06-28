package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// func forceShutdownIfNeed() {
// 	time.AfterFunc(time.Minute, func() {
// 		os.Exit((1))
// 	})
// }

// func shutdown() {
// 	// 执行各种动作
// 	services.Range(func(key, value interface{}) bool {
// 		service := value.(Service)
// 		go func() {
// 			service.ShutDown()
// 		}()
// 		return true
// 	})
// }

func main() {

	AddService()

	http.HandleFunc("/", hander)

	go func() {
		listenSignal()
	}()

}

func listenSignal() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		<-signals
		os.Exit(1)
	}

}
