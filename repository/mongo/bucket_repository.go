package mongo

import (
	"context"
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/pkg/s3"
	"github.com/antoniocapizzi95/fakeS3/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BucketRepositoryMongo struct {
	collection *mongo.Collection
}

func (b *BucketRepositoryMongo) CreateBucket(ctx context.Context, bucket s3.Bucket) error {
	// add if bucket don't exists (using name as unique identifier)
	res, err := b.collection.UpdateOne(
		ctx,
		bson.M{"name": bucket.Name}, bson.M{"$setOnInsert": bucket},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	if res.UpsertedCount == 0 {
		return fmt.Errorf("this bucket already exists")
	}
	return nil
}

func (b *BucketRepositoryMongo) UpdateBucket(ctx context.Context, bucket s3.Bucket) error {
	result, err := b.collection.UpdateOne(ctx, bson.M{"name": bucket.Name}, bson.M{"$set": bucket})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("the bucket %s doesn't exists", bucket.Name)
	}
	return nil
}

func (b *BucketRepositoryMongo) GetBucket(ctx context.Context, bucketName string) (*s3.Bucket, error) {
	var bucket s3.Bucket
	err := b.collection.FindOne(ctx, bson.M{"name": bucketName}).Decode(&bucket)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func GetBucketRepositoryMongo(db *mongo.Database) repository.BucketRepository {
	return &BucketRepositoryMongo{collection: db.Collection("bucket")}
}
