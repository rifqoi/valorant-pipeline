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

	if err := dataLino.WriteMatchLocalJSON(match); err != nil {
		log.Println(err)
	}

	dataNanda := ingest.NewPlayerData("iNeChan", "uwu")

	match, err = dataNanda.GetMatchData()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := dataNanda.WriteMatchLocalJSON(match); err != nil {
		log.Println(err)
	}

	dataIU := ingest.NewPlayerData("I U", "8400")

	match, err = dataIU.GetMatchData()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := dataIU.WriteMatchLocalJSON(match); err != nil {
		log.Println(err)
	}
}
