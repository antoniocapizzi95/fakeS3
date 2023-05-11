package repository

import (
	"context"
	"github.com/antoniocapizzi95/fakeS3/models"
)

type BucketRepository interface {
	GetBucket(ctx context.Context, bucketName string) (*models.Bucket, error)
	CreateBucket(ctx context.Context, bucket models.Bucket) error
	UpdateBucket(ctx context.Context, bucket models.Bucket) error
}
