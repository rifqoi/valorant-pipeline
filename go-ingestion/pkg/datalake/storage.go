package datalake

import "github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"

type StorageHelper interface {
	WritePlayerJSON() error
	WriteMatchJSON(match *model.Match) error
}

type Storage interface {
	WritePlayerJSON() error
	WriteMatchJSON(match *model.Match) error
}
