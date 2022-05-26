package ingest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
)

type PlayerData struct {
	Name         string
	Tag          string
	Region       string
	Puuid        string
	AccountLevel int
}

func NewPlayerData(name string, tag string) *PlayerData {
	d := new(PlayerData)

	d.Name = name
	d.Tag = tag

	player, err := d.getPlayerData()

	if err != nil {
		log.Fatal(err)
	}

	d.Region = player.Data.Region
	d.Puuid = player.Data.Puuid
	d.AccountLevel = player.Data.AccountLevel

	return d
}

func (d *PlayerData) getPlayerData() (*model.Player, error) {
	player := &model.Player{}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/account/%s/%s", d.Name, d.Tag)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&player)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return player, nil
}

func (d *PlayerData) GetMatchData() (*model.Match, error) {
	player := &model.Match{}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v3/matches/%s/%s/%s", d.Region, d.Name, d.Tag)
	fmt.Println("GET: ", url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&player)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return player, nil
}
