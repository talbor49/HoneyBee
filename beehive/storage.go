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

<<<<<<< HEAD
=======
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

>>>>>>> 84b6353f02a9bd6662913c839e967d330cb40c0d
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

	for i := len(pairs) - 1; i >= 0; i-- {
		pair := pairs[i]
		colonIndex := strings.Index(pair, ":")
<<<<<<< HEAD
		if colonIndex == -1 {
=======
		if colonIndex <= 0 {
>>>>>>> 84b6353f02a9bd6662913c839e967d330cb40c0d
			continue
		}
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
