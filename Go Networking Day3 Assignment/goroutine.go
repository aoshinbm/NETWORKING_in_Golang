package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
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

	wg.Wait()
}

func handleConnection(conn net.Connwg *sync.WaitGroup) {
	
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
	defer wg.Done()
	time.Sleep(20 * time.Second) // 5 seconds delay
}
