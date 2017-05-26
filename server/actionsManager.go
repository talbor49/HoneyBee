package server

import (
	"fmt"
	"time"

	"github.com/talbor49/HoneyBee/beehive"
	"log"
)

const (
	ERROR_NO_SUCH_BUCKET = "ERROR no bucket named %s exists."
)

// in the background, clean "cold" (unused) records from RAM

// RULE OF THUMB - UPDATE LOGS WHATEVER YOU
// current decision - don't compress keys, only compress values

//QueueRequestsHandler will automatically pop actions from the action priority queue.
//This method will always run as a goroutine.
func QueueRequestsHandler() {
	// log.Println("Entered queue worker")
	// defer log.Println("Quit queue worker")
	for {
		if Queue.Len() == 0 {
			// log.Println("queue is empty :(")
			time.Sleep(1 * time.Millisecond)
		} else {
			var action = PopFromRequestQueue()
			reqType := action.RequestType
			switch reqType {
			case "GET":
				processGetRequest(action.Request.(GetRequest))
			case "SET":
				processSetRequest(action.Request.(SetRequest))
			case "DELETE":
				processDeleteRequest(action.Request.(DeleteRequest))
			case "USE":
				processUseRequest(action.Request.(UseRequest))
			case "CREATE":
				processCreateRequest(action.Request.(CreateRequest))
			case "AUTH":
				processAuthRequest(action.Request.(AuthRequest))
			}
			log.Printf("Popped request type: %s", action.RequestType)
		}
	}
}

func processDeleteRequest(req DeleteRequest) {
	message, _ := beehive.DeleteFromHardDriveBucket(req.Object, req.ObjectType, req.Conn.Bucket)
	req.Conn.Write([]byte(message))
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
	if req.Conn.Bucket == "" {
		req.Conn.Write([]byte("ERROR client needs to authorize before sending requests"))
		return
	}
	if !beehive.BucketExists(req.Conn.Bucket) {
		req.Conn.Write([]byte(errBucketDoesNotExist))
		return
	}

	message, _ := beehive.ReadFromHardDriveBucket(req.Key, req.Conn.Bucket)
	req.Conn.Write([]byte(message + "\n"))
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

	if req.Conn.Bucket == "" {
		req.Conn.Write([]byte("ERROR client needs to authorize before sending requests\n"))
		return
	}
	if !beehive.BucketExists(req.Conn.Bucket) {
		req.Conn.Write([]byte(errBucketDoesNotExist))
		return
	}

	// Write to hard disk
	message, _ := beehive.WriteToHardDriveBucket(req.Key, req.Value, req.Conn.Bucket)
	req.Conn.Write([]byte(message))
}

func processUseRequest(req UseRequest) {
	log.Printf("Checking if there is a database at path: %s", req.BucketName)
	// If the bucket does not exist - create it.
	if beehive.BucketExists(req.BucketName) {
		req.Conn.Bucket = req.BucketName
		req.Conn.Write([]byte(OK + "\n"))
	} else {
		errorMessage := fmt.Sprintf(ERROR_NO_SUCH_BUCKET, req.BucketName)
		req.Conn.Write([]byte(errorMessage))
		log.Println(errorMessage)
	}
}

func processCreateRequest(req CreateRequest) {
	if beehive.BucketExists(req.BucketName) {
		req.Conn.Write([]byte(errBucketAlreadyExists + "\n"))
		return
	}

	message, _ := beehive.CreateHardDriveBucket(req.BucketName)
	req.Conn.Write([]byte(message))
	log.Println(message)
}

func processAuthRequest(req AuthRequest) {
	if credentialsValid(req.Username, req.Password) {
		req.Conn.Username = req.Username
	}
}