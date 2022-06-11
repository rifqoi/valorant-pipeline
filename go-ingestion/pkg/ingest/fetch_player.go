package ingest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/avast/retry-go"
	"github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"
)

type Player struct {
	Name       string
	Tag        string
	Region     string
	PlayerData *model.PlayerData
}

func NewPlayerData(name string, tag string) *Player {
	player := &Player{}

	player.Name = name
	player.Tag = tag

	log.Printf("Fetching %s#%s player data from API. ", name, tag)
	fetchData, err := getPlayerData(name, tag)
	if err != nil {
		log.Fatal(err)
	}
	player.PlayerData = fetchData
	// if err := isPlayerDataExists(name, tag); err != nil {
	// 	log.Println("PlayerData not found. Fetching data from API.")
	// 	fetchData, err := getPlayerData(name, tag)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	playerData = fetchData
	// 	player.PlayerData = playerData
	// 	if err := player.WritePlayerLocalJSON(); err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Print("asd")
	// } else {
	// 	log.Println("PlayerData found!")
	// 	dir := fmt.Sprintf("Player/%s#%s", name, tag)
	// 	jsonBytes, err := os.ReadFile(path.Join(dir, name+".json"))
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	err = json.Unmarshal(jsonBytes, playerData)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	player.PlayerData = playerData
	// }
	player.Region = fetchData.Data.Region
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

func getJson(url string) []byte {
	var body []byte
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return err
			}

			if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
				log.Printf("Response Status: %v", resp.StatusCode)
				return err
			}

			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)

			return nil
		},
		retry.Delay(3*time.Minute),
	)

	if err != nil {
		log.Println(err)
	}

	return body
}

func getPlayerData(name string, tag string) (*model.PlayerData, error) {
	player := &model.PlayerData{}
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/account/%s/%s", name, tag)

	jsonBytes := getJson(url)

	err := json.Unmarshal(jsonBytes, &player)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return player, nil
}
