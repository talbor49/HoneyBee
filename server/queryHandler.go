package server

import (
	"github.com/talbor49/HoneyBee/grammar"
	"log"
)

const (
	illegalRequestTemplate = "Illegal request by client, No such command '%d'"
)

func checkRequirements(request grammar.Request, conn *DatabaseConnection, requestParamsLength int, requiresLogin bool, requiresBucket bool) (err byte){
	if requiresLogin && conn.Username == "" {
		return grammar.RESP_STATUS_ERR_UNAUTHORIZED
	}
	if requiresBucket && conn.Bucket == "" {
		return grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
	}
	if len(request.RequestData) != requestParamsLength {
		return grammar.RESP_STATUS_ERR_INVALID_AMOUNT_OF_PARAMS
	}
	return 0
}

// HandleRequest recieves a plain text string query, and hanles it.
// In most cases it adds it to the requests queue.
// Whilst in AUTH requests it validates the credentials and returns an answer.
func HandleRequest(query []byte, conn *DatabaseConnection) {
	log.Printf("Handling raw query: %s", query)
	log.Printf("Parsing request...")
	request, err := grammar.ParseRequest(query)
	log.Printf("Parsed request")
	var response grammar.Response

	if err != nil {
		log.Printf("Error in request parsing! %s", err.Error())
		response.Type = grammar.UNKNOWN_TYPE_RESPONSE
		response.Status = grammar.RESP_STATUS_ERR_INVALID_QUERY
		response.Data = err.Error()
		conn.Write(grammar.GetBufferFromResponse(response))
	}

	switch request.Type {
	case grammar.AUTH_REQUEST:
		// AUTH {username} {password}
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_AUTH_REQUEST,false, false)
		if errorStatus != 0 {
			log.Printf("Error in AUTH request! %d", errorStatus)
			response.Status = errorStatus
			break
		}
		username := request.RequestData[0]
		password := request.RequestData[1]
		// bucketname := tokens[2]
		log.Printf("Client wants to authenticate.<username>:<password> %s:%s", username, password)

		authRequest := AuthRequest{Username: username, Password: password, Conn:conn}
		response = processAuthRequest(authRequest)
	case grammar.SET_REQUEST:
		// SET {key} {value} [ttl] [nooverride]
		request.Type = grammar.SET_RESPONSE
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_SET_REQUEST,true, true)
		if errorStatus != 0 {
			log.Printf("Error in SET request! %d", errorStatus)
			response.Status = errorStatus
			break
		}

		key := request.RequestData[0]
		value := request.RequestData[1]
		log.Printf("Setting %s:%s", key, value)
		setRequest := SetRequest{Key: key, Value: value, Conn: conn}
		response = processSetRequest(setRequest)

	case grammar.GET_REQUEST:
		// GET {key}
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_GET_REQUEST,true, true)
		if errorStatus != 0 {
			log.Printf("Error in GET request! %d", errorStatus)
			response.Status = errorStatus
			break
		}

		key := request.RequestData[0]
		log.Printf("Client wants to get key '%s'", key)
		getRequest := GetRequest{Key: key, Conn: conn}
		response = processGetRequest(getRequest)

	case grammar.DELETE_REQUEST:
		// DELETE {key}
		log.Println("Client wants to delete a bucket/key")
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_DELETE_REQUEST,true, true)
		if errorStatus != 0 {
			log.Printf("Error in DELETE request! %d", errorStatus)
			response.Status = errorStatus
			break
		}
		// TODO implement
	case grammar.CREATE_BUCKET_REQUEST:
		log.Println("Client wants to create a bucket")
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_CREATE_BUCKET_REQUEST,true, false)
		if errorStatus != 0 {
			log.Printf("Error in CREATE bucket request! %d", errorStatus)
			response.Status = errorStatus
			break
		}

		bucketName := request.RequestData[0]
		createBucketRequest := CreateBucketRequest{BucketName: bucketName, Conn: conn}

		response = processCreateBucketRequest(createBucketRequest)
	case grammar.CREATE_USER_REQUEST:
		log.Printf("Client wants to create a user")
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_CREATE_USER_REQUEST,false, false)
		if errorStatus != 0 {
			log.Printf("Error in CREATE user request! %d", errorStatus)
			response.Status = errorStatus
			break
		}

		username := request.RequestData[0]
		password := request.RequestData[1]
		createUserRequest := CreateUserRequest{Username: username, Password:password, Conn: conn}

		response = processCreateUserRequest(createUserRequest)
	case grammar.USE_REQUEST:
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_USE_REQUEST,true, false)
		if errorStatus != 0 {
			log.Printf("Error in USE request! %d", errorStatus)
			response.Status = errorStatus
			break
		}

		bucketname := request.RequestData[0]

		useRequest := UseRequest{BucketName: bucketname, Conn: conn}
		response = processUseRequest(useRequest)
	default:
		log.Printf(illegalRequestTemplate, request.Type)
		response.Type = grammar.UNKNOWN_TYPE_RESPONSE
		response.Status = grammar.RESP_STATUS_ERR_UNKNOWN_COMMAND
	}
	if response.Status != 0 {
		log.Printf("Error in request. status: %d", response.Status)
	}
	conn.Write(grammar.GetBufferFromResponse(response))
	log.Printf("Wrote buffer: %s to client", grammar.GetBufferFromResponse(response))

}
