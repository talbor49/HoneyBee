package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	port      = "4590"
	ip        = "0.0.0.0"
	bufferLen = 1024
)

//DatabaseConnection is an extension of the net.Conn struct, added additional required properties.
type DatabaseConnection struct {
	net.Conn
	Bucket      string
	Connections int
	Username    string
}

// StartServer starts the database server - listens at a specific port for any incoming TCP connections.
func StartServer() {
	l, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error listening on port "+port, err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on: " + ip + ":" + port)
	// Close the listener socket when the application closes.
	defer l.Close()

	go PriorityQueueWorker()

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
	buff := make([]byte, bufferLen)

	for strings.TrimSpace(data) != "quit" {
		reqLen, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading buffer: ", err.Error())
			return
		}
		data = string(buff[:reqLen])
		for _, req := range strings.Split(data, "\n") {
			fmt.Println("Request got: " + req)
			returnMessage := HandleQuery(req, &conn)
			fmt.Println("Query handles with code " + returnMessage)
		}

		// conn.Write([]byte(returnMessage + "\n"))
	}
	fmt.Println("Closed connection")
}
