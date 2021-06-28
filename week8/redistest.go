package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Dialredis() (redis.Conn, error) {
	rc, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, fmt.Errorf("redis链接失败，%w", err)
	}
	return rc, nil
}

func file(num int) string {
	var m string
	for i := 0; i < num; i++ {
		m = m + "a"
	}
	return m
}

func main() {
	rc, err := Dialredis()
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	m := file(500000)

	// set
	_, err = rc.Do("set", "abc8", m)
	if err != nil {
		fmt.Println(err.Error())
	}

	// // get
	// rec1, _ := rc.Do("Get", "abc")
	// fmt.Println(string(rec1.([]byte)))

	// // delete
	// _, err = rc.Do("DEL", "abc")
	// if err != nil {
	// 	fmt.Println("redis delelte failed:", err)

	// }

}
