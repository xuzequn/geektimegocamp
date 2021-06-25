package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

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
