package beehive

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	dataFolder = "data"
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

func WriteToHardDriveBucket(key string, value string, bucketName string) (string, error) {
	fmt.Println(bucketName + "->" + key + ":" + value)

	bucketPath := getBucketPath(bucketName)

	fmt.Println("bucketPath: " + bucketPath)

	f, err := os.OpenFile(bucketPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	value = strings.Replace(value, "\n", "\\n", -1)

	if _, err = f.WriteString(hashedKey + ":" + value + "\n"); err != nil {
		return "Error while trying to write key to bucket\n", err
	}
	return "Succesfully wrote key to bucket\n", nil
}

func ReadFromHardDriveBucket(key string, bucketName string) (string, error) {
	bucketPath := getBucketPath(bucketName)

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	fmt.Println("bucketPath: " + bucketPath)

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

func CreateHardDriveBucket(bucketName string) (string, error) {
	bucketPath := getBucketPath(bucketName)
	fmt.Println("Creating bucket: " + bucketName + " on path" + bucketPath)
	_, err := os.Create(bucketPath)
	if err != nil {
		return "Error while creating bucket\n", err
	}
	return ("Successfully created bucket\n"), err
}

func DeleteFromHardDriveBucket(object string, objectType string, bucketName string) (string, error) {
	return "", nil
}
