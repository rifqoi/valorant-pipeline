package gcs

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"path"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
)

const BUCKET_NAME = "dtc_data_lake_erudite-bonbon-352111"

type cloudStorage struct {
	player     *ingest.Player
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle
	w          io.Writer
	ctx        context.Context
}

func NewCloudStorage(player *ingest.Player) *cloudStorage {
	ctx := context.Background()
	cloudStorage := &cloudStorage{}
	cloudStorage.bucketName = BUCKET_NAME
	cloudStorage.player = player

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

func (cs *cloudStorage) errorf(format string, args ...interface{}) {
	fmt.Fprintln(cs.w, fmt.Sprintf(format, args...))
	log.Println(cs.ctx, fmt.Sprintf(format, args...))
}

func writeObjectFile(cs *cloudStorage, stringToWrite string, objectPath string) error {
	wc := cs.bucket.Object(objectPath).NewWriter(cs.ctx)
	wc.ContentType = "text/plain"

	if _, err := wc.Write([]byte(stringToWrite)); err != nil {
		cs.errorf("Create Log: unable to create %s to bucket %q, file %q: %v", objectPath, cs.bucketName, objectPath, err)
		return err
	}
	if err := wc.Close(); err != nil {
		// TODO: handle error.
		cs.errorf("Create Log: unable to create %s to bucket %q, file %q: %v", objectPath, cs.bucketName, objectPath, err)
		return err
	}
	log.Println("updated object:", objectPath)
	return nil
}

func readLines(r io.Reader) ([]string, error) {
	var linesSlices []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		linesSlices = append(linesSlices, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return linesSlices, nil
}

func (cs *cloudStorage) logFileExists(objectPath string) error {
	objectPath = path.Join(objectPath, "history.log")
	_, err := cs.bucket.Object(objectPath).Attrs(cs.ctx)
	if err != nil {
		log.Println(objectPath, "does not exists")
		return err
	}
	log.Println(objectPath, "exists")
	return nil
}

func (cs *cloudStorage) createLogFile(playerPath string) (string, error) {
	objectPath := path.Join(playerPath, "history.log")

	if err := writeObjectFile(cs, "", objectPath); err != nil {
		return "", err
	}

	return objectPath, nil
}

func (cs *cloudStorage) readLogFile(objectPath string) (string, error) {
	objectPath = path.Join(objectPath, "history.log")
	rc, err := cs.bucket.Object(objectPath).NewReader(cs.ctx)
	if err != nil {
		return "", err
	}
	defer rc.Close()

	linesSlices, err := readLines(rc)
	if err != nil {
		return "", err
	}

	if len(linesSlices) == 0 {
		fmt.Println("kosong cuy")
		return "", nil
	}
	lastLine := linesSlices[len(linesSlices)-1]

	return lastLine, nil
}

func (cs *cloudStorage) writeJSON(objectPath string, jsonInterface interface{}) error {

	jsonBytes, err := json.MarshalIndent(jsonInterface, "", "   ")
	if err != nil {
		return err
	}
	wc := cs.bucket.Object(objectPath).NewWriter(cs.ctx)
	wc.ContentType = "application/json"

	if _, err := wc.Write(jsonBytes); err != nil {
		cs.errorf("Create Log: unable to create %s to bucket %q, file %q: %v", objectPath, cs.bucketName, objectPath, err)
		return err
	}
	if err := wc.Close(); err != nil {
		// TODO: handle error.
		cs.errorf("Create Log: unable to create %s to bucket %q, file %q: %v", objectPath, cs.bucketName, objectPath, err)
		return err
	}
	log.Println("updated object:", objectPath)
	return nil
}

func (cs *cloudStorage) appendHistory(objectPath string, logInformation string) error {
	rc, err := cs.bucket.Object(objectPath).NewReader(cs.ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	linesSlices, err := readLines(rc)
	if err != nil {
		return err
	}

	linesSlices = append(linesSlices, logInformation)

	logString := strings.Join(linesSlices, "\n")

	if err := writeObjectFile(cs, logString, objectPath); err != nil {
		return err
	}

	return nil
}

func (cs *cloudStorage) WritePlayerJSON() error {
	playerPath := fmt.Sprintf("Player/%s#%s", cs.player.Name, cs.player.Tag)
	jsonName := fmt.Sprintf("%s.json", cs.player.Name)
	objectPath := path.Join(playerPath, jsonName)

	if err := cs.writeJSON(objectPath, cs.player.PlayerData); err != nil {
		return err
	}

	return nil
}

func (cs *cloudStorage) WriteMatchJSON(match *model.Match) error {
	// Last Match ID
	lastMatchID := match.Data[0].Metadata.Matchid
	// Directory Player/<PlayerName>#>Tag>/
	dir := fmt.Sprintf("Player/%s#%s/matches", cs.player.Name, cs.player.Tag)
	// File Player/<PlayerName>#>Tag>/<match-id>.json
	filename := fmt.Sprintf("%s.json", lastMatchID)
	fullPath := path.Join(dir, filename)
	logPath := path.Join(dir, "history.log")

	log.Println("Checking log history")
	if err := cs.logFileExists(dir); err != nil {
		log.Println("Log file not found. Creating a new log file....")
		logPath, err := cs.createLogFile(dir)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Appending match id to log file...")
		if err := cs.appendHistory(logPath, lastMatchID); err != nil {
			log.Fatal(err)
		}

		log.Println("Writing json....")
		if err = cs.writeJSON(fullPath, match); err != nil {
			log.Fatal(err)
		}

		log.Println("Json sucessfully created!")
		return nil
	}

	log.Println("Reading log file....")
	lastHistory, err := cs.readLogFile(dir)
	if err != nil {
		cs.bucket.Object(logPath).Delete(cs.ctx)
		return fmt.Errorf("Empty history")
	}

	if lastHistory == lastMatchID {
		return fmt.Errorf("There was an existing match: %s", lastMatchID)
	}

	log.Println("New match found!")
	log.Printf("Appending %s to %s", lastMatchID, fullPath)
	if err := cs.appendHistory(logPath, lastMatchID); err != nil {
		return err
	}

	log.Printf("Creating %s.json file", lastMatchID)
	if err = cs.writeJSON(fullPath, match); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cs)
	return nil
}
