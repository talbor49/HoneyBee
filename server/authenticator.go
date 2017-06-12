package server

import (
	"github.com/talbor49/HoneyBee/beehive"
	"crypto/sha1"
	"io"
)

const SALTS_BUCKET = "salts_bucket"
const USERS_BUCKET = "users_bucket"


func credentialsValid(username string, password string) bool {
	salt, err := beehive.ReadFromHardDriveBucket(username, SALTS_BUCKET)
	if err != nil { return false }
	saltedPassword := password + salt

	hashedSaltedPassword := hash(saltedPassword)

	realHashedSaltedPassword, _ := beehive.ReadFromHardDriveBucket(username, USERS_BUCKET)

	if err != nil { return false }

	return hashedSaltedPassword == realHashedSaltedPassword
}

func hash(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return string(h.Sum(nil))
}