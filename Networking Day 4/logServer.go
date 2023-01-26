package main

import (
	"fmt"
	"net"
)

const listenAddr = ":8080"

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on", listenAddr)

	// Handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// Read log information from the client
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		logInfo := string(buf[:n])

		// Print log information to the console
		fmt.Println("Received log:", logInfo)

		// Send response to the client
		response := "Log received"
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}
