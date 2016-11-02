package server

import (
	"fmt"
	"time"

	"github.com/talbor49/HoneyBee/beehive"
)

// in the background, clean "cold" (unused) records from RAM

// RULE OF THUMB - UPDATE LOGS WHATEVER YOU DO

// current decision - don't compress keys, only compress values

func PriorityQueueWorker() {
	// fmt.Println("Entered queue worker")
	// defer fmt.Println("Quit queue worker")
	for {
		if Queue.Len() == 0 {
			// fmt.Println("queue is empty :(")
			time.Sleep(1 * time.Millisecond)
		} else {
			var action *Action = Queue.Pop().(*Action)
			reqType := action.RequestType
			switch reqType {
			case "GET":
				processGetRequest(action.Request.(GetRequest))
			case "SET":
				processSetRequest(action.Request.(SetRequest))
			}
			fmt.Println("Popped request type: " + action.RequestType)
		}
	}
}

func processGetRequest(req GetRequest) {
	/*
		if IS IN RAM {
			return FROM RAM
		} ELSE IF IS IN HARD DISK {
			// calculate if record is hot enough to be put in RAM
			return FROM HARD DISK
		} else {
			return NOT FOUND
		}
	*/
	val := beehive.Read_from_hard_drive_bucket(req.Key, req.Conn.Bucket)
	req.Conn.Write([]byte(val + "\n"))
}

func processSetRequest(req SetRequest) {
	/*
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
	// Write to hard disk
	beehive.Write_to_hard_drive_bucket(req.Key, req.Value, req.Conn.Bucket)
	req.Conn.Write([]byte(OK + "\n"))
}
