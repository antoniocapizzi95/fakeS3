package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteFile(basePath string, bucketName string, key string, file io.Reader) (string, error) {
	entirePath := GetEntirePath(basePath, bucketName, key)
	dirPath := getDirPath(entirePath)
	err := createDirectoriesFromPath(dirPath)
	if err != nil {
		return "", fmt.Errorf("error creating folders: %s", err)
	}
	out, err := os.Create(entirePath)
	if err != nil {
		return "", err
	}
	hash := CalculateHash(out)
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return hash, nil
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
