package s3

import (
	"context"
	"encoding/xml"
	"time"
)

type Bucket struct {
	Name         string    `bson:"name"`
	CreationDate time.Time `bson:"creation_date"`
	Objects      []Object  `bson:"objects"`
}

type Object struct {
	Key          string    `bson:"key" xml:"Key"`
	ETag         string    `bson:"etag" xml:"ETag"`
	CreationDate time.Time `bson:"creation_date"`
	LastModified time.Time `bson:"last_modified" xml:"LastModified"`
	Size         uint64    `bson:"size" xml:"Size"`
}

type BucketHandler interface {
	GetBucket(ctx context.Context, bucketName string) (*Bucket, error)
	CreateBucket(ctx context.Context, bucket Bucket) error
	UpdateBucket(ctx context.Context, bucket Bucket) error
}

type ListObjectsOutput struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	Contents    []Object `xml:"Contents"`
}
