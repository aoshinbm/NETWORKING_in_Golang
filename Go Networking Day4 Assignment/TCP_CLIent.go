package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Received message: ", string(buf[:n]))
}
