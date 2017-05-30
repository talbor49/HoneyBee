package beehive

import (
	"crypto/sha1"
	"errors"
	"log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/talbor49/HoneyBee/grammar"
)

const (
	dataFolder = "data"
)

const (
	KEY_WRITE_SUCCESS = iota
	KEY_WRITE_FAILURE = iota
)

func getBucketPath(bucketName string) string {
	bucketPath, err := filepath.Abs(filepath.Join(dataFolder, bucketName+".hb"))
	if err != nil {
		panic(err)
	}
	return bucketPath
}

func BucketExists(bucketName string) bool {
	bucketPath := getBucketPath(bucketName)

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		// Bucket does not exist
		return false
	}
	return true
}

func WriteToHardDriveBucket(key string, value string, bucketName string) (byte, error) {
	log.Printf("Setting data in bucket %s -> %s:%s", bucketName, key, value)

	bucketPath := getBucketPath(bucketName)

	log.Printf("Bucket path: %s", bucketPath)

	f, err := os.OpenFile(bucketPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	value = strings.Replace(value, "\n", "\\n", -1)

	if _, err = f.WriteString(hashedKey + ":" + value + "\n"); err != nil {
		return grammar.RESP_STATUS_ERR_COULD_NOT_WRITE_TO_DISK, err
	}
	return grammar.RESP_STATUS_SUCCESS, nil
}

func ReadFromHardDriveBucket(key string, bucketName string) (result string, error error) {
	bucketPath := getBucketPath(bucketName)

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	log.Printf("Bucket path: %s", bucketPath)

	data, err := ioutil.ReadFile(bucketPath)
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(string(data), "\n")

	for i := len(pairs) - 1; i >= 0; i-- {
		pair := pairs[i]
		colonIndex := strings.Index(pair, ":")
		if colonIndex <= 0 {
			continue
		}
		pairKey := pair[:colonIndex]
		if pairKey == hashedKey {
			pairValue := pair[colonIndex+1:]
			return pairValue, nil
		}
	}

	return "", errors.New("Key not found")
}

func CreateHardDriveBucket(bucketName string) (byte, error) {
	bucketPath := getBucketPath(bucketName)
	log.Printf("Creating bucket: %s in path %s", bucketName, bucketPath)
	_, err := os.Create(bucketPath)
	if err != nil {
		return grammar.RESP_STATUS_ERR_COULD_NOT_CREATE_BUCKET, err
	}
	return grammar.RESP_STATUS_SUCCESS, err
}

func DeleteFromHardDriveBucket(object string, objectType string, bucketName string) (status byte, err error) {
	// TODO implement
	return grammar.RESP_STATUS_SUCCESS, nil
}
