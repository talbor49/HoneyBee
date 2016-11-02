package server

const (
	OK  = "OK"
	ERR = "ERR"
)

func handleSetRequest(request SetRequest) string {
	// For now, just blindly push it.
	// TODO: validate that it doesn't contradict any other action in the queue.
	// Check many other stuff?
	// TODO: calculate priority
	PushRequestToActionQueue(request, "SET", 6)
	return OK
}
func handleGetRequest(request GetRequest) string {
	// For now, just blindly push it.
	// TODO: validate that it doesn't contradict any other action in the queue.
	// Check many other stuff?
	// TODO: calculate priority
	PushRequestToActionQueue(request, "GET", 5)
	return OK
}

// in the background, clean "cold" (unused) records from RAM

// RULE OF THUMB - UPDATE LOGS WHATEVER YOU DO

// current decision - don't compress keys, only compress values

// priorityQue.Push(query)
// if request is WRITE {
// 	inconsistentKeys.append(key)  // When someone tries to access one of the keys in this list, push it up the priority queue (at least above the GET request)
// }
