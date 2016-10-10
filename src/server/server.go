package main

import (
	"fmt"
	"net"
	"os"
)

const (
	PORT = "4590"
)

func main() {
	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Println("Error listening on port "+PORT, err.Error())
		os.Exit(1)
	}
	// Close the listener socket when the application closes.
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buff := make([]byte, 1024)
	reqLen, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading buffer: ", err.Error())
	}
	data := string(buff[:reqLen])
	fmt.Println("data recieved: " + data)
	conn.Write([]byte("Message recieved."))
	conn.Close()
}
