package service

import (
	"context"
	"time"

	"github.com/SashaVolohov/mapsRequestServer/internal/repository"
)

type API interface {
	CreateValueByKey(key string, value string, time time.Time) error
	GetValueByKey(key string) (string, error)
	DeleteValueByKey(key string) error
	KeyCollector(ctx context.Context)
}

type Service struct {
	API
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		API: NewAPIService(repos.API),
	}
}
