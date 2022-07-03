package datalake

import "github.com/rifqoi/valorant-pipeline/go-ingestion/pkg/ingest/model"

type storageType struct {
	storage Storage
}

func (s *storageType) WritePlayerJSON() error {
	return s.storage.WritePlayerJSON()
}

func (s *storageType) WriteMatchJSON(match *model.Match) error {
	return s.storage.WriteMatchJSON(match)
}

func (s *storageType) CheckStorageForNewFile() error {
	return s.storage.CheckStorageForNewFile()
}

func NewStorage(storage Storage) StorageHelper {
	return &storageType{
		storage,
	}
}
