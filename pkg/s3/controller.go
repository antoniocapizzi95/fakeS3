package s3

import (
	"encoding/xml"
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/config"
	"github.com/antoniocapizzi95/fakeS3/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

type S3Service interface {
	CreateBucket(c *fiber.Ctx) error
	PutObject(c *fiber.Ctx) error
	ListObjects(c *fiber.Ctx) error
	GetObject(c *fiber.Ctx) error
}

// s3Service is the implementation of S3Service interface
type s3Service struct {
	bucketHandler BucketHandler
	conf          config.Config
}

// New this return a new S3Service object
func New(bucketHandler BucketHandler, conf config.Config) S3Service {
	return &s3Service{
		bucketHandler: bucketHandler,
		conf:          conf,
	}
}

func (s *s3Service) CreateBucket(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	bucket := buildNewBucket(bucketName)
	err := s.bucketHandler.CreateBucket(c.Context(), bucket)
	if err != nil {
		return err
	}
	return c.Status(200).Type("application/xml").SendString("")
}

func (s *s3Service) PutObject(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	key := c.Params("+")
	ctx := c.Context()
	file := c.Body()
	object := buildNewObject(key, file)
	bucket, err := s.bucketHandler.GetBucket(ctx, bucketName)
	if err != nil {
		return err
	}
	if bucket == nil {
		return fmt.Errorf("bucket with name %s not found", bucketName)
	}

	err = utils.WriteFile(s.conf.StoragePath, bucketName, key, file)
	if err != nil {
		return err
	}
	bucket.Objects = appendOrUpdateObject(bucket.Objects, object)
	err = s.bucketHandler.UpdateBucket(ctx, *bucket)
	if err != nil {
		return err
	}
	c.Set("ETag", object.ETag)
	return c.Status(200).Type("application/xml").SendString("")
}

func (s *s3Service) ListObjects(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	maxKeys, _ := strconv.Atoi(c.Query("max-keys"))
	prefix := c.Query("prefix")
	marker := c.Query("marker")
	bucket, err := s.bucketHandler.GetBucket(c.Context(), bucketName)
	if err != nil {
		return err
	}
	if bucket == nil {
		return fmt.Errorf("bucket with name %s not found", bucketName)
	}

	objects := bucket.Objects
	if prefix != "" {
		objects = filterObjectsByPrefix(objects, prefix)
	}
	output := buildListOutput(bucketName, int(maxKeys), prefix, marker, objects)
	xmlBytes, err := xml.Marshal(output)
	if err != nil {
		return fmt.Errorf("error marshaling to XML: %v", err)
	}
	xmlString := string(xmlBytes)
	return c.Status(fiber.StatusOK).Type("application/xml").SendString(xmlString)
}

func (s *s3Service) GetObject(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	key := c.Params("+")
	path := utils.GetEntirePath(s.conf.StoragePath, bucketName, key)
	bucket, err := s.bucketHandler.GetBucket(c.Context(), bucketName)
	if err != nil {
		return err
	}

	object := getObject(*bucket, key)
	if object == nil {
		return fmt.Errorf("object not found")
	}
	c.Set("ETag", object.ETag)
	return c.SendFile(path)
}

func buildNewBucket(bucketName string) Bucket {
	return Bucket{
		Name:         bucketName,
		CreationDate: time.Now(),
	}
}

func buildNewObject(key string, file []byte) Object {
	return Object{
		Key:          key,
		CreationDate: time.Now(),
		LastModified: time.Now(),
		Size:         len(file),
		ETag:         utils.CalculateHash(file),
	}
}

func appendOrUpdateObject(objects []Object, newObj Object) []Object {
	for i, o := range objects {
		if o.Key == newObj.Key {
			newObj.CreationDate = o.CreationDate
			objects[i] = newObj
			return objects
		}
	}
	return append(objects, newObj)
}

func filterObjectsByPrefix(objects []Object, prefix string) []Object {
	var result []Object
	for _, obj := range objects {
		if strings.HasPrefix(obj.Key, prefix) {
			result = append(result, obj)
		}
	}
	return result
}

func buildListOutput(bucketName string, maxKeys int, prefix string, marker string, objects []Object) ListObjectsOutput {
	return ListObjectsOutput{
		Name:        bucketName,
		Prefix:      prefix,
		MaxKeys:     maxKeys,
		Marker:      marker,
		Contents:    objects,
		IsTruncated: false,
	}
}

func getObject(bucket Bucket, key string) *Object {
	for _, obj := range bucket.Objects {
		if obj.Key == key {
			return &obj
		}
	}
	return nil
}
