package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
func readFile(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 62)
	stat, err := os.Stat(fname)
	start := stat.Size() - 62
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("%s\n", buf)
	}

}

func main() {
	dataLino := ingest.NewPlayerData("Aledania", "3110")

	match, err := dataLino.GetMatchData()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := dataLino.WritePlayerJSON(); err != nil {
		log.Println(err)
	}

	if err := dataLino.WriteMatchJSON(match); err != nil {
		log.Println(err)
	}

	dataNanda := ingest.NewPlayerData("iNeChan", "uwu")

	if err := dataNanda.WritePlayerJSON(); err != nil {
		log.Println(err)
	}
	match, err = dataNanda.GetMatchData()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := dataNanda.WriteMatchJSON(match); err != nil {
		log.Println(err)
	}

}
