package server

type Request interface {
	ExecuteRequest() string
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type GetRequest struct {
	Conn *DatabaseConnection
	Key  string
}

func (r GetRequest) ExecuteRequest() string {
	return "OK"
}

type SetRequest struct {
	Conn     *DatabaseConnection
	Key      string
	Value    string
	TTL      int
	Override bool
}

func (r SetRequest) ExecuteRequest() string {
	return "OK"
}
