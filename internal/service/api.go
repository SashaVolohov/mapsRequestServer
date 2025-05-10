package service

import (
	"context"
	"fmt"
	"time"

	"github.com/SashaVolohov/mapsRequestServer/internal/repository"
	"github.com/sirupsen/logrus"
)

type APIService struct {
	repo repository.API
}

func NewAPIService(repo repository.API) *APIService {
	return &APIService{repo: repo}
}

func (s *APIService) CreateValueByKey(key string, value string, time time.Time) error {
	return s.repo.CreateValueByKey(key, value, time)
}

func (s *APIService) GetValueByKey(key string) (string, error) {
	return s.repo.GetValueByKey(key)
}

func (s *APIService) DeleteValueByKey(key string) error {
	return s.repo.DeleteValueByKey(key)
}

func (s *APIService) KeyCollector(ctx context.Context) {

	for {
		select {

		case <-ctx.Done():
			return

		case <-time.After(time.Second):

			for _, key := range s.repo.GetKeys() {

				lifeTime, err := s.repo.GetLifeTimeByKey(key)
				if err != nil {
					logrus.Errorf(fmt.Sprintf("Unable to get key lifetime - %s", err.Error()))
				}

				if time.Now().Unix() >= lifeTime {
					err := s.repo.DeleteValueByKey(key)

					if err != nil {
						logrus.Errorf(fmt.Sprintf("Unable to delete old map key - %s", err.Error()))
					}

				}

			}

		}
	}
}
