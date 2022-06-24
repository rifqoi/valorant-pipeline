package ingest

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/avast/retry-go"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
)

func (p *Player) GetMatchData() (*model.Match, error) {
	match := &model.Match{}
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v3/matches/%s/%s/%s", p.Region, p.Name, p.Tag)
	log.Println("GET: ", url)

	err := retry.Do(
		func() error {
			jsonBytes := getJson(url)
			err := json.Unmarshal(jsonBytes, &match)
			if err != nil {
				return err
			}
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return match, nil
}
