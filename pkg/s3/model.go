package s3

import (
	"context"
	"time"
)

type Bucket struct {
	Name         string    `bson:"name"`
	CreationDate time.Time `bson:"creation_date"`
	Objects      []Object  `bson:"objects"`
}

type Object struct {
	Key          string    `bson:"key"`
	ETag         string    `bson:"etag"`
	CreationDate time.Time `bson:"creation_date"`
	LastModified time.Time `bson:"last_modified"`
	Size         uint64    `bson:"size"`
}

type BucketHandler interface {
	CreateBucket(ctx context.Context, bucket Bucket) error
	AddObject(ctx context.Context, bucketName string, object Object) error
	GetObjects(ctx context.Context, bucketName string, max uint, prefix string, marker string) ([]Object, error)
}
