package mapsRequest

import (
	"sync"
	"time"
)

type MapObject struct {
	value    string
	lifeTime time.Time
	mu       *sync.Mutex
}

func NewMapObject(value string, lifeTime time.Time) MapObject {
	return MapObject{value: value, lifeTime: lifeTime, mu: &sync.Mutex{}}
}

func (m *MapObject) GetValue() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.value
}

func (m *MapObject) GetLifeTime() int64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.lifeTime.Unix()
}
