package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/datalake"
	cs "github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/datalake/gcs"
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

		match, err := dataPlayer.GetMatchData()
		if err != nil {
			log.Println(err)
		}

		storage := chooseStorage("gcs", dataPlayer)
		if err := storage.WritePlayerJSON(); err != nil {
			log.Println(err)
		}
		if err := storage.WriteMatchJSON(match); err != nil {
			log.Println(err)
		}
		fmt.Println()
	}

	_storage := chooseStorage("gcs", nil)
	if err := _storage.CheckStorageForNewFile(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nice")
}

func chooseStorage(storageType string, player *ingest.Player) datalake.StorageHelper {
	switch storageType {
	case "local":
		storage := ls.NewLocalStorage(player)

		return storage
	case "gcs":
		storage := cs.NewCloudStorage(player)
		return storage
	}

	return nil
}
