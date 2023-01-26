package main

import (
	"fmt"
	"net"
)

func main() {
	port := ":8082"
	addr, _ := net.ResolveUDPAddr("udp", port)
	conn, _ := net.ListenUDP("udp", addr)

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, _ := conn.ReadFromUDP(buffer)
		fmt.Println("Received: ", string(buffer[:n]), " from ", addr)
	}
}
