package server

import "github.com/talbor49/HoneyBee/grammar"

const (
	OK  = "OK"
	ERR = "ERR"
)

func pushSetRequestToQ(request SetRequest) string {
	// For now, just blindly push it.
	// TODO: validate that it doesn't contradict any other action in the queue.
	// Check many other stuff?
	// TODO: calculate priority
	PushRequestToActionQueue(request, grammar.SET_REQUEST, 6)
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
	PushRequestToActionQueue(request, grammar.GET_REQUEST, 5)
	return OK
}

func pushUseRequestToQ(request UseRequest) string {

	PushRequestToActionQueue(request, grammar.USE_REQUEST, 10)
	return OK
}

func pushCreateRequestToQ(request CreateRequest) string {
	PushRequestToActionQueue(request, grammar.CREATE_REQUEST, 10)
	return OK
}

func pushAuthRequestToQ(request AuthRequest) string {
	PushRequestToActionQueue(request, grammar.AUTH_REQUEST, 10)
	return OK
}
