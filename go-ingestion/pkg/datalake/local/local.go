package local

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
)

type localStorage struct {
	player *ingest.Player
}

func NewLocalStorage(player *ingest.Player) *localStorage {
	return &localStorage{
		player: player,
	}
}

func logFileExists(dir string) error {
	logPath := path.Join(dir, "history.log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		return err
	}
	return nil

}

func createLogFile(dir string) (string, error) {
	logPath := path.Join(dir, "history.log")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	os.Create(logPath)
	return logPath, nil

}

func (ls *localStorage) readLogFile(filepath string) (string, error) {
	f, err := os.Open(path.Join(filepath, "history.log"))
	if err != nil {
		return "", err
	}
	defer f.Close()

	var linesSlices []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		linesSlices = append(linesSlices, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(linesSlices) == 0 {
		return "", nil
	}
	lastLine := linesSlices[len(linesSlices)-1]

	return lastLine, nil

}

func writeJSON(filepath string, jsonInterface interface{}) error {
	jsonFile, err := json.MarshalIndent(jsonInterface, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath, jsonFile, 0755)
	if err != nil {
		return err
	}
	return nil

}

func appendHistory(filepath string, logInformation string) error {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Println(fmt.Sprintf("Appending %s to %s", logInformation, filepath))
	if _, err = f.WriteString(logInformation + "\n"); err != nil {
		return err
	}

	return nil

}

func (ls *localStorage) WritePlayerJSON() error {
	dir := fmt.Sprintf("Player/%s#%s", ls.player.Name, ls.player.Tag)
	filename := fmt.Sprintf("%s.json", ls.player.Name)
	fullPath := path.Join(dir, filename)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	log.Printf("Creating %s", fullPath)
	if err := writeJSON(fullPath, ls.player.PlayerData); err != nil {
		return err
	}
	return nil

}

func (ls *localStorage) WriteMatchJSON(match *model.Match) error {
	// Last Match ID
	lastMatchID := match.Data[0].Metadata.Matchid
	// Directory Player/<PlayerName>#>Tag>/
	dir := fmt.Sprintf("Player/%s#%s/matches", ls.player.Name, ls.player.Tag)
	// File Player/<PlayerName>#>Tag>/<match-id>.json
	filename := fmt.Sprintf("%s.json", lastMatchID)
	fullPath := path.Join(dir, filename)
	logPath := path.Join(dir, "history.log")

	// Check whether the last fetched data is the same or not
	log.Println("Checking log history")
	if err := logFileExists(dir); err != nil {
		log.Println("Log file not found. Creating a new log file....")
		logPath, err := createLogFile(dir)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Appending match id to log file...")
		if err := appendHistory(logPath, lastMatchID); err != nil {
			log.Fatal(err)
		}

		log.Println("Writing json....")
		if err = writeJSON(fullPath, match); err != nil {
			log.Fatal(err)
		}

		log.Println("Json sucessfully created!")
		return nil
	}

	log.Println("Reading log file....")
	lastHistory, err := ls.readLogFile(dir)
	if err != nil {
		os.Remove(logPath)
		return fmt.Errorf("Empty history")
	}

	if lastHistory == lastMatchID {
		return fmt.Errorf("There was an existing match: %s", lastMatchID)
	}

	log.Println("New match found!")
	log.Printf("Appending %s to %s", lastMatchID, fullPath)
	if err := appendHistory(logPath, lastMatchID); err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating %s.json file", lastMatchID)
	if err = writeJSON(fullPath, match); err != nil {
		log.Fatal(err)
	}

	return nil
}
