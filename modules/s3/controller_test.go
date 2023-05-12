package s3

import (
	"github.com/antoniocapizzi95/fakeS3/models"
	"github.com/antoniocapizzi95/fakeS3/utils"
	"testing"
	"time"
)

func TestAppendOrUpdateObject(t *testing.T) {
	// Create a slice of objects for testing
	objects := []models.Object{
		{Key: "abc", CreationDate: time.Now()},
		{Key: "def", CreationDate: time.Now()},
		{Key: "ghi", CreationDate: time.Now()},
	}

	// Test appending a new object
	newObj := models.Object{Key: "jkl", CreationDate: time.Now()}
	result := appendOrUpdateObject(objects, newObj)
	if len(result) != 4 {
		t.Errorf("Expected result length of 4, got %d", len(result))
	}
	if result[3].Key != "jkl" {
		t.Errorf("Expected object with key 'jkl', got key '%s'", result[3].Key)
	}

	// Test updating an existing object
	newObj = models.Object{Key: "def", CreationDate: time.Now()}
	result = appendOrUpdateObject(objects, newObj)
	if len(result) != 3 {
		t.Errorf("Expected result length of 3, got %d", len(result))
	}
}

func TestFilterObjectsByPrefix(t *testing.T) {
	// Create a slice of objects for testing
	objects := []models.Object{
		{Key: "abc", CreationDate: time.Now()},
		{Key: "def", CreationDate: time.Now()},
		{Key: "ghi", CreationDate: time.Now()},
	}

	// Test filtering by prefix
	result := filterObjectsByPrefix(objects, "ab")
	if len(result) != 1 {
		t.Errorf("Expected result length of 1, got %d", len(result))
	}
	if result[0].Key != "abc" {
		t.Errorf("Expected object with key 'abc', got key '%s'", result[0].Key)
	}
}

func TestBuildListOutput(t *testing.T) {
	// Test building ListObjectsOutput
	bucketName := "test-bucket"
	maxKeys := 100
	prefix := "test/"
	marker := "abc"
	objects := []models.Object{
		{Key: "test/abc", CreationDate: time.Now()},
		{Key: "test/def", CreationDate: time.Now()},
	}
	expected := models.ListObjectsOutput{
		Name:        bucketName,
		Prefix:      prefix,
		MaxKeys:     maxKeys,
		Marker:      marker,
		Contents:    objects,
		IsTruncated: false,
	}
	result := buildListOutput(bucketName, maxKeys, prefix, marker, objects)
	if !utils.CompareStructs(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetObject(t *testing.T) {
	// Create a bucket for testing
	bucket := models.Bucket{
		Name: "test-bucket",
		Objects: []models.Object{
			{Key: "abc", CreationDate: time.Now()},
			{Key: "def", CreationDate: time.Now()},
			{Key: "ghi", CreationDate: time.Now()},
		},
	}

	// Test getting an existing object
	key := "def"
	result := getObject(bucket, key)
	if result == nil {
		t.Errorf("Expected object with key '%s', got nil", key)
	}
	if result.Key != key {
		t.Errorf("Expected object with key '%s', got key '%s'", key, result.Key)
	}

	// Test getting a non-existing object
	key = "jkl"
	result = getObject(bucket, key)
	if result != nil {
		t.Errorf("Expected nil object, got object '%s'", result.Key)
	}
}
