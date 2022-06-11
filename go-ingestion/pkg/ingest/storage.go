package ingest

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
)

type CloudStorage struct {
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle
	w          io.Writer
	ctx        context.Context
	// cleanUp is a list of filenames that need cleaning up at the end of the demo.
	cleanUp []string
	// failed indicates that one or more of the demo steps failed.
	failed bool
}

type Storage struct {
	Player       *Player
	CloudStorage *CloudStorage
}

type StorageUtils interface {
}
