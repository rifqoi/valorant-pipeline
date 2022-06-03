package ingest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

type Player struct {
	Name       string
	Tag        string
	Region     string
	PlayerData *PlayerData
}

func NewPlayerData(name string, tag string) *Player {
	player := &Player{}

	player.Name = name
	player.Tag = tag

	playerData := &PlayerData{}
	if err := isPlayerDataExists(name, tag); err != nil {
		log.Println("PlayerData not found. Fetching data from API.")
		fetchData, err := getPlayerData(name, tag)
		if err != nil {
			log.Fatal(err)
		}
		playerData = fetchData
		player.PlayerData = playerData
		if err := player.WritePlayerLocalJSON(); err != nil {
			log.Fatal(err)
		}
		fmt.Print("asd")
	} else {
		log.Println("PlayerData found!")
		dir := fmt.Sprintf("Player/%s#%s", name, tag)
		jsonBytes, err := os.ReadFile(path.Join(dir, name+".json"))
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(jsonBytes, playerData)
		if err != nil {
			log.Fatal(err)
		}
		player.PlayerData = playerData
	}
	player.Region = playerData.Data.Region
	return player
}

func isPlayerDataExists(name string, tag string) error {
	dir := path.Join("Player", name+"#"+tag)
	path := path.Join(dir, name+".json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println(path + " is not exist")
		return err
	}
	log.Println(path + " is exist")
	return nil
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
	log.Println("GET: ", url)

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
