package server

// GetRequest describes the GET request to the database - when the client wants to know the value of a certain key
type GetRequest struct {
	Conn *DatabaseConnection
	Key  string
}

//SetRequest describes the SET request to the database - when the client wants to set a value to a certain key
type SetRequest struct {
	Conn     *DatabaseConnection
	Key      string
	Value    string
	TTL      int
	Override bool
}

//DeleteRequest describes the DELETE request to the database - when the client wants to delete a bucket or a key
type DeleteRequest struct {
	Conn       *DatabaseConnection
	ObjectType string
	Object     string
}

//UseRequest is a request that sets the current bucket that is used
type UseRequest struct {
	Conn       *DatabaseConnection
	BucketName string
}

//CreateRequest creates a database file(bucket) in the db folder
type CreateRequest struct {
	Conn       *DatabaseConnection
	BucketName string
}
