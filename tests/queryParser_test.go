package tests

import (
	"testing"

	"github.com/talbor49/HoneyBee/grammar"
)

func TestParseGetRequest(t *testing.T) {
	requestType, parsedTokens, err := grammar.ParseQuery("GET foo")

	if err != nil {
		t.Error("Error parsing legit query")
	}
	if requestType != "GET" {
		t.Error("Query parsing did not recognize the right request type")
	}
	if parsedTokens[0] != "foo" {
		t.Error("Tokens parsed wrongly")
	}

	_, _, err = grammar.ParseQuery("GET")

	if err == nil {
		t.Error("Succeed in parsing an invalid query")
	}

	// requestType, parsedTokens, err = grammar.ParseQuery("GET ASMDQWE2309123 123====SAC A-S---- ;LQAWE,ZWD;LAZOW,E1PO;243KE0P2O1-3EOI90135I1123   12333333333123123\"\" --- ;")

}

func TestParseSetRequest(t *testing.T) {
	requestType, parsedTokens, err := grammar.ParseQuery("SET foo bar")

	if err != nil {
		t.Error("Error parsing legit query")
	}
	if requestType != "SET" {
		t.Error("Query parsing did not recognize the right request type")
	}
	if len(parsedTokens) != 2 || parsedTokens[0] != "foo" || parsedTokens[1] != "bar" {
		t.Error("Tokens parsed wrongly")
	}

	_, _, err = grammar.ParseQuery("SET")

	if err == nil {
		t.Error("Succeed in parsing an invalid query")
	}
}

func TestParseAuthRequest(t *testing.T) {
	requestType, parsedTokens, err := grammar.ParseQuery("AUTH username password bucket")

	if err != nil {
		t.Error("Error parsing legit query")
	}
	if requestType != "AUTH" {
		t.Error("Query parsing did not recognize the right request type")
	}
	if len(parsedTokens) != 3 || parsedTokens[0] != "username" || parsedTokens[1] != "password" || parsedTokens[2] != "bucket" {
		t.Error("Tokens parsed wrongly")
	}

	_, _, err = grammar.ParseQuery("AUTH ASMDQWE2309123 123====SAC A-S---- ;LQAWE,ZWD;LAZOW,E1PO;243KE0P2O1-3EOI90135I1123   12333333333123123\"\" --- ;")
	if err == nil {
		t.Error("Succeed in parsing an invalid query")
	}
}

func TestParseDeleteRequest(t *testing.T) {
	requestType, parsedTokens, err := grammar.ParseQuery("DELETE KEY foo")

	if err != nil {
		t.Error("Error parsing legit query")
	}
	if requestType != "DELETE" {
		t.Error("Query parsing did not recognize the right request type")
	}
	if len(parsedTokens) != 2 || parsedTokens[0] != "KEY" || parsedTokens[1] != "foo" {
		t.Error("Tokens parsed wrongly")
	}
}

func TestEmptyRequest(t *testing.T) {
	requestType, _, err := grammar.ParseQuery("")
	if err == nil || requestType != "" {
		t.Error("Succeed in parsing an invalid query")
	}
}

func TestNonsenseRequest(t *testing.T) {
	requestType, _, err := grammar.ParseQuery("1OP2M34IO1P2M3AWMDL;KA,SC;LZXS,C AOWMSXCOIAL MC0123945 I2103965I24-6")
	if err == nil || requestType != "" {
		t.Error("Succeed in parsing an invalid query")
	}
}
