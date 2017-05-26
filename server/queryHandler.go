package server

import (
	"github.com/talbor49/HoneyBee/grammar"
	"log"
)

const (
	success                = "OK"
	errNoSuchCommand                  = "No such command"
	errNoBucket            = "You are not connected to any bucket, use 'USE {BUCKET}'."
	errNotLoggedIn         = "You are not logged in, use 'Auth {username} {password}'."
	errBucketDoesNotExist  = "Bucket does not exist, use 'CREATE {BUCKET}'"
	errBucketAlreadyExists = "Can not create bucket, a bucket with that name already exists"
	illegalRequestTemplate = "Illegal request by client, %s"
)

// HandleQuery recieves a plain text string query, and hanles it.
// In most cases it adds it to the requests queue.
// Whilst in AUTH requests it validates the credentials and returns an answer.
func HandleQuery(query string, conn *DatabaseConnection) (returnCode string) {
	log.Printf("Handling query: %s", query)
	requestType, tokens, err := grammar.ParseQuery(query)

	if err != nil {
		return err.Error()
	}

	switch requestType {
	case "AUTH":
		// AUTH {username} {password} {database_name}
		log.Println("Client wants to authenticate.")
		username := tokens[0]
		password := tokens[1]
		// bucketname := tokens[2]

		authRequest := AuthRequest{Username: username, Password: password, Conn:conn}
		log.Printf("User logged in as: %s", username)
		return pushAuthRequestToQ(authRequest)
	case "SET":
		// SET {key} {value} [ttl] [nooverride]
		if conn.Bucket == "" {
			log.Printf(illegalRequestTemplate, errNoBucket)
			return errNoBucket
		}
		if conn.Username == "" {
			log.Printf(illegalRequestTemplate, errNoBucket)
			return errNotLoggedIn
		}

		key := tokens[0]
		value := tokens[1]
		log.Printf("Setting %s:%s", key, value)
		setRequest := SetRequest{Key: key, Value: value, Conn: conn}
		return pushSetRequestToQ(setRequest)

	case "GET":
		// GET {key}
		log.Println("Client wants to get key")
		if conn.Bucket == "" {
			log.Printf(illegalRequestTemplate, errNoBucket)
			return errNoBucket
		}
		if conn.Username == "" {
			log.Printf(illegalRequestTemplate, errNotLoggedIn)
			return errNotLoggedIn
		}
		key := tokens[0]
		log.Printf("Client asked for value of key: %s", key)
		getRequest := GetRequest{Key: key, Conn: conn}
		return pushGetRequestToQ(getRequest)

	case "DELETE":
		// DELETE {key}
		log.Println("Client wants to delete a bucket/key")
		if conn.Bucket == "" {
			log.Printf(illegalRequestTemplate, errNoBucket)
			return errNoBucket
		}
		if conn.Username == "" {
			log.Printf(illegalRequestTemplate, errNotLoggedIn)
			return errNotLoggedIn
		}
		// TODO implement
		return success
	case "CREATE":
		log.Println("Client wants to create a bucket")
		if conn.Username == "" {
			log.Printf(illegalRequestTemplate, errNotLoggedIn)
			return errNotLoggedIn
		}

		bucketName := tokens[0]

		createRequest := CreateRequest{BucketName: bucketName, Conn: conn}

		return pushCreateRequestToQ(createRequest)
	case "USE":
		if conn.Username == "" {
			log.Printf(illegalRequestTemplate, errNotLoggedIn)
			return errNotLoggedIn
		}

		bucketname := tokens[0]

		useRequest := UseRequest{BucketName: bucketname, Conn: conn}
		return pushUseRequestToQ(useRequest)
	default:
		log.Printf(illegalRequestTemplate, errNoSuchCommand)
		return errNoSuchCommand
	}

}
