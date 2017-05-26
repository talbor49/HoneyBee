package server

import (
	"fmt"
	"net"
	"os"
	"log"
	"github.com/talbor49/HoneyBee/grammar"
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

	var rawRequest []byte
	buff := make([]byte, bufferLen)

	for len(rawRequest) == 0 || rawRequest[0] != grammar.DELETE_REQUEST {
		reqLen, err := conn.Read(buff)
		if err != nil {
			log.Printf("Error reading buffer. %s", err)
			return
		}
		rawRequest = buff[:reqLen]
		returnMessage := HandleRequest(rawRequest, &conn)
		conn.Write([]byte(returnMessage + "\n"))
		log.Printf("Query handles with code %s", returnMessage)

	}
	log.Println("Closed connection")
}
