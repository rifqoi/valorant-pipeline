package ingest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Player struct {
	Name       string
	Tag        string
	Region     string
	PlayerData *PlayerData
}

func NewPlayerData(name string, tag string) *Player {
	player := &Player{}

	playerData, err := getPlayerData(name, tag)
	if err != nil {
		log.Fatal(err)
	}
	player.Name = name
	player.Tag = tag
	player.Region = playerData.Data.Region
	player.PlayerData = playerData
	return player
}

func getPlayerData(name string, tag string) (*PlayerData, error) {
	player := &PlayerData{}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/account/%s/%s", name, tag)

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

func (p *Player) GetMatchData() (*Match, error) {
	match := &Match{}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v3/matches/%s/%s/%s", p.Region, p.Name, p.Tag)
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

	err = json.NewDecoder(response.Body).Decode(&match)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return match, nil
}
