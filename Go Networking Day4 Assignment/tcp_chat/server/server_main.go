package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// keep track of no of connections we have
// every we add a connection it is added to the map
var openConnections = make(map[net.Conn]bool)

// accepts or passes the connection type
var newConnection = make(chan net.Conn)

// a connection dies or closed
var deadConnection = make(chan net.Conn)

/*
var (

	openConnections = make(map[net.Conn]bool)
	newConnection   = make(chan net.Conn)
	deadConnection  = make(chan net.Conn)

)
*/
func main() {
	listn, err := net.Listen("tcp", ":8080")
	fmt.Println("Test server")
	logFatal(err)

	defer listn.Close() //close the listener after all the code is executed

	go func() {
		for {
			//invoke Accept() method on listener interface
			//n Accept will return connection interface
			//n this will happen to every new client who want to connect to server
			conn, err := listn.Accept()
			logFatal(err)

			openConnections[conn] = true

			//to use this connection outside the goroutine
			//then we use newConnection channel to pass it around
			//now it will be outside the goroutine as main func() n this goroutine will run concurrently
			newConnection <- conn
		}
	}()

	//cheecking a connection is der or not
	//pass a new connection to println
	//fmt.Println(<-newConnection)

	//to receive the client side message
	//pass newConnection which was inside goroutine to connection
	connection := <-newConnection

	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	logFatal(err)
	fmt.Println(message)

}
