package server

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
