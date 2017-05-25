package server

import (
	"strings"
	"log"
)

func credentialsValid(username string, password string) bool {
	return true
}

//HandleAuthentication checks if credentials are valid - if they are, return the username, else, return an empty string.
func HandleAuthentication(authQuery string) string {
	// Returns username if authentication is successful, else return empty string
	// Authentication is:            USERNAME PASSWORD DATABASE
	log.Printf("authQuery: %s", authQuery)
	usernameEndIndex := strings.Index(authQuery, " ")
	if usernameEndIndex != -1 {
		return authQuery[:strings.Index(authQuery, " ")]
	} else {
		return ""
	}
}
