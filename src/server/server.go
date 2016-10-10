package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	PORT = "4590"
	IP   = "0.0.0.0"
)

func Startserver() {
	l, err := net.Listen("tcp", IP+":"+PORT)
	if err != nil {
		fmt.Println("Error listening on port "+PORT, err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on: " + IP + ":" + PORT)
	// Close the listener socket when the application closes.
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)
	reqLen, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading buffer: ", err.Error())
		return
	}

	authQuery := string(buff[:reqLen])
	username := HandleAuthentication(authQuery)

	fmt.Println("username: " + username)

	if username == "" {
		fmt.Println("Invalid credentials.")
		conn.Write([]byte("Invalid credentials"))
	} else {
		fmt.Println("Valid credentials")
		conn.Write([]byte("Succesfully connected"))
		data := "undefined"
		for strings.TrimSpace(data) != "" {
			reqLen, err = conn.Read(buff)
			if err != nil {
				fmt.Println("Error reading buffer: ", err.Error())
			}
			data = string(buff[:reqLen-1])
			HandleQuery(data)
		}
	}
}
