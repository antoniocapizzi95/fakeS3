package database

import (
	"context"
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionMongo(config config.Config) *mongo.Database {
	connString := fmt.Sprintf("mongodb://%s:%s", config.MongoDBHost, config.MongoDBPort)
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database(config.MongoDBDatabase)
}
