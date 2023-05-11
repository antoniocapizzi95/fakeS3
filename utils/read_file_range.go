package utils

import (
	"fmt"
	"os"
)

func ReadFileRange(path string, start, end int) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if end >= len(content) {
		end = len(content) - 1
	}

	if start >= end {
		return nil, fmt.Errorf("invalid range")
	}

	return content[start : end+1], nil
}
