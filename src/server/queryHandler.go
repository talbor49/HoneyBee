package server

import (
	"fmt"
	"strings"
)

func HandleQuery(query string) {
	fmt.Println("query: " + query)
}

func HandleAuthentication(authQuery string) string {
	fmt.Println("authQuery: " + authQuery)
	usernameEndIndex := strings.Index(authQuery, " ")
	if usernameEndIndex != -1 {
		return authQuery[:strings.Index(authQuery, " ")]
	} else {
		return ""
	}
}
