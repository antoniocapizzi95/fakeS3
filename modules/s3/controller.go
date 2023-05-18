package s3

import (
	"encoding/xml"
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/config"
	"github.com/antoniocapizzi95/fakeS3/models"
	"github.com/antoniocapizzi95/fakeS3/repository"
	"github.com/antoniocapizzi95/fakeS3/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

type S3Controller interface {
	CreateBucket(c *fiber.Ctx) error
	PutObject(c *fiber.Ctx) error
	ListObjects(c *fiber.Ctx) error
	GetObject(c *fiber.Ctx) error
}

// s3Controller is the implementation of s3Controller interface
type s3Controller struct {
	bucketRepository repository.BucketRepository
	conf             config.Config
}

// New this return a new s3Controller object
func New(bucketRepository repository.BucketRepository, conf config.Config) S3Controller {
	return &s3Controller{
		bucketRepository: bucketRepository,
		conf:             conf,
	}
}

func (s *s3Controller) CreateBucket(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	bucket := buildNewBucket(bucketName)
	err := s.bucketRepository.CreateBucket(c.Context(), bucket)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).Type("application/xml").SendString("")
}

func (s *s3Controller) PutObject(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	key := c.Params("+")
	ctx := c.Context()
	file := c.Context().RequestBodyStream()
	bodyLength := c.Request().Header.ContentLength()

	etag, err := utils.WriteFile(s.conf.StoragePath, bucketName, key, file)
	if err != nil {
		return err
	}
	object := buildNewObject(key, etag, bodyLength)
	bucket, err := s.bucketRepository.GetBucket(ctx, bucketName)
	if err != nil {
		return err
	}
	if bucket == nil {
		return fmt.Errorf("bucket with name %s not found", bucketName)
	}
	bucket.Objects = appendOrUpdateObject(bucket.Objects, object)
	err = s.bucketRepository.UpdateBucket(ctx, *bucket)
	if err != nil {
		return err
	}
	c.Set("ETag", object.ETag)
	return c.Status(fiber.StatusOK).Type("application/xml").SendString("")
}

func (s *s3Controller) ListObjects(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	maxKeys, _ := strconv.Atoi(c.Query("max-keys"))
	prefix := c.Query("prefix")
	marker := c.Query("marker")
	bucket, err := s.bucketRepository.GetBucket(c.Context(), bucketName)
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
	output := buildListOutput(bucketName, maxKeys, prefix, marker, objects)
	xmlBytes, err := xml.Marshal(output)
	if err != nil {
		return fmt.Errorf("error marshaling to XML: %v", err)
	}
	xmlString := string(xmlBytes)
	return c.Status(fiber.StatusOK).Type("application/xml").SendString(xmlString)
}

func (s *s3Controller) GetObject(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	rangeParam := c.Get("Range")
	key := c.Params("+")

	path := utils.GetEntirePath(s.conf.StoragePath, bucketName, key)
	bucket, err := s.bucketRepository.GetBucket(c.Context(), bucketName)
	if err != nil {
		return err
	}

	object := getObject(*bucket, key)
	if object == nil {
		return fmt.Errorf("object not found")
	}
	c.Set("ETag", object.ETag)

	// handle range
	if rangeParam != "" {
		start, end, err := utils.ExtractByteRanges(rangeParam)
		if err != nil {
			return err
		}
		partialFile, err := utils.ReadFileRange(path, start, end)
		if err != nil {
			return err
		}
		return c.Send(partialFile)
	}

	return c.SendFile(path)
}

func buildNewBucket(bucketName string) models.Bucket {
	return models.Bucket{
		Name:         bucketName,
		CreationDate: time.Now(),
	}
}

func buildNewObject(key string, etag string, bodyLength int) models.Object {
	return models.Object{
		Key:          key,
		CreationDate: time.Now(),
		LastModified: time.Now(),
		Size:         bodyLength,
		ETag:         etag,
	}
}

func appendOrUpdateObject(objects []models.Object, newObj models.Object) []models.Object {
	for i, o := range objects {
		if o.Key == newObj.Key {
			newObj.CreationDate = o.CreationDate
			objects[i] = newObj
			return objects
		}
	}
	return append(objects, newObj)
}

func filterObjectsByPrefix(objects []models.Object, prefix string) []models.Object {
	var result []models.Object
	for _, obj := range objects {
		if strings.HasPrefix(obj.Key, prefix) {
			result = append(result, obj)
		}
	}
	return result
}

func buildListOutput(bucketName string, maxKeys int, prefix string, marker string, objects []models.Object) models.ListObjectsOutput {
	return models.ListObjectsOutput{
		Name:        bucketName,
		Prefix:      prefix,
		MaxKeys:     maxKeys,
		Marker:      marker,
		Contents:    objects,
		IsTruncated: false,
	}
}

func getObject(bucket models.Bucket, key string) *models.Object {
	for _, obj := range bucket.Objects {
		if obj.Key == key {
			return &obj
		}
	}
	return nil
}
