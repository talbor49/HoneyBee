package grammar

import (
	"bytes"
	"errors"
	"reflect"
)

func byte2dSliceToStringSlice(byteslice [][]byte) []string {
	var stringslice []string
	for _, bslice := range byteslice {
		if len(bslice) > 0 && !reflect.DeepEqual(bslice, []byte("\n")) && !reflect.DeepEqual(bslice, []byte{0}) {
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

	return request, nil
}
