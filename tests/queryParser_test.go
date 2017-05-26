package tests

import (
	"testing"

	"github.com/talbor49/HoneyBee/grammar"
)

func TestParseGetRequest(t *testing.T) {
	var rawGetRequest []byte
	rawGetRequest = append(rawGetRequest, grammar.GET_REQUEST)
	rawGetRequest = append(rawGetRequest, grammar.REQUEST_STATUS)
	rawGetRequest = append(rawGetRequest, []byte("foo")...)
	rawGetRequest = append(rawGetRequest, 0)
	request, err := grammar.ParseRequest(rawGetRequest)

	if err != nil {
		t.Error("Error parsing legit GET query")
	}
	if request.Type != grammar.GET_REQUEST {
		t.Error("Query parsing did not recognize the right request type")
	}
	if request.RequestData[0] != "foo" {
		t.Error("Tokens parsed wrongly")
	}
}

func TestParseSetRequest(t *testing.T) {
	var rawSetRequest []byte
	rawSetRequest = append(rawSetRequest, grammar.SET_REQUEST)
	rawSetRequest = append(rawSetRequest, grammar.REQUEST_STATUS)
	rawSetRequest = append(rawSetRequest, []byte("foo")...)
	rawSetRequest = append(rawSetRequest, 0)
	rawSetRequest = append(rawSetRequest, []byte("bar")...)
	request, err := grammar.ParseRequest(rawSetRequest)

	if err != nil {
		t.Error("Error parsing legit SET query")
	}
	if request.Type != grammar.SET_REQUEST {
		t.Error("Query parsing did not recognize the right request type")
	}
	if len(request.RequestData) != 2 || request.RequestData[0] != "foo" || request.RequestData[1] != "bar" {
		t.Error("Tokens parsed wrongly. got: ")
	}
}

func TestParseAuthRequest(t *testing.T) {
	var rawAuthRequest []byte
	rawAuthRequest = append(rawAuthRequest, grammar.AUTH_REQUEST)
	rawAuthRequest = append(rawAuthRequest, grammar.REQUEST_STATUS)
	rawAuthRequest = append(rawAuthRequest, []byte("username")...)
	rawAuthRequest = append(rawAuthRequest, 0)
	rawAuthRequest = append(rawAuthRequest, []byte("password")...)
	request, err := grammar.ParseRequest(rawAuthRequest)

	if err != nil {
		t.Error("Error parsing legit AUTH query")
	}
	if request.Type != grammar.AUTH_REQUEST {
		t.Error("Query parsing did not recognize the right request type")
	}
	if len(request.RequestData) != 2 || request.RequestData[0] != "username" || request.RequestData[1] != "password" {
		t.Error("Tokens parsed wrongly")
	}
}

func TestParseDeleteRequest(t *testing.T) {
	var rawAuthRequest []byte
	rawAuthRequest = append(rawAuthRequest, grammar.DELETE_REQUEST)
	rawAuthRequest = append(rawAuthRequest, grammar.REQUEST_STATUS)
	rawAuthRequest = append(rawAuthRequest, []byte("KEY")...)
	rawAuthRequest = append(rawAuthRequest, 0)
	rawAuthRequest = append(rawAuthRequest, []byte("foo")...)
	request, err := grammar.ParseRequest(rawAuthRequest)

	if err != nil {
		t.Error("Error parsing legit AUTH query")
	}
	if request.Type != grammar.DELETE_REQUEST {
		t.Error("Query parsing did not recognize the right request type")
	}
	if len(request.RequestData) != 2 || request.RequestData[0] != "KEY" || request.RequestData[1] != "foo" {
		t.Error("Tokens parsed wrongly")
	}
}

func TestEmptyRequest(t *testing.T) {
	_, err := grammar.ParseRequest([]byte{99})
	if err == nil {
		t.Error("Succeed in parsing an invalid query")
	}
}

func TestNonsenseRequest(t *testing.T) {
	_, err := grammar.ParseRequest([]byte("1OP2M34IO1P2M3AWMDL;KA,SC;LZXS,C AOWMSXCOIAL MC0123945 I2103965I24-6"))
	if err == nil {
		t.Error("Succeed in parsing an invalid query")
	}
}
