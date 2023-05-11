package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(basePath string, bucketName string, key string, file []byte) error {
	entirePath := fmt.Sprintf("%s/%s/%s", basePath, bucketName, key)
	dirPath := strings.Join(strings.Split(filepath.Dir(entirePath), "/")[:2], "/")
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

func createDirectoriesFromPath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}
