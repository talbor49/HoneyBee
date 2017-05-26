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

