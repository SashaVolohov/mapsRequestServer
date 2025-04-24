package repository

import (
	"fmt"
	"time"

	"maps"

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

	r.mapItems[key] = mapsRequest.MapObject{
		Value:    value,
		LifeTime: time,
	}

	return nil
}

func (r *APIMaps) GetValueByKey(key string) (string, error) {
	mapObject, ok := r.mapItems[key]

	if !ok {
		return "", fmt.Errorf("this key is not in the map")
	}

	return mapObject.Value, nil
}

func (r *APIMaps) DeleteValueByKey(key string) error {

	if _, ok := r.mapItems[key]; !ok {
		return fmt.Errorf("this key is not in the map")
	}

	delete(r.mapItems, key)
	return nil
}

func (r *APIMaps) GetMaps() map[string]mapsRequest.MapObject {
	copiedMap := make(map[string]mapsRequest.MapObject)
	maps.Copy(copiedMap, r.mapItems)
	return copiedMap
}
