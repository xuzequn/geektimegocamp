package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
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

		// 创建一个bufio.Reader 实体 接收数据
		reader := bufio.NewReader(conn)
		for {
			// 创建一个接收数据的切片数组
			// var data = make([]byte, 5)
			//  通过ReaderSlice方法 以“\n” 进行切分数据
			// data, err := reader.ReadSlice('\n')
			// 通过peek 去前几个字节的数据返回一个[]byte
			peek, err := reader.Peek(4)

			// n, err := conn.Read(data)
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				} else {
					break
				}
			}

			//  通过返回的[]byte 创建一个byte类型的buffer
			buffer := bytes.NewBuffer(peek)
			var length int32
			// 字节缓存 转化成 数字 ， binary.BigEndian 转化规则方法集。 三个参数分别是 转换前的字符数据， 转换方法集， 转换成的数据
			err = binary.Read(buffer, binary.BigEndian, &length)
			if err != nil {
				log.Println(err)
			}

			if int32(reader.Buffered()) < length+4 {
				continue
			}
			data := make([]byte, length+4)
			_, err = reader.Read(data)
			if err != nil {
				continue
			}

			log.Println("received msg", string(data[:4]))

			// log.Println("received msg", len(data), "bytes:", string(data))

			// if n > 0 {
			// log.Println("received msg", n, "bytes:", string(data[:n]))

			// }

		}
	}

}

func main() {
	start()

}
