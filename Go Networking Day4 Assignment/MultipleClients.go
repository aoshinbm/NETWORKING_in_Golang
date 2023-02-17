package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//listen connection
	dstream, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstream.Close()

	for {
		//accept connection
		//conn handles connection n read some data from it
		conn, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnectionn(conn)
	}
}

// handle connection -->thread
// sort of thread n its runs separately from ur main thread
func handleConnectionn(conn net.Conn) {
	for {
		//data, err := ioutil.ReadAll(conn)
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
	conn.Close()
}
