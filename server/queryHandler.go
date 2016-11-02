package server

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/talbor49/HoneyBee/grammar"
)

const (
	SUCCESS               = "OK"
	ERROR                 = "ERR"
	ERR_UNAUTHORIZED_USER = "Unauthorized user. please authorize using the AUTH command."
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
		bucketname := tokens[2]
		if credentialsValid(username, password) {
			conn.Username = username

			bucketPath, _ := filepath.Abs(path.Join("data", bucketname+".hb"))

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
		}
		fmt.Println("User logged in as: ", username, password, " to database: "+bucketname)
		return SUCCESS
	case "SET":
		// SET {key} {value} [ttl] [nooverride]
		fmt.Println("Client wants to set key")
		if conn.Bucket != "" {
			key := tokens[0]
			value := tokens[1]
			fmt.Println("Setting " + key + ":" + value)
			setRequest := SetRequest{Key: key, Value: value, Conn: conn}
			return handleSetRequest(setRequest)
		} else {
			return ERR_UNAUTHORIZED_USER
		}

	case "GET":
		// GET {key}
		fmt.Println("Client wants to get key")
		if conn.Bucket != "" {
			key := tokens[0]
			fmt.Println("Returning value of key: " + key)
			getRequest := GetRequest{Key: key, Conn: conn}
			return handleGetRequest(getRequest)
		} else {
			return ERR_UNAUTHORIZED_USER
		}
	case "DELETE":
		// DELETE {key}
		fmt.Println("Client wants to delete a bucket/key")
		if conn.Bucket != "" {
			return SUCCESS
		} else {
			return ERR_UNAUTHORIZED_USER
		}
	default:
		return ERROR
	}

}
