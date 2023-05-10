package s3

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

type S3Service interface {
	CreateBucket(c *fiber.Ctx) error
	PutObject(c *fiber.Ctx) error
	ListObjects(c *fiber.Ctx) error
}

// s3Service is the implementation of S3Service interface
type s3Service struct {
	bucketHandler BucketHandler
}

// New this return a new S3Service object
func New(bucketHandler BucketHandler) S3Service {
	return &s3Service{
		bucketHandler: bucketHandler,
	}
}

func (s *s3Service) CreateBucket(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	bucket := buildNewBucket(bucketName)
	err := s.bucketHandler.CreateBucket(c.Context(), bucket)
	if err != nil {
		return err
	}
	return c.SendString(fmt.Sprintf("Bucket %s successful created!", bucketName))
}

func (s *s3Service) PutObject(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	key := c.Params("key")
	object := buildNewObject(key)
	s.bucketHandler.AddObject(c.Context(), bucketName, object)
	return c.SendString(fmt.Sprintf("Object %s added in Bucket %s with success!", key, bucketName))
}

func (s *s3Service) ListObjects(c *fiber.Ctx) error {
	bucket := c.Params("bucket")
	maxKeys := c.Query("max-keys")
	prefix := c.Query("prefix")
	marker := c.Query("marker")
	return c.SendString(fmt.Sprintf("List objects of Buket %s with maxKeys=%s, prefix=%s, marker=%s", bucket, maxKeys, prefix, marker))
}

func buildNewBucket(bucketName string) Bucket {
	return Bucket{
		Name:         bucketName,
		CreationDate: time.Now(),
	}
}

func buildNewObject(key string) Object {
	return Object{
		Key:          key,
		CreationDate: time.Now(),
		LastModified: time.Now(),
		Size:         rand.Uint64(),
	}
}
