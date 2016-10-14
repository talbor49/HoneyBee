package server

import (
	"fmt"
	"strings"
)

const (
	SUCCESS = "OK"
	ERROR   = "ERR"
)

func HandleQuery(query string) (returnCode string) {
	query = strings.TrimSpace(query)
	fmt.Println("query: " + query)

	tokens := strings.Split(query, " ")
	fmt.Println("tokens[0]: ", tokens[0])
	switch tokens[0] {
	case "AUTH":
		// AUTH {username} {password} {database_name}
		fmt.Println("Client wants to authenticate.")
		if len(tokens) == 4 {
			username := tokens[1]
			password := tokens[2]
			dbname := tokens[3]
			fmt.Println("User logged in as: ", username, password, " to database: "+dbname)
			return SUCCESS
		} else {
			fmt.Println("Request is of invalid form. ")
			return ERROR
		}
	case "SET":
		// SET {key} {value} [ttl] [nooverride]
		fmt.Println("Client wants to set key")
		if len(tokens) >= 3 {
			key := tokens[1]
			value := tokens[2]
			fmt.Println("Setting " + key + ":" + value)
			//write_to_hard_disk(key, value, database)
			return SUCCESS
		} else {
			fmt.Println("Request is of invalid form. ")
			return ERROR
		}
	case "GET":
		// GET {key}
		fmt.Println("Client wants to get key")
		if len(tokens) == 2 {
			key := tokens[1]
			fmt.Println("Returning value of key: " + key)
			return SUCCESS
		} else {
			fmt.Println("Request is of invalid form.")
			return ERROR
		}
	case "DELETE":
		// DELETE {key}
		fmt.Println("Client wants to set key")
		return SUCCESS
	default:
		return ERROR
	}

	// PUT RECORD {KEY} {VALUE} IN {BUCKET}

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

func write_to_hard_disk(key string, value string, database string) {

}

func HandleAuthentication(authQuery string) string {
	// Authentication is:            USERNAME PASSWORD DATABASE
	fmt.Println("authQuery: " + authQuery)
	usernameEndIndex := strings.Index(authQuery, " ")
	if usernameEndIndex != -1 {
		return authQuery[:strings.Index(authQuery, " ")]
	} else {
		return ""
	}
}
