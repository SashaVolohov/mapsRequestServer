package repository

import (
	"time"
)

type API interface {
	CreateValueByKey(key string, value string, time time.Time) error
	GetValueByKey(key string) (string, error)
	GetLifeTimeByKey(key string) (int64, error)
	DeleteValueByKey(key string) error
	GetKeys() []string
}

type Repository struct {
	API
}

func NewRepository() *Repository {
	return &Repository{
		API: NewAPIMaps(),
	}
}
