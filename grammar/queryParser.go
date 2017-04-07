package grammar

import (
	"errors"
	"fmt"
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
		validLength := len(tokens) == 2
		if validLength {
			return requestType, tokens, nil
		} else {
			fmt.Println("Tokens recieved: ")
			fmt.Print(tokens)
			return "", nil, errors.New("AUTH request should look like 'AUTH username password'")
		}
	case "DELETE":
		// form: "DELETE [BUCKET/KEY] key"
		validLength := len(tokens) >= 2
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("A valid DELETE request looks like: 'DELETE [BUCKET/KEY] <name>'")
		}
	case "USE":
		validLength := len(tokens) == 1
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("A valid USE request looks like: 'USE <bucket>'")
		}
	case "CREATE":
		validLength := len(tokens) == 1
		if validLength {
			return requestType, tokens, nil
		} else {
			return "", nil, errors.New("A valid CREATE request looks like: 'CREATE <bucket>'")
		}
	default:
		return "", tokens, errors.New("Unimplemented request type")
	}

}
