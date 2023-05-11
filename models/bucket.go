package models

import (
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
	Size         int       `bson:"size" xml:"Size"`
}

type ListObjectsOutput struct {
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	Contents    []Object `xml:"Contents"`
}
