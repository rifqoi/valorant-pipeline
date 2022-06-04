package main

import (
	"encoding/json"
	"log"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func main() {
	players := make(map[string]string)

	players["Aledania"] = "3110"
	players["iNeChan"] = "uwu"
	players["I U"] = "8400"

	for key, value := range players {
		dataPlayer := ingest.NewPlayerData(key, value)
		match, err := dataPlayer.GetMatchData()
		if err != nil {
			log.Fatal(err)
		}
		if err := dataPlayer.WriteMatchLocalJSON(match); err != nil {
			log.Println(err)
		}
	}
}
