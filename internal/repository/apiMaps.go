package repository

import (
	"fmt"
	"time"

	mapsRequest "github.com/SashaVolohov/mapsRequestServer"
)

type APIMaps struct {
	mapItems map[string]mapsRequest.MapObject
}

func NewAPIMaps() *APIMaps {
	return &APIMaps{
		mapItems: make(map[string]mapsRequest.MapObject),
	}
}

func (r *APIMaps) CreateValueByKey(key string, value string, time time.Time) error {

	if _, ok := r.mapItems[key]; ok {
		return fmt.Errorf("this key is already taken by a value")
	}

	r.mapItems[key] = mapsRequest.NewMapObject(value, time)

	return nil
}

func (r *APIMaps) GetValueByKey(key string) (string, error) {
	mapObject, ok := r.mapItems[key]

	if !ok {
		return "", fmt.Errorf("this key is not in the map")
	}

	return mapObject.GetValue(), nil
}

func (r *APIMaps) GetLifeTimeByKey(key string) (int64, error) {
	mapObject, ok := r.mapItems[key]

	if !ok {
		return 0, fmt.Errorf("this key is not in the map")
	}

	return mapObject.GetLifeTime(), nil
}

func (r *APIMaps) DeleteValueByKey(key string) error {

	if _, ok := r.mapItems[key]; !ok {
		return fmt.Errorf("this key is not in the map")
	}

	delete(r.mapItems, key)
	return nil
}

func (r *APIMaps) GetKeys() []string {

	keys := make([]string, 0, len(r.mapItems))
	for k := range r.mapItems {
		keys = append(keys, k)
	}

	return keys
}
