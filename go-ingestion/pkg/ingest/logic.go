package ingest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetPlayerData(name string, tag string) (*Player, error) {
	player := &Player{}
	client := &http.Client{}
	baseUrl := "https://api.henrikdev.xyz/valorant/v1/account"
	url := fmt.Sprintf("%s/%s/%s", baseUrl, name, tag)
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
		return nil, err
	}

	return player, nil
}
