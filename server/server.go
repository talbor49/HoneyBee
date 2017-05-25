package server

import (
	"fmt"
	"net"
	"os"
	"strings"
	"log"
)

const (
	port      = "8080"
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
	addr := fmt.Sprintf("%s:%s", ip, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Listening on: %s", addr)
	// Close the listener socket when the application closes.
	defer listener.Close()

	go QueueRequestsHandler()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting message from client, %s", err)
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
			log.Printf("Error reading buffer. %s", err)
			return
		}
		data = string(buff[:reqLen])
		for _, req := range strings.Split(data, "\n") {
			if len(req) == 0 {
				continue
			}
			returnMessage := HandleQuery(req, &conn)
			conn.Write([]byte(returnMessage + "\n"))
			log.Printf("Query handles with code %s", returnMessage)
		}

	}
	fmt.Println("Closed connection")
}
