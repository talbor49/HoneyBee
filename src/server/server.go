package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type DatabaseConnection struct {
	net.Conn
	dbname      string
	connections int
	username    string
}

const (
	PORT       = "8080"
	IP         = "0.0.0.0"
	BUFFER_LEN = 1024
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
			fmt.Println("Error accepting message. ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		dbconn := DatabaseConnection{conn, "", 0, ""}
		go handleConnection(dbconn)
	}
}

func handleConnection(conn DatabaseConnection) {
	// authenticate and process further requests
	defer conn.Close()

	data := ""
	buff := make([]byte, BUFFER_LEN)

	for strings.TrimSpace(data) != "quit" {
		reqLen, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading buffer: ", err.Error())
			return
		}
		data = string(buff[:reqLen])

		returnMessage := HandleQuery(data, &conn)
		conn.Write([]byte(returnMessage + "\n"))
	}
	fmt.Println("Closed connection")
}
