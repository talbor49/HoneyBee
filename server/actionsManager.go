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


func processDeleteRequest(req DeleteRequest) (response grammar.Response) {
	response.Type = grammar.DELETE_RESPONSE
	status, err := beehive.DeleteFromHardDriveBucket(req.Object, req.ObjectType, req.Conn.Bucket)
	response.Status = status
	response.Data = err.Error()
	return
}

func processGetRequest(req GetRequest) (response grammar.Response) {
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
	response.Type = grammar.GET_RESPONSE
	if req.Conn.Bucket == "" {
		response.Status = grammar.RESP_STATUS_ERR_UNAUTHORIZED
		return
	}
	if !beehive.BucketExists(req.Conn.Bucket) {
		response.Status = grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
		return
	}

	message, err := beehive.ReadFromHardDriveBucket(req.Key, req.Conn.Bucket)
	if err != nil {
		response.Status = grammar.RESP_STATUS_ERR_KEY_NOT_FOUND
		return
	}
	response.Status = grammar.RESP_STATUS_SUCCESS
	response.Data = message
	return
}

func processSetRequest(req SetRequest) (response grammar.Response){
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

	response.Type = grammar.SET_RESPONSE
	if req.Conn.Bucket == "" {
		response.Status = grammar.RESP_STATUS_ERR_UNAUTHORIZED
		return
	}
	if !beehive.BucketExists(req.Conn.Bucket) {
		response.Status = grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
		return
	}

	log.Printf("Setting %s->%s in bucket %s", req.Key, req.Value, req.Conn.Bucket)
	// Write to hard disk
	status, err := beehive.WriteToHardDriveBucket(req.Key, req.Value, req.Conn.Bucket)
	response.Status = status
	response.Data = err.Error()
	return
}

func processUseRequest(req UseRequest) (response grammar.Response){
	response.Type = grammar.USE_RESPONSE
	log.Printf("Checking if there is a database at path: %s", req.BucketName)
	// If the bucket does not exist - create it.
	if beehive.BucketExists(req.BucketName) {
		req.Conn.Bucket = req.BucketName
		response.Status = grammar.RESP_STATUS_SUCCESS
	} else {
		response.Status = grammar.RESP_STATUS_ERR_NO_SUCH_BUCKET
		log.Printf("Error - no bucket named %s found on disk.", req.BucketName)
	}
	return
}

func processCreateRequest(req CreateRequest) (response grammar.Response){
	response.Type= grammar.CREATE_RESPONSE
	if beehive.BucketExists(req.BucketName) {
		response.Status = grammar.RESP_STATUS_ERR_BUCKET_ALREADY_EXISTS
		return
	}

	status, err := beehive.CreateHardDriveBucket(req.BucketName)
	response.Status = status
	response.Data = err.Error()
	return
}

func processAuthRequest(req AuthRequest) (response grammar.Response) {
	response.Type = grammar.AUTHORIZE_RESPONSE
	if credentialsValid(req.Username, req.Password) {
		req.Conn.Username = req.Username
		log.Printf("User logged in as: %s", req.Username)
		response.Status = grammar.RESP_STATUS_SUCCESS
	} else {
		response.Status = grammar.RESP_STATUS_ERR_WRONG_CREDENTIALS
	}
	return
}