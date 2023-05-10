package s3

import "time"

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
	Size         int64     `bson:"size"`
}

type BucketHandler interface {
	CreateBucket(bucket Bucket) error
	AddObject(bucketName string, object Object) error
	GetObjects(bucketName string, max uint, prefix string, marker string)
}
