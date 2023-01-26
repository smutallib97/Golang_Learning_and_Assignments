package main

import (
	"fmt"
	"net"
	"time"
)

const interval = 5 * time.Minute
const remoteAddr = "localhost:8080"

func main() {
	// Create a connection to the remote logging server
	conn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		fmt.Println("Error while connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send log information every 5 minutes
	for {
		logInfo := "This is some log information"
		_, err := conn.Write([]byte(logInfo))
		if err != nil {
			fmt.Println("Error while writing to server:", err)
			return
		}

		// Wait for response and print it to the console
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Error while reading from server:", err)
			return
		}
		fmt.Println("Server response:", string(response[:n]))

		time.Sleep(interval)
	}
}
