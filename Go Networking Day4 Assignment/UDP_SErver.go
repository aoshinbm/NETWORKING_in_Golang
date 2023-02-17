/*
UDP is particularly quick since it doesn't offer error checking, correction, or packet retransmission.
UDP is typically preferred over TCP when speed is more important than reliability.
Online gaming, video chatting, and other real-time applications frequently employ UDP.
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on a port
	/*
		For UDP and IP networks, if the host in the address parameter is empty or a literal unspecified IP address,
		ListenPacket listens on all available IP addresses of the local system except multicast IP addresses.
	*/
	addr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		fmt.Println(err)
		return
	}

	// listen to incoming udp packets
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//create a loop for continuously receiving packages
	for {
		buffer := make([]byte, 4096)
		// Read data
		//With the help of the Read() function, we can get the data from the UDP client
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the data
		go handleData(conn, addr, buffer[:n])
	}
}

func handleData(conn *net.UDPConn, addr *net.UDPAddr,
	data []byte) {
	// Send a response
	conn.WriteToUDP([]byte("Hello, client! \n And Say..."), addr)

	// Print the received data
	fmt.Println(string(data))
}
