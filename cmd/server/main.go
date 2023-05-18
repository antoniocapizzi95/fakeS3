package main

import (
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/config"
	"github.com/antoniocapizzi95/fakeS3/database"
	"github.com/antoniocapizzi95/fakeS3/modules/s3"
	"github.com/antoniocapizzi95/fakeS3/repository/mongo"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	conf := config.NewConfig()
	db := database.ConnectionMongo(conf)

	app := fiber.New()

	// if I change the database, I only have to change the following row
	bucketRepository := mongo.NewBucketRepositoryMongo(db)
	// S3 configuration
	s3Controller := s3.New(bucketRepository, conf)

	// Endpoint to create a Bucket
	app.Put("/:bucket", s3Controller.CreateBucket)

	// Endpoint to add an object inside a Bucket
	app.Put("/:bucket/+", s3Controller.PutObject)

	// Endpoint to list objects inside a Bucket
	app.Get("/:bucket", s3Controller.ListObjects)

	// Endpoint to get an object
	app.Get("/:bucket/+", s3Controller.GetObject)

	app.Server().StreamRequestBody = true

	if err := app.Listen(fmt.Sprintf(":%s", conf.Port)); err != nil {
		log.Fatalf("Error during server running: %v\n", err)
	}
}
