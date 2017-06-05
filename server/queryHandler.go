package server

import (
	"github.com/talbor49/HoneyBee/grammar"
	"log"
)

const (
	illegalRequestTemplate = "Illegal request by client, No such command '%d'"
)

func checkRequirements(request grammar.Request, conn *DatabaseConnection, requestParamsLength int, requiresLogin bool, requiresBucket bool) (err byte){
	if requiresBucket && conn.Bucket == "" {
		return grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
	}
	if requiresLogin && conn.Username == "" {
		return grammar.RESP_STATUS_ERR_UNAUTHORIZED
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
	request, err := grammar.ParseRequest(query)
	var response grammar.Response

	if err != nil {
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
			request.Status = errorStatus
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
			request.Status = errorStatus
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
			request.Status = errorStatus
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
			request.Status = errorStatus
			break
		}
		// TODO implement
	case grammar.CREATE_REQUEST:
		log.Println("Client wants to create a bucket")
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_CREATE_REQUEST,true, false)
		if errorStatus != 0 {
			request.Status = errorStatus
			break
		}

		bucketName := request.RequestData[0]
		createRequest := CreateRequest{BucketName: bucketName, Conn: conn}

		response = processCreateRequest(createRequest)
	case grammar.USE_REQUEST:
		errorStatus := checkRequirements(request, conn, grammar.LENGTH_OF_USE_REQUEST,true, false)
		if errorStatus != 0 {
			request.Status = errorStatus
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
	log.Printf("Writing buffer to client: %s", grammar.GetBufferFromResponse(response))
	conn.Write(grammar.GetBufferFromResponse(response))
	log.Printf("Wrote buffer to client")

}
