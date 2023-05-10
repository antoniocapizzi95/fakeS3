package mongo

import (
	"context"
	"fmt"
	"github.com/antoniocapizzi95/fakeS3/pkg/s3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BucketHandlerMongo struct {
	collection *mongo.Collection
}

func (b *BucketHandlerMongo) CreateBucket(ctx context.Context, bucket s3.Bucket) error {
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

func (b *BucketHandlerMongo) AddObject(ctx context.Context, bucketName string, object s3.Object) error {
	bucket, err := b.getBucket(ctx, bucketName)
	if err != nil {
		return err
	}
	if bucket == nil {
		return fmt.Errorf("not found bucket with name %s", bucketName)
	}
	bucket.Objects = append(bucket.Objects, object)
	_, err = b.collection.UpdateOne(ctx, bson.M{"name": bucketName}, bson.M{"objects": bucket.Objects})
	if err != nil {
		return err
	}
	return nil
}

func (b *BucketHandlerMongo) GetObjects(ctx context.Context, bucketName string, max uint, prefix string, marker string) ([]s3.Object, error) {

	return nil, nil
}

func (b *BucketHandlerMongo) getBucket(ctx context.Context, bucketName string) (*s3.Bucket, error) {
	var bucket s3.Bucket
	err := b.collection.FindOne(ctx, bson.M{"name": bucketName}).Decode(&bucket)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func GetBucketHandlerMongo(collection *mongo.Collection) s3.BucketHandler {
	return &BucketHandlerMongo{collection: collection}
}
