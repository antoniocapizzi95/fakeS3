package utils

import "fmt"

func GetEntirePath(basePath string, bucketName string, key string) string {
	return fmt.Sprintf("%s/%s/%s", basePath, bucketName, key)
}
