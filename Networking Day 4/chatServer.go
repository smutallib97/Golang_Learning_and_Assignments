package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clients   = make(map[net.Conn]string)
	history   = make(map[string][]string)
	broadcast = make(chan string)
	mu        sync.Mutex
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Get client name
	name := strings.TrimSpace(readMessage(conn))
	clients[conn] = name
	history[name] = []string{}

	broadcast <- fmt.Sprintf("%s has joined the chat", name)

	for {
		message := readMessage(conn)
		if message == "" {
			delete(clients, conn)
			broadcast <- fmt.Sprintf("%s has left the chat", name)
			break
		}

		// Store message in history
		mu.Lock()
		history[name] = append(history[name], message)
		mu.Unlock()

		broadcast <- fmt.Sprintf("%s: %s", name, message)
	}
}

func readMessage(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	return message
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	go func() {
		for {
			conn, _ := listener.Accept()
			go handleConnection(conn)
		}
	}()

	for {
		select {
		case message := <-broadcast:
			for conn := range clients {
				conn.Write([]byte(message))
			}
		}
	}
}
