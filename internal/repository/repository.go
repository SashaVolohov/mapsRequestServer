package repository

import (
	"time"

	mapsRequest "github.com/SashaVolohov/mapsRequestServer"
)

type API interface {
	CreateValueByKey(key string, value string, time time.Time) error
	GetValueByKey(key string) (string, error)
	DeleteValueByKey(key string) error
	GetMaps() map[string]mapsRequest.MapObject
}

type Repository struct {
	API
}

func NewRepository() *Repository {
	return &Repository{
		API: NewAPIMaps(),
	}
}
