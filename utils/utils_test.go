package utils

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	input := "hello world"
	expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"

	result := CalculateHash([]byte(input))
	if result != expected {
		t.Errorf("Expected hash '%s', got '%s'", expected, result)
	}
	if len(result) != 32 {
		t.Errorf("Expected hash length of 32, got length %d", len(result))
	}
}

func TestExtractByteRanges(t *testing.T) {
	start, end, err := ExtractByteRanges("bytes=5-10")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if start != 5 {
		t.Errorf("expected start to be 5, got %d", start)
	}

	if end != 10 {
		t.Errorf("expected end to be 10, got %d", end)
	}
}

func TestExtractNumber(t *testing.T) {
	num, err := extractNumber("a1b2c3")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if num != 1 {
		t.Errorf("expected number to be 1, got %d", num)
	}
}

func TestGetEntirePath(t *testing.T) {
	basePath := "/home/user"
	bucketName := "my-bucket"
	key := "path/to/file.txt"

	expectedPath := "/home/user/my-bucket/path/to/file.txt"
	actualPath := GetEntirePath(basePath, bucketName, key)

	if actualPath != expectedPath {
		t.Errorf("expected path to be %s, but got %s", expectedPath, actualPath)
	}
}
