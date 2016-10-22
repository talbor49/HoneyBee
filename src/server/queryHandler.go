package server

import (
	"fmt"
	"grammar"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	SUCCESS = "OK"
	ERROR   = "ERR"
)

func HandleQuery(query string, conn *DatabaseConnection) (returnCode string) {
	// TODO: write query in plain text to log

	// queriesQueue priorityQue

	/*
		Parse key / value / stuff
	*/

	// in the background, clean "cold" (unused) records from RAM

	// RULE OF THUMB - UPDATE LOGS WHATEVER YOU DO

	// current decision - don't compress keys, only compress values

	// priorityQue.Push(query)
	// if request is WRITE {
	// 	inconsistentKeys.append(key)  // When someone tries to access one of the keys in this list, push it up the priority queue (at least above the GET request)
	// }

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
		write_to_hard_disk(key, value, conn.dbname)
		return SUCCESS
	case "GET":
		// GET {key}
		fmt.Println("Client wants to get key")
		key := tokens[0]
		fmt.Println("Returning value of key: " + key)
		return SUCCESS
	case "DELETE":
		// DELETE {key}
		fmt.Println("Client wants to set key")
		return SUCCESS
	default:
		return ERROR
	}

	/*
		if GET REQUEST {
			if IS IN RAM {
				return FROM RAM
			} ELSE IF IS IN HARD DISK {
				// calculate if record is hot enough to be put in RAM
				return FROM HARD DISK
			} else {
				return NOT FOUND
			}
		}
		else if SET REQUEST {
			FIRST:
				// DECIDE IF TO KEEP A POINTER TO THE VALUE IN MEMORY OR THE VALUE OF ITSELF
				PUT IN RAM
				REMOVE FROM INCONSISTENTKEYS
			THEN:
				// COMPRESS VALUE WHEN WRITING TO HARD DISK
				PUT IN HARD DISK
				UPDATE CACHED MEMORY
		}
	*/

}

func credentialsValid(username string, password string) bool {
	return true
}

func write_to_hard_disk(key string, value string, database string) {
	fmt.Println(database + "->" + key + ":" + value)

	dbPath, _ := filepath.Abs(path.Join("data", database+".hb"))

	fmt.Println("dbPath: " + dbPath)

	f, err := os.OpenFile(dbPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(key + ":" + value + "\n"); err != nil {
		panic(err)
	}
}

func HandleAuthentication(authQuery string) string {
	// Returns username if authentication is successful, else return empty string
	// Authentication is:            USERNAME PASSWORD DATABASE
	fmt.Println("authQuery: " + authQuery)
	usernameEndIndex := strings.Index(authQuery, " ")
	if usernameEndIndex != -1 {
		return authQuery[:strings.Index(authQuery, " ")]
	} else {
		return ""
	}
}
