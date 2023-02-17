package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :8080...")

	for {
		// Wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Echo incoming data
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		data := buffer[:n]
		fmt.Println("Received:", string(data))

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
