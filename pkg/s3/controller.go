package s3

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type S3Service interface {
	CreateBucket(c *fiber.Ctx) error
	PutObject(c *fiber.Ctx) error
	ListObjects(c *fiber.Ctx) error
}

// s3Service is the implementation of S3Service interface
type s3Service struct{}

// New this return a new S3Service object
func New() S3Service {
	return &s3Service{}
}

func (s *s3Service) CreateBucket(c *fiber.Ctx) error {
	bucket := c.Params("bucket")
	return c.SendString(fmt.Sprintf("Bucket %s successful created!", bucket))
}

func (s *s3Service) PutObject(c *fiber.Ctx) error {
	bucket := c.Params("bucket")
	key := c.Params("key")
	return c.SendString(fmt.Sprintf("Object %s added in Bucket %s with success!", key, bucket))
}

func (s *s3Service) ListObjects(c *fiber.Ctx) error {
	bucket := c.Params("bucket")
	maxKeys := c.Query("max-keys")
	prefix := c.Query("prefix")
	marker := c.Query("marker")
	return c.SendString(fmt.Sprintf("List objects of Buket %s with maxKeys=%s, prefix=%s, marker=%s", bucket, maxKeys, prefix, marker))
}
