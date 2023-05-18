package utils

import (
	"fmt"
	"os"
)

func ReadFileRange(path string, start, end int) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := stat.Size()

	if end >= int(fileSize) {
		end = int(fileSize) - 1
	}

	if start >= end {
		return nil, fmt.Errorf("invalid range")
	}

	blockSize := end - start + 1

	if _, err := file.Seek(int64(start), 0); err != nil {
		return nil, err
	}

	content := make([]byte, blockSize)
	if _, err := file.Read(content); err != nil {
		return nil, err
	}

	return content, nil
}
