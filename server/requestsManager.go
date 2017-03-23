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

// if request is WRITE {
// 	inconsistentKeys.append(key)  // When someone tries to access one of the keys in this list, push it up the priority queue (at least above the GET request)
// }

func handleGetRequest(request GetRequest) string {
	// For now, just blindly push it.
	// TODO: validate that it doesn't contradict any other action in the queue.
	// Check many other stuff?
	// TODO: calculate priority
	PushRequestToActionQueue(request, "GET", 5)
	return OK
}

func handleUseRequest(request UseRequest) string {

	PushRequestToActionQueue(request, "USE", 10)
	return OK
}
