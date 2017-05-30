package server

import (
	"time"

	"github.com/talbor49/HoneyBee/beehive"
	"log"
	"github.com/talbor49/HoneyBee/grammar"
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
			case grammar.GET_REQUEST:
				processGetRequest(action.Request.(GetRequest))
			case grammar.SET_REQUEST:
				processSetRequest(action.Request.(SetRequest))
			case grammar.DELETE_REQUEST:
				processDeleteRequest(action.Request.(DeleteRequest))
			case grammar.USE_REQUEST:
				processUseRequest(action.Request.(UseRequest))
			case grammar.CREATE_REQUEST:
				processCreateRequest(action.Request.(CreateRequest))
			case grammar.AUTH_REQUEST:
				processAuthRequest(action.Request.(AuthRequest))
			}
			log.Printf("Popped request type: %s", action.RequestType)
		}
	}
}

func processDeleteRequest(req DeleteRequest) {
	response := grammar.Response{Type:grammar.DELETE_FINAL_ANSWER}
	status, err := beehive.DeleteFromHardDriveBucket(req.Object, req.ObjectType, req.Conn.Bucket)
	response.Status = status
	response.Data = err.Error()
	req.Conn.Write(grammar.GetBufferFromResponse(response))
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
	response := grammar.Response{Type:grammar.GET_FINAL_ANSWER}
	if req.Conn.Bucket == "" {
		response.Status = grammar.RESP_STATUS_ERR_UNAUTHORIZED
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		return
	}
	if !beehive.BucketExists(req.Conn.Bucket) {
		response.Status = grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		return
	}

	message, err := beehive.ReadFromHardDriveBucket(req.Key, req.Conn.Bucket)
	if err != nil {
		response.Status = grammar.RESP_STATUS_ERR_KEY_NOT_FOUND
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		return
	}
	response.Status = grammar.RESP_STATUS_SUCCESS
	response.Data = message
	req.Conn.Write(grammar.GetBufferFromResponse(response))
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

	response := grammar.Response{Type:grammar.SET_FINAL_ANSWER}
	if req.Conn.Bucket == "" {
		response.Status = grammar.RESP_STATUS_ERR_UNAUTHORIZED
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		return
	}
	if !beehive.BucketExists(req.Conn.Bucket) {
		response.Status = grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		return
	}

	log.Printf("Setting %s->%s in bucket %s", req.Key, req.Value, req.Conn.Bucket)
	// Write to hard disk
	status, err := beehive.WriteToHardDriveBucket(req.Key, req.Value, req.Conn.Bucket)
	response.Status = status
	response.Data = err.Error()
	req.Conn.Write(grammar.GetBufferFromResponse(response))
}

func processUseRequest(req UseRequest) {
	response := grammar.Response{Type:grammar.USE_FINAL_ANSWER}
	log.Printf("Checking if there is a database at path: %s", req.BucketName)
	// If the bucket does not exist - create it.
	if beehive.BucketExists(req.BucketName) {
		req.Conn.Bucket = req.BucketName
		response.Status = grammar.RESP_STATUS_SUCCESS
		req.Conn.Write(grammar.GetBufferFromResponse(response))
	} else {
		response.Status = grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		log.Printf("Error - no bucket named %s found on disk.", req.BucketName)
	}
}

func processCreateRequest(req CreateRequest) {
	response := grammar.Response{Type:grammar.CREATE_FINAL_ANSWER}
	if beehive.BucketExists(req.BucketName) {
		response.Status = grammar.RESP_STATUS_ERR_BUCKET_ALREADY_EXISTS
		req.Conn.Write(grammar.GetBufferFromResponse(response))
		return
	}

	status, err := beehive.CreateHardDriveBucket(req.BucketName)
	response.Status = status
	response.Data = err.Error()
	req.Conn.Write(grammar.GetBufferFromResponse(response))
}

func processAuthRequest(req AuthRequest) {
	response := grammar.Response{Type:grammar.AUTHORIZE_FINAL_ANSWER}
	if credentialsValid(req.Username, req.Password) {
		req.Conn.Username = req.Username
		response.Status = grammar.RESP_STATUS_SUCCESS
	} else {
		response.Status = grammar.RESP_STATUS_ERR_WRONG_CREDENTIALS
	}
	req.Conn.Write(grammar.GetBufferFromResponse(response))
}