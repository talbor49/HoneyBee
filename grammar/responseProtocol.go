package grammar

const (
	GET_INITIAL_RESPONSE = 128 + 1 // 		1000 0001
	SET_INITIAL_RESPONSE = 128 + 2 // 		1000 0002
	DELETE_INITIAL_RESPONSE = 128 + 3 // 		1000 0011
	AUTHORIZE_INITIAL_RESPONSE = 128 + 4 // 	1000 0100
	CREATE_INITIAL_RESPONSE = 128 + 5 // 		1000 0101
	USE_INITIAL_RESPONSE = 128 + 6 // 		1000 0110
	GET_FINAL_ANSWER = 128 + 64 + 1 // 		1100 0001
	SET_FINAL_ANSWER = 128 + 64 + 2 // 		1100 0002
	DELETE_FINAL_ANSWER = 128 + 64 + 3 // 		1100 0011
	AUTHORIZE_FINAL_ANSWER = 128 + 64 + 4 //	1100 0100
	CREATE_FINAL_ANSWER = 128 + 64 + 5 // 		1100 0101
	USE_FINAL_ANSWER = 128 + 64 + 6 // 		1100 0110
)


const (
	RESP_STATUS_SUCCESS = 0
	RESP_STATUS_ERR_UNAUTHORIZED = 1
	RESP_STATUS_ERR_NO_SUCH_BUCKET = 2
	RESP_STATUS_ERR_KEY_NOT_FOUND = 3
	RESP_STATUS_ERR_COULD_NOT_WRITE_TO_DISK = 4
	RESP_STATUS_ERR_BUCKET_ALREADY_EXISTS = 5
	RESP_STATUS_ERR_COULD_NOT_CREATE_BUCKET = 6
	RESP_STATUS_ERR_WRONG_CREDENTIALS = 7
)

type Response struct {
	Type byte
	Status byte
	Data string
}

func GetBufferFromResponse(response Response) []byte {
	var buffer []byte

	buffer = append(buffer, response.Type)
	buffer = append(buffer, response.Status)
	buffer = append(buffer, []byte(response.Data)...)

	return buffer
}