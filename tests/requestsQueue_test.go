package tests

import (
	"github.com/talbor49/HoneyBee/server"
	"testing"
)

func TestNullValues(t *testing.T) {
	if server.PopFromRequestQueue() != nil {
		t.Error("Pop from empty queue didn't return nil")
	}
}

func TestSanity(t *testing.T) {
	setRequest1 := server.SetRequest{Key: "foo", Value: "bar", Conn: nil}
	setRequest2 := server.SetRequest{Key: "gooooo", Value: "bar", Conn: nil}
	setRequest3 := server.SetRequest{Key: "bar", Value: "foo", Conn: nil}
	setRequest4 := server.SetRequest{Key: "Marco", Value: "Polo", Conn: nil}

	server.InitPriorityQueue()
	server.PushRequestToActionQueue(setRequest1, "SET", 123)
	server.PushRequestToActionQueue(setRequest2, "SET", 111)
	server.PushRequestToActionQueue(setRequest3, "SET", 155)
	server.PushRequestToActionQueue(setRequest4, "SET", 104)

	if server.Queue.Len() == 0 {
		t.Error("Queue was empty even though requests were pushed and not popped.")
	}

	poppedSetRequest3 := server.PopFromRequestQueue().Request.(server.SetRequest)

	// log.Println("First pop key: " + poppedSetRequest3.Key)
	if poppedSetRequest3.Key != setRequest3.Key || poppedSetRequest3.Value != setRequest3.Value {
		t.Error("Error in priority queue - request did not pop in the correct order")
	}

	poppedSetRequest1 := server.PopFromRequestQueue().Request.(server.SetRequest)

	// log.Println("Second pop key: " + poppedSetRequest1.Key)
	if poppedSetRequest1.Key != setRequest1.Key || poppedSetRequest1.Value != setRequest1.Value {
		t.Error("Error in priority queue - request did not pop in the correct order")
	}

	poppedSetRequest2 := server.PopFromRequestQueue().Request.(server.SetRequest)

	// log.Println("Third pop key: " + poppedSetRequest2.Key)
	if poppedSetRequest2.Key != setRequest2.Key || poppedSetRequest2.Value != setRequest2.Value {
		t.Error("Error in priority queue - request did not pop in the correct order")
	}

	poppedSetRequest4 := server.PopFromRequestQueue().Request.(server.SetRequest)

	// log.Println("Fourth pop key: " + poppedSetRequest4.Key)
	if poppedSetRequest4.Key != setRequest4.Key || poppedSetRequest4.Value != setRequest4.Value {
		t.Error("Error in priority queue - request did not pop in the correct order")
	}
}
