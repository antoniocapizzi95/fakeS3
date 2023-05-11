package main

import (
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/config"
	"github.com/antoniocapizzi95/fakeS3/database"
	"github.com/antoniocapizzi95/fakeS3/pkg/s3"
	"github.com/antoniocapizzi95/fakeS3/repository/mongo"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	conf := config.NewConfig()
	db := database.ConnectionMongo(conf)

	app := fiber.New()

	// if I change the database, I only have to change the following row
	bucketHandler := mongo.GetBucketHandlerMongo(db.Collection("bucket"))
	// S3 configuration
	s3Service := s3.New(bucketHandler, conf)

	// Endpoint to create a Bucket
	app.Put("/:bucket", s3Service.CreateBucket)

	// Endpoint to add an object inside a Bucket
	app.Put("/:bucket/+", s3Service.PutObject)

	// Endpoint to list objects inside a Bucket
	app.Get("/:bucket", s3Service.ListObjects)

	// Endpoint to get an object
	app.Get("/:bucket/+", s3Service.GetObject)

	if err := app.Listen(fmt.Sprintf(":%s", conf.Port)); err != nil {
		log.Fatalf("Error during server running: %v\n", err)
	}
}
