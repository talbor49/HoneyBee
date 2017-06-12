package grammar

const (
	GET_REQUEST    = 1
	SET_REQUEST    = 2
	DELETE_REQUEST = 3
	AUTH_REQUEST   = 4
	CREATE_BUCKET_REQUEST = 5
	USE_REQUEST    = 6
	QUIT_REQUEST = 7
	CREATE_USER_REQUEST = 8
)

const (
	REQUEST_STATUS = 1
	KEY = 2
	BUCKET = 3
	USER = 4
)

const (
	LENGTH_OF_GET_REQUEST           = 1
	LENGTH_OF_SET_REQUEST           = 2
	LENGTH_OF_DELETE_REQUEST        = 1
	LENGTH_OF_AUTH_REQUEST          = 2
	LENGTH_OF_CREATE_BUCKET_REQUEST = 1
	LENGTH_OF_CREATE_USER_REQUEST 	= 2
	LENGTH_OF_USE_REQUEST           = 1
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
			reqType == CREATE_USER_REQUEST ||
			reqType == CREATE_BUCKET_REQUEST ||
			reqType == USE_REQUEST
	return validRequest
}

func validRequestStatus(reqStatus byte) bool {
	return true
}

func totalLength(st []string) int{
	length := 0
	for _, element := range st {
		length += len(element)
	}
	return length
}

func BuildRawRequest(request Request) (buffer []byte) {
	buffer = make([]byte, 2 + totalLength(request.RequestData))
	buffer[0] = request.Type
	buffer[1] = request.Status
	if len(request.RequestData) > 0 {
		for _, element := range request.RequestData {
			buffer = append(buffer, element...)
			buffer = append(buffer, []byte{0}...)
		}
	}
	buffer = append(buffer, []byte("\n")...)
	return buffer
}
