package main

import (
	"fmt"
	"net"
)

func main() {
	port := ":8082"
	// Resolve server address
	addr, _ := net.ResolveUDPAddr("udp", port)

	// Dial to server
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	// Send data
	conn.Write([]byte("Hello, server!"))

	// Read response
	buffer := make([]byte, 4096)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
