package ingest

import (
	"context"
	"fmt"
	"log"
	"path"

	"cloud.google.com/go/storage"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
)

// type CloudStorage struct {
// 	player     *Player
// 	client     *storage.Client
// 	bucketName string
// 	bucket     *storage.BucketHandle
// 	w          io.Writer
// 	ctx        context.Context
// 	// cleanUp is a list of filenames that need cleaning up at the end of the demo.
// 	cleanUp []string
// 	// failed indicates that one or more of the demo steps failed.
// 	failed bool
// }

const BUCKET_NAME = "dtc_data_lake_erudite-bonbon-352111"

func NewCloudStorage(player *Player) *CloudStorage {
	ctx := context.Background()
	cloudStorage := &CloudStorage{}
	cloudStorage.bucketName = BUCKET_NAME

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}

	cloudStorage.client = client
	cloudStorage.bucket = client.Bucket(BUCKET_NAME)
	cloudStorage.ctx = ctx

	return cloudStorage
}

func (cs *CloudStorage) errorf(format string, args ...interface{}) {
	cs.failed = true
	fmt.Fprintln(cs.w, fmt.Sprintf(format, args...))
	log.Println(cs.ctx, fmt.Sprintf(format, args...))
}

func (cs *CloudStorage) cloudLogFileExists(bucket string) error {
	return nil
}

func (cs *CloudStorage) CreateCloudLogFile(playerPath string) (string, error) {
	objectPath := path.Join(playerPath, "history.log")
	wc := cs.bucket.Object(objectPath).NewWriter(cs.ctx)

	wc.ContentType = "text/plain"

	if _, err := wc.Write([]byte("")); err != nil {
		cs.errorf("Create Log: unable to create history.log to bucket %q, file %q: %v", cs.bucketName, objectPath, err)
		return "", err
	}
	if err := wc.Close(); err != nil {
		// TODO: handle error.
		cs.errorf("Create Log: unable to create history.log to bucket %q, file %q: %v", cs.bucketName, objectPath, err)
		return "", err
	}
	fmt.Println("updated object:", wc.Attrs())

	return objectPath, nil
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
