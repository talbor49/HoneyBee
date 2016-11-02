package main

import (
	"fmt"
	"server"
	"testing"
)

func main() {
	server.InitPriorityQueue()
	server.StartServer()
}

func TestSanity(t *testing.T) {
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	server.InitPriorityQueue()

	// heap.Push(&Queue, getRequest1)
	// heap.Push(&Queue, getRequest2)
	// heap.Push(&Queue, getRequest3)
	getRequest1 := server.GetRequest{Key: "shouldgo1", Conn: nil}
	getRequest2 := server.GetRequest{Key: "shouldgo3", Conn: nil}
	getRequest3 := server.GetRequest{Key: "shouldgo2", Conn: nil}

	server.PushRequestToActionQueue(getRequest1, "GET", 40)
	server.PushRequestToActionQueue(getRequest2, "GET", 20)
	server.PushRequestToActionQueue(getRequest3, "GET", 30)
	// heap.Push(&Queue, &Action{Request: getRequest1, priority: 40, RequestType: "GET"})
	// heap.Push(&Queue, &Action{Request: getRequest2, priority: 20, RequestType: "GET"})
	// heap.Push(&Queue, &Action{Request: getRequest3, priority: 30, RequestType: "GET"})

	if server.Queue.Len() == 0 {
		t.Error("Queue was empty even though requests were pushed and not popped.")
	}

	poppedGetRequest1 := server.PopFromRequestQueue().Request.(server.GetRequest)

	fmt.Println("first pop key: " + poppedGetRequest1.Key)
	if poppedGetRequest1.Key != "shouldgo1" {
		t.Error("Error in priority queue")
	}

	poppedGetRequest3 := server.PopFromRequestQueue().Request.(server.GetRequest)

	fmt.Println("second pop key: " + poppedGetRequest3.Key)
	if poppedGetRequest3.Key != "shouldgo2" {
		t.Error("Error in priority queue")
	}

	poppedGetRequest2 := server.PopFromRequestQueue().Request.(server.GetRequest)

	fmt.Println("first pop key: " + poppedGetRequest2.Key)
	if poppedGetRequest2.Key != "shouldgo3" {
		t.Error("Error in priority queue")
	}
}
