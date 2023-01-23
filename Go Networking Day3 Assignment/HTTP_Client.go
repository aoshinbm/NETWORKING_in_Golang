package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	request := "GET / HTTP/1.1\r\n" +
		"Host: localhost:8080\r\n" +
		"Connection: close\r\n" +
		"\r\n"

	conn.Write([]byte(request))

	response := make([]byte, 1024)
	length, _ := conn.Read(response)
	fmt.Println(string(response[:length]))
}
