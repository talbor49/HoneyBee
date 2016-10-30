package server

import (
	"fmt"
	"grammar"
)

const (
	SUCCESS = "OK"
	ERROR   = "ERR"
)

func HandleQuery(query string, conn *DatabaseConnection) (returnCode string) {
	// TODO: write query in plain text to log
	requestType, tokens, err := grammar.ParseQuery(query)

	if err != nil {
		return err.Error()
	}

	switch requestType {
	case "AUTH":
		// AUTH {username} {password} {database_name}
		fmt.Println("Client wants to authenticate.")
		username := tokens[0]
		password := tokens[1]
		dbname := tokens[2]
		if credentialsValid(username, password) {
			conn.username = username
			conn.dbname = dbname
		}
		fmt.Println("User logged in as: ", username, password, " to database: "+dbname)
		return SUCCESS
	case "SET":
		// SET {key} {value} [ttl] [nooverride]
		fmt.Println("Client wants to set key")
		key := tokens[0]
		value := tokens[1]
		fmt.Println("Setting " + key + ":" + value)
		setRequest := SetRequest{Key: key, Value: value, Conn: conn}
		return handleSetRequest(setRequest)
	case "GET":
		// GET {key}
		fmt.Println("Client wants to get key")
		key := tokens[0]
		fmt.Println("Returning value of key: " + key)
		getRequest := GetRequest{Key: key, Conn: conn}
		return handleGetRequest(getRequest)
	case "DELETE":
		// DELETE {key}
		fmt.Println("Client wants to set key")
		return SUCCESS
	default:
		return ERROR
	}

}
