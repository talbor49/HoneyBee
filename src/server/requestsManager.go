package server

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const (
	OK = "OK"
)

func handleSetRequest(request SetRequest) string {
	return OK
}
func handleGetRequest(request GetRequest) string {
	return OK
}

func write_to_hard_disk(key string, value string, database string) {
	fmt.Println(database + "->" + key + ":" + value)

	dbPath, _ := filepath.Abs(path.Join("data", database+".hb"))

	fmt.Println("dbPath: " + dbPath)

	f, err := os.OpenFile(dbPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	if _, err = f.WriteString(hashedKey + ":" + value + "\n"); err != nil {
		panic(err)
	}
}
