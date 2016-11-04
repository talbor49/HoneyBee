package server

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/talbor49/HoneyBee/grammar"
)

const (
	success     = "OK"
	error       = "ERR"
	errNoBucket = "You are not connected to any bucket. use the 'SET {BUCKET}'"
)

// HandleQuery recieves a plain text string query, and hanles it.
// In most cases it adds it to the requests queue.
// Whilst in AUTH requests it validates the credentials and returns an answer.
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
		// bucketname := tokens[2]
		if credentialsValid(username, password) {
			conn.Username = username
		}
		fmt.Println("User logged in as: ", username)
		return success
	case "SET":
		// SET {key} {value} [ttl] [nooverride]
		fmt.Println("Client wants to set key")
		if conn.Bucket != "" {
			key := tokens[0]
			value := tokens[1]
			fmt.Println("Setting " + key + ":" + value)
			setRequest := SetRequest{Key: key, Value: value, Conn: conn}
			return handleSetRequest(setRequest)
		}
		return errNoBucket

	case "GET":
		// GET {key}
		fmt.Println("Client wants to get key")
		if conn.Bucket != "" {
			key := tokens[0]
			fmt.Println("Returning value of key: " + key)
			getRequest := GetRequest{Key: key, Conn: conn}
			return handleGetRequest(getRequest)
		}
		return errNoBucket

	case "DELETE":
		// DELETE {key}
		fmt.Println("Client wants to delete a bucket/key")
		if conn.Bucket != "" {
			return success
		}
		return errNoBucket
	case "CREATE":
		fmt.Println("Client wants to create a bucket")
		return success
	case "USE":
		fmt.Println("Client wants to use a specific bucket")
		bucketname := tokens[0]
		bucketPath, _ := filepath.Abs(path.Join("data", bucketname+".hb"))

		fmt.Println("Checking if there is a database at path: " + bucketPath)
		// If the bucket does not exist - create it.
		if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
			// Bucket does not exist
			f, err := os.Create(bucketPath)
			if err != nil {
				panic(err)
			}
			f.Close()
		}
		conn.Bucket = bucketname
		return success
	default:
		return error
	}

}
