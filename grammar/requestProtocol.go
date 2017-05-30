package grammar

const (
	GET_REQUEST    = 1
	SET_REQUEST    = 2
	DELETE_REQUEST = 3
	AUTH_REQUEST   = 4
	CREATE_REQUEST = 5
	USE_REQUEST    = 6
	QUIT_REQUEST = 7
)

const (
	REQUEST_STATUS = 1
)

const (
	LENGTH_OF_GET_REQUEST    = 1
	LENGTH_OF_SET_REQUEST    = 2
	LENGTH_OF_DELETE_REQUEST = 1
	LENGTH_OF_AUTH_REQUEST   = 2
	LENGTH_OF_CREATE_REQUEST = 1
	LENGTH_OF_USE_REQUEST    = 1
)

type Request struct {
	Type byte
	Status byte
	RequestData []string
}

func validRequestType(reqType byte) bool {
	validRequest := reqType == GET_REQUEST ||
			reqType == SET_REQUEST ||
			reqType == DELETE_REQUEST ||
			reqType == AUTH_REQUEST ||
			reqType == CREATE_REQUEST ||
			reqType == USE_REQUEST
	return validRequest
}

func validRequestStatus(reqStatus byte) bool {
	return true
}


func BuildRawRequest(request Request) (buffer []byte) {
	buffer = append(buffer, request.Type)
	buffer = append(buffer, request.Status)
	if len(buffer) > 2 {
		for _, element := range request.RequestData {
			buffer = append(buffer, element...)
			buffer = append(buffer, 0)
		}
	}
}
