package server

const (
	OK  = "OK"
	ERR = "ERR"
)

func pushSetRequestToQ(request SetRequest) string {
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

func pushGetRequestToQ(request GetRequest) string {
	// For now, just blindly push it.
	// TODO: validate that it doesn't contradict any other action in the queue.
	// Check many other stuff?
	// TODO: calculate priority
	PushRequestToActionQueue(request, "GET", 5)
	return OK
}

func pushUseRequestToQ(request UseRequest) string {

	PushRequestToActionQueue(request, "USE", 10)
	return OK
}

func pushCreateRequesToQ(request CreateRequest) string {
	PushRequestToActionQueue(request, "CREATE", 10)
	return OK
}
