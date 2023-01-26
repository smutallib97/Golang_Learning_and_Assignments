package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello, server!"))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received: ", string(buf[:n]))
}
