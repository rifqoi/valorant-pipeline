package model

type PlayerCard struct {
	Small string `json:"small"`
	Large string `json:"large"`
	Wide  string `json:"wide"`
	Id    string `json:"id"`
}
type PlayerStats struct {
	Puuid        string     `json:"puuid"`
	Region       string     `json:"region"`
	AccountLevel int        `json:"account_level"`
	Name         string     `json:"name"`
	Tag          string     `json:"tag"`
	Card         PlayerCard `json:"card"`
	LastUpdate   string     `json:"last_update"`
}

type Player struct {
	Status int         `json:"status"`
	Data   PlayerStats `json:"data"`
}
