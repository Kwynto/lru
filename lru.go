package lru

import (
	"encoding/json"
	"sync"
	"time"
)

// dataCache

type dataCache struct {
	result any
	queue  time.Time
}

// Cache

type Cache interface {
	Process(in any) (any, error)
}

type cache struct {
	data   sync.Map
	volume int
}

func New(size int) Cache {
	return cache{
		data:   sync.Map{},
		volume: size,
	}
}

func (c cache) Process(in any) (any, error) {
	return nil, nil
}

// Helper functions

func MarshalKey(in any) (string, bool) {
	out, err := json.Marshal(in)
	if err != nil {
		return "", false
	}
	return string(out), true
}
