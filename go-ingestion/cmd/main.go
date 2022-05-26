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
	dataLino := ingest.NewPlayerData("Aledania", "3110")

	req, err := dataLino.GetMatchData()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	res, err := PrettyStruct(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
