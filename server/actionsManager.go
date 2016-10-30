package server

import (
	"beehive"
	"fmt"
	"time"
)

// An Action is something we manage in a priority queue.
type Action struct {
	requestType string
	request     interface{}
	priority    int // The priority of the action in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the action in the heap.
}

func PriorityQueueWorker() {
	// fmt.Println("Entered queue worker")
	// defer fmt.Println("Quit queue worker")
	for {
		if queue.Len() == 0 {
			// fmt.Println("queue is empty :(")
			time.Sleep(1 * time.Millisecond)
		} else {
			var action *Action = queue.Pop().(*Action)
			reqType := action.requestType
			switch reqType {
			case "GET":
				processGetRequest(action.request.(GetRequest))
			case "SET":
				processSetRequest(action.request.(SetRequest))
			}
			fmt.Println("Popped request type: " + action.requestType)
		}
	}
}

func processGetRequest(req GetRequest) {
	val := beehive.Read_from_hard_drive_bucket(req.Key, req.Conn.Bucket)
	req.Conn.Write([]byte(val + "\n"))
}

func processSetRequest(req SetRequest) {
	// Write to hard disk
	beehive.Write_to_hard_drive_bucket(req.Key, req.Value, req.Conn.Bucket)
	req.Conn.Write([]byte(OK + "\n"))
}
