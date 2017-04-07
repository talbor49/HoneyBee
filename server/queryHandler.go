package server

import (
	"fmt"

	"github.com/talbor49/HoneyBee/grammar"
)

const (
<<<<<<< HEAD
	success               = "OK"
	error                 = "ERR"
	errNoBucket           = "You are not connected to any bucket, use 'USE {BUCKET}'."
	errNotLoggedIn        = "You are not logged in, use 'Auth {username} {password}'."
	errBucketDoesNotExist = "Bucket does not exist, use 'CREATE {BUCKET}'"
=======
	success        = "OK"
	error          = "ERR"
	errNoBucket    = "You are not connected to any bucket. use the 'USE <BUCKET>'"
	errNotLoggedIn = "You are not logged in. Use 'AUTH <username> <password>'"
>>>>>>> 84b6353f02a9bd6662913c839e967d330cb40c0d
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
		if conn.Bucket == "" {
			return errNoBucket
		}
		if conn.Username == "" {
			return errNotLoggedIn
		}

		key := tokens[0]
		value := tokens[1]
		fmt.Println("Setting " + key + ":" + value)
		setRequest := SetRequest{Key: key, Value: value, Conn: conn}
		return handleSetRequest(setRequest)

	case "GET":
		// GET {key}
		fmt.Println("Client wants to get key")
		if conn.Bucket == "" {
			return errNoBucket
		}
		if conn.Username == "" {
			return errNotLoggedIn
		}
		key := tokens[0]
		fmt.Println("Returning value of key: " + key)
		getRequest := GetRequest{Key: key, Conn: conn}
		return handleGetRequest(getRequest)

	case "DELETE":
		// DELETE {key}
		fmt.Println("Client wants to delete a bucket/key")
		if conn.Bucket == "" {
			return errNoBucket
		}
		if conn.Username == "" {
			return errNotLoggedIn
		}
		return success
	case "CREATE":
		fmt.Println("Client wants to create a bucket")
		if conn.Bucket == "" {
			return errNoBucket
		}
		if conn.Username == "" {
			return errNotLoggedIn
		}

		return success
	case "USE":
<<<<<<< HEAD
=======
		fmt.Println("Client wants to use a specific bucket")
>>>>>>> 84b6353f02a9bd6662913c839e967d330cb40c0d
		if conn.Username == "" {
			return errNotLoggedIn
		}

		bucketname := tokens[0]

		useRequest := UseRequest{BucketName: bucketname, Conn: conn}

		return handleUseRequest(useRequest)
	default:
		return error
	}

}
