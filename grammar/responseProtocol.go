package grammar

import (
	"bytes"
	"fmt"
)

const (
	UNKNOWN_TYPE_RESPONSE = 0            //			0000 0000
	GET_RESPONSE          = 128 + 64 + 1 // 		1100 0001
	SET_RESPONSE          = 128 + 64 + 2 // 		1100 0002
	DELETE_RESPONSE       = 128 + 64 + 3 // 		1100 0011
	AUTHORIZE_RESPONSE    = 128 + 64 + 4 //			1100 0100
	CREATE_RESPONSE       = 128 + 64 + 5 // 		1100 0101
	USE_RESPONSE          = 128 + 64 + 6 // 		1100 0110
)

const (
	RESP_STATUS_SUCCESS                      = 0
	RESP_STATUS_ERR_UNAUTHORIZED             = 1
	RESP_STATUS_ERR_NO_SUCH_BUCKET           = 2
	RESP_STATUS_ERR_KEY_NOT_FOUND            = 3
	RESP_STATUS_ERR_COULD_NOT_WRITE_TO_DISK  = 4
	RESP_STATUS_ERR_BUCKET_ALREADY_EXISTS    = 5
	RESP_STATUS_ERR_COULD_NOT_CREATE_BUCKET  = 6
	RESP_STATUS_ERR_WRONG_CREDENTIALS        = 7
	RESP_STATUS_ERR_INVALID_QUERY            = 8
	RESP_STATUS_ERR_INVALID_AMOUNT_OF_PARAMS = 9
	RESP_STATUS_ERR_UNKNOWN_COMMAND          = 10
	RESP_STATUS_ERR_USERNAME_EXISTS 		 = 11
)

type Response struct {
	Type   byte
	Status byte
	Data   string
}

func GetBufferFromResponse(response Response) []byte {
	var buffer []byte

	buffer = append(buffer, response.Type)
	buffer = append(buffer, response.Status)
	buffer = append(buffer, []byte(response.Data)...)
	buffer = append(buffer, []byte("\n")...)

	return buffer
}

func GetResponseFromBuffer(buffer []byte) (response Response) {
	if len(buffer) < 1 {
		fmt.Errorf("Length of buffer is smaller than 1 byte :\\")
		return
	}
	response.Type = buffer[0]
	if len(buffer) > 1 {
		response.Status = buffer[1]
	} else {
		response.Status = RESP_STATUS_SUCCESS
	}
	if len(buffer) > 2 {
		response.Data = byte2dSliceToStringSlice(bytes.Split(buffer, []byte{0}))[0]
	} else {
		response.Data = ""
	}
	return
}

func ResponseToString(response Response) (str string) {
	switch response.Type {
		case UNKNOWN_TYPE_RESPONSE:
			str += "UNKNOWN_TYPE_RESPONSE "
		case GET_RESPONSE         :
			str += "GET_RESPONSE          "
		case SET_RESPONSE         :
			str += "SET_RESPONSE          "
		case DELETE_RESPONSE      :
			str += "DELETE_RESPONSE       "
		case AUTHORIZE_RESPONSE   :
			str += "AUTHORIZE_RESPONSE    "
		case CREATE_RESPONSE      :
			str += "CREATE_RESPONSE       "
		case USE_RESPONSE         :
			str += "USE_RESPONSE          "
	}
	switch response.Status {
		case RESP_STATUS_SUCCESS                      :
			str += "RESP_STATUS_SUCCESS                      "
		case RESP_STATUS_ERR_UNAUTHORIZED             :
			str += "RESP_STATUS_ERR_UNAUTHORIZED             "
		case RESP_STATUS_ERR_NO_SUCH_BUCKET           :
			str += "RESP_STATUS_ERR_NO_SUCH_BUCKET           "
		case RESP_STATUS_ERR_KEY_NOT_FOUND            :
			str += "RESP_STATUS_ERR_KEY_NOT_FOUND            "
		case RESP_STATUS_ERR_COULD_NOT_WRITE_TO_DISK  :
			str += "RESP_STATUS_ERR_COULD_NOT_WRITE_TO_DISK  "
		case RESP_STATUS_ERR_BUCKET_ALREADY_EXISTS    :
			str += "RESP_STATUS_ERR_BUCKET_ALREADY_EXISTS    "
		case RESP_STATUS_ERR_COULD_NOT_CREATE_BUCKET  :
			str += "RESP_STATUS_ERR_COULD_NOT_CREATE_BUCKET  "
		case RESP_STATUS_ERR_WRONG_CREDENTIALS        :
			str += "RESP_STATUS_ERR_WRONG_CREDENTIALS        "
		case RESP_STATUS_ERR_INVALID_QUERY            :
			str += "RESP_STATUS_ERR_INVALID_QUERY            "
		case RESP_STATUS_ERR_INVALID_AMOUNT_OF_PARAMS :
			str += "RESP_STATUS_ERR_INVALID_AMOUNT_OF_PARAMS "
		case RESP_STATUS_ERR_USERNAME_EXISTS:
			str += "RESP_STATUS_ERR_USERNAME_EXISTS  "
	}
	str += "data: " + response.Data
	return
}
