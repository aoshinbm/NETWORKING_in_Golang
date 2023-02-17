package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve server address
	//use the ResolveUDPAddr() and DialUDP() to establish a connection to the UDP server
	//ResolveUDPAddr returns an address of UDP end point.
	//The network must be a UDP network name
	addr, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Dial to server
	//DialUDP acts like Dial for UDP networks.
	//Dial connects to the address on the named network.
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	//close the connection
	defer conn.Close()

	//Write() and Read() function to send and receive messages from the server
	// Send data
	conn.Write([]byte("Hello, server! A UDP Message"))

	// Read response
	buffer := make([]byte, 4096)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
