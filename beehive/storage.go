package beehive

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Write_to_hard_drive_bucket(key string, value string, database string) {
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

	value = strings.Replace(value, "\n", "\\n", -1)

	if _, err = f.WriteString(hashedKey + ":" + value + "\n"); err != nil {
		panic(err)
	}
}

func Read_from_hard_drive_bucket(key string, database string) string {
	dbPath, _ := filepath.Abs(path.Join("data", database+".hb"))

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
