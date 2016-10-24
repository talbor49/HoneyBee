package grammar

import (
	"errors"
	"strings"
)

func ParseQuery(query string) (requestType string, parsedTokens []string, err error) {
	query = strings.TrimSpace(query)
	words := strings.Split(query, " ")
	requestType = strings.ToUpper(words[0])

	if requestType == "" {
		return "", nil, errors.New("Request type is blank")
	}

	tokens := words[1:]

	switch requestType {
	case "GET":
		// form: "GET key"
		validLength := len(tokens) >= 1
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("GET request requires a key parameter.")
		}
	case "SET":
		// form: "SET key value"
		validLength := len(tokens) >= 2
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("SET request requires parameters: key and value")
		}
	case "AUTH":
		// form: "AUTH username password bucket"
		validLength := len(tokens) == 3
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("AUTH request should look like 'AUTH username password bucket'")
		}
	case "DELETE":
		// form: "DELETE [BUCKET/KEY] key"
		validLength := len(tokens) == 2
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("A valid DELETE request looks like: 'DELETE [BUCKET/KEY] key'")
		}
	default:
		return requestType, tokens, errors.New("Unimplemented request type")
	}

}