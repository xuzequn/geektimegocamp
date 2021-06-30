package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func start() {
	listener, err := net.Listen("tcp", "127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		reader := bufio.NewReader(conn)
		for {
			// var data = make([]byte, 5)
			data, err := reader.ReadSlice('\n')

			// n, err := conn.Read(data)
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				} else {
					break
				}
			}
			log.Println("received msg", len(data), "bytes:", string(data))

			// if n > 0 {
			// log.Println("received msg", n, "bytes:", string(data[:n]))

			// }

		}
	}

}

func main() {
	start()

}
