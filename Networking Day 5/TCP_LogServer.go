package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client.")
				return
			}
			log.Println("Error reading from connection:", err)
			return
		}

		remoteAddr := conn.RemoteAddr().String()
		log.Println("Received " + strconv.Itoa(n) + " bytes from " + remoteAddr)
		log.Println(string(buf[:n]))

		// Write the request to a file
		now := time.Now()
		logfile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening log file:", err)
			return
		}
		defer logfile.Close()
		var headers string
		log.SetOutput(logfile)
		log.Println("Time:", now)
		log.Println("Remote IP:", remoteAddr)
		log.Println("Payload size:", n)
		log.Println("Headers:", headers)
		log.Println("Payload:", string(buf[:n]))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
	}
	defer listener.Close()

	log.Println("TCP server running on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
