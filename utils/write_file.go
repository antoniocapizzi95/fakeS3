package utils

import (
	"fmt"
	"os"
	"strings"
)

func WriteFile(basePath string, bucketName string, key string, file []byte) error {
	entirePath := GetEntirePath(basePath, bucketName, key)
	dirPath := getDirPath(entirePath)
	err := createDirectoriesFromPath(dirPath)
	if err != nil {
		return fmt.Errorf("error creating folders: %s", err)
	}
	err = os.WriteFile(entirePath, file, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}
	return nil
}

func getDirPath(path string) string {
	directories := strings.Split(path, "/")
	return strings.Join(directories[:len(directories)-1], "/")
}

func createDirectoriesFromPath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}
