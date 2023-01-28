package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Hello, server!... From Client 2"))

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
	time.Sleep(time.Second * 5)
}
