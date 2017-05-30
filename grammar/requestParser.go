package grammar

import (
	"errors"
	"log"
	"bytes"
)

func byte2dSliceToStringSlice(byteslice [][]byte) []string {
	var stringslice []string
	for _, bslice := range byteslice {
		if len(bslice) > 0 {
			stringslice = append(stringslice, string(bslice))
		}
	}
	return stringslice
}

func ParseRequest(data []byte) (Request, error) {
	if len(data) < 2 {
		return Request{}, errors.New("Request is too short.")
	}
	requestType := data[0]
	requestStatus := data[1]
	var tokens []string
	if len(data) > 2 {
		rawData := data[2:]
		tokens = byte2dSliceToStringSlice(bytes.Split(rawData, []byte{0}))
	}

	request := Request{Type: requestType, Status: requestStatus, RequestData: tokens}

	if !validRequestType(requestType) {
		return request, errors.New("Invalid request type")
	}
	if !validRequestStatus(requestStatus) {
		return request, errors.New("Invalid request status")
	}

	switch requestType {
	case GET_REQUEST:
		// form: "GET key"
		validLength := len(tokens) >= 1
		if validLength {
			return request, nil
		} else {
			return request, errors.New("GET request requires a key parameter.")
		}
	case SET_REQUEST:
		// form: "SET key value"
		validLength := len(tokens) >= 2
		if validLength {
			return request, nil
		} else {
			return request, errors.New("SET request requires parameters: key and value")
		}
	case AUTH_REQUEST:
		// form: "AUTH username password bucket"
		validLength := len(tokens) == 2
		if validLength {
			return request, nil
		} else {
			log.Println("Tokens recieved: ")
			log.Println(tokens)
			return request, errors.New("AUTH request should look like 'AUTH username password'")
		}
	case DELETE_REQUEST:
		// form: "DELETE [BUCKET/KEY] key"
		validLength := len(tokens) >= 2
		if validLength {
			return request, nil
		} else {
			return request, errors.New("A valid DELETE request looks like: 'DELETE [BUCKET/KEY] <name>'")
		}
	case USE_REQUEST:
		validLength := len(tokens) == 1
		if validLength {
			return request, nil
		} else {
			return request, errors.New("A valid USE request looks like: 'USE <bucket>'")
		}
	case CREATE_REQUEST:
		validLength := len(tokens) == 1
		if validLength {
			return request, nil
		} else {
			return request, errors.New("A valid CREATE request looks like: 'CREATE <bucket>'")
		}
	default:
		return request, errors.New("Unimplemented request type")
	}

}
