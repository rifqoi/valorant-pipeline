package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func main() {
	playerName := "Aledania"
	playerTag := "3110"

	req, err := ingest.GetPlayerData(playerName, playerTag)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	res, err := PrettyStruct(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
