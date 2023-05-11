package repository

import (
	"context"
	"github.com/antoniocapizzi95/fakeS3/pkg/s3"
)

type BucketRepository interface {
	GetBucket(ctx context.Context, bucketName string) (*s3.Bucket, error)
	CreateBucket(ctx context.Context, bucket s3.Bucket) error
	UpdateBucket(ctx context.Context, bucket s3.Bucket) error
}
