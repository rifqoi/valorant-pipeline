package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/datalake"
	ls "github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/datalake/local"
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

	// To store the keys in slice in sorted order
	var keys []string
	for k := range players {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		dataPlayer := ingest.NewPlayerData(key, players[key])

		storage := chooseStorage("local", dataPlayer)
		match, err := dataPlayer.GetMatchData()
		if err != nil {
			log.Println(err)
		}
		if err := storage.WriteMatchJSON(match); err != nil {
			log.Println(err)
		}
		fmt.Println("")
	}
	fmt.Println("Nice")
	// dataPlayer := ingest.NewPlayerData("Aledania", "3110")

	// gcs := ingest.NewCloudStorage(dataPlayer)
	// gcs.CreateCloudLogFile("Player/Aledania#3110")

}

func chooseStorage(storageType string, player *ingest.Player) datalake.StorageHelper {
	switch storageType {
	case "local":
		storage := ls.NewLocalStorage(player)

		return storage
	case "gcs":
		return nil
	}

	return nil
}
