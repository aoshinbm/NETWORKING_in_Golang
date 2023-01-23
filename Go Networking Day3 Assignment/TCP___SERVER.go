package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Listen on a port
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Send data
	conn.Write([]byte("Hello, chillax! Lets rock my client......."))

	// Read response
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
	time.Sleep(30 * time.Second)
}
