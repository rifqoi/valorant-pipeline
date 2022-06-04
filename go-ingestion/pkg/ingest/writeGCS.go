package ingest

import (
	"io"

	"cloud.google.com/go/storage"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
	"golang.org/x/net/context"
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

func cloudLogFileExists(bucket string) error {
	return nil
}

func createCloudLogFile(bucket string) (string, error) {
	return "", nil
}

func readCloudLogFile(objectPath string) (string, error) {
	return "", nil
}

func writeCloudJson(objectPath string, jsonInterface interface{}) error {
	return nil
}

func appendHistoryToCloudLogFile(objectPath string, logInformation string) error {
	return nil
}

func (p *Player) WritePlayerCloudJSON() error {
	return nil
}

func (p *Player) WriteMatchCloudJSON(match *model.Match) error {
	return nil
}
