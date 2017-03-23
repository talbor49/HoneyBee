package beehive

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DATA_FOLDER = "data"
)

func BucketExists(bucketName string) bool {
	bucketPath, _ := filepath.Abs(filepath.Join(DATA_FOLDER, bucketName+".hb"))

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		// Bucket does not exist
		f, err := os.Create(bucketPath)
		if err != nil {
			panic(err)
		}
		f.Close()
		return false
	}
	return true
}

func WriteToHardDriveBucket(key string, value string, bucketName string) {
	fmt.Println(bucketName + "->" + key + ":" + value)

	dbPath, _ := filepath.Abs(filepath.Join(DATA_FOLDER, bucketName+".hb"))

	fmt.Println("dbPath: " + dbPath)

	f, err := os.OpenFile(dbPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	value = strings.Replace(value, "\n", "\\n", -1)

	if _, err = f.WriteString(hashedKey + ":" + value + "\n"); err != nil {
		panic(err)
	}
}

func ReadFromHardDriveBucket(key string, bucketName string) string {
	dbPath, _ := filepath.Abs(filepath.Join(DATA_FOLDER, bucketName+".hb"))

	keyHash := sha1.New()
	hashedKey := string(keyHash.Sum([]byte(key)))

	fmt.Println("dbPath: " + dbPath)

	data, err := ioutil.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(string(data), "\n")

	for _, pair := range pairs {
		colonIndex := strings.Index(pair, ":")
		pairKey := pair[:colonIndex]
		if pairKey == hashedKey {
			pairValue := pair[colonIndex+1:]
			return pairValue
		}
	}

	return ""
}

func DeleteFromHardDriveBucket(object string, objectType string, bucketName string) error {
	return nil
}
