package server

import (
	"fmt"
	"strings"
)

func credentialsValid(username string, password string) bool {
	return true
}

func HandleAuthentication(authQuery string) string {
	// Returns username if authentication is successful, else return empty string
	// Authentication is:            USERNAME PASSWORD DATABASE
	fmt.Println("authQuery: " + authQuery)
	usernameEndIndex := strings.Index(authQuery, " ")
	if usernameEndIndex != -1 {
		return authQuery[:strings.Index(authQuery, " ")]
	} else {
		return ""
	}
}
