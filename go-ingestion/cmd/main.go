package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func retryOnTimeout(player *ingest.Player) *ingest.Match {
	matchInterface := &ingest.Match{}
	if err := retry.Do(
		func() error {
			match, err := player.GetMatchData()
			if err != nil {
				log.Println(err)
				return err
			}
			matchInterface = match
			return nil
		},
		retry.Delay(2*time.Minute+40*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	return matchInterface

}

func main() {
	players := make(map[string]string)

	players["Aledania"] = "3110"
	players["iNeChan"] = "uwu"
	players["I U"] = "8400"

	for key, value := range players {
		dataPlayer := ingest.NewPlayerData(key, value)
		match := retryOnTimeout(dataPlayer)
		if err := dataPlayer.WriteMatchLocalJSON(match); err != nil {
			log.Println(err)
		}
	}
	fmt.Println("Nice")
}
