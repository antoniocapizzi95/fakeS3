package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func CalculateHash(file *os.File) string {

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	hashInString := fmt.Sprintf("%x", hash.Sum(nil))
	return hashInString[:32]
}
