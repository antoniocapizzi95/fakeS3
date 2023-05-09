package main

import (
	"github.com/antoniocapizzi95/fakeS3/pkg/s3"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	// S3 configuration
	s3Service := s3.New()

	// Endpoint to create a Bucket
	app.Put("/:bucket", s3Service.CreateBucket)

	// Endpoint to add an object inside a Bucket
	app.Put("/:bucket/:key", s3Service.PutObject)

	// Endpoint to list objects inside a Bucket
	app.Get("/:bucket", s3Service.ListObjects)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error during server running: %v\n", err)
	}
}
