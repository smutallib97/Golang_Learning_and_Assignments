package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer listener.Close()
	fmt.Println("Listening on :8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		wg.Add(1)
		go handleConnection(conn, &wg)
	}
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	fmt.Println("Connection established:", conn.RemoteAddr())
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Received %d bytes: %s\n", n, string(buffer[:n]))

		_, err = conn.Write(buffer[:n])
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Sent %d bytes: %s\n", n, string(buffer[:n]))
	}
	fmt.Println("Connection closed:", conn.RemoteAddr())
}
