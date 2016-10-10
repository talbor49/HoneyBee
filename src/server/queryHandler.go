package server

import (
	"fmt"
	"strings"
)

func HandleQuery(query string) {
	// PUT RECORD {KEY} {VALUE} IN {BUCKET}
	
	// TODO: write query in plain text to log
	
	// queriesQueue priorityQue
	
	/*
		Parse key / value / stuff
	*/
	
	// in the background, clean "cold" (unused) records from RAM
	
	// RULE OF THUMB - UPDATE LOGS WHATEVER YOU DO
	
	// current decision - don't compress keys, only compress values
	
	// priorityQue.Push(query)
	// if request is WRITE {
	// 	inconsistentKeys.append(key)  // When someone tries to access one of the keys in this list, push it up the priority queue (at least above the GET request)
	// }
	
	/*
	if GET REQUEST {
		if IS IN RAM {
			return FROM RAM
		} ELSE IF IS IN HARD DISK {
			// calculate if record is hot enough to be put in RAM
			return FROM HARD DISK
		} else {
			return NOT FOUND
		}
	}
	else if SET REQUEST {
		FIRST:
			// DECIDE IF TO KEEP A POINTER TO THE VALUE IN MEMORY OR THE VALUE OF ITSELF
			PUT IN RAM
			REMOVE FROM INCONSISTENTKEYS
		THEN:
			// COMPRESS VALUE WHEN WRITING TO HARD DISK
			PUT IN HARD DISK
			UPDATE CACHED MEMORY
	}
	*/
		
	
	fmt.Println("query: " + query)
}

func HandleAuthentication(authQuery string) string {
	// Authentication is:            USERNAME PASSWORD DATABASE
	fmt.Println("authQuery: " + authQuery)
	usernameEndIndex := strings.Index(authQuery, " ")
	if usernameEndIndex != -1 {
		return authQuery[:strings.Index(authQuery, " ")]
	} else {
		return ""
	}
}
