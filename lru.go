package lru

import (
	"encoding/json"
	"errors"
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
	Get(key any) (any, error)
	Set(key any, value any) bool
	СheckUp(key any) bool
}

type cache struct {
	data   sync.Map
	volume int
}

func New(size int) Cache {
	return &cache{
		data:   sync.Map{},
		volume: size,
	}
}

func (c *cache) Get(key any) (any, error) {
	keyStr, ok := MarshalKey(key)
	if !ok {
		return nil, errors.New("can't get result from cache")
	}

	value, ok := c.data.Load(keyStr)
	if !ok {
		return nil, errors.New("can't get result from cache")
	}

	return value, nil
}

func (c *cache) Set(key any, value any) bool {
	keyStr, ok := MarshalKey(key)
	if !ok {
		return false
	}
	c.data.Store(keyStr, value)
	return true
}

func (c *cache) СheckUp(key any) bool {
	keyStr, ok := MarshalKey(key)
	if !ok {
		return false
	}

	_, ok = c.data.Load(keyStr)
	return ok
}

// Helper functions

func MarshalKey(in any) (string, bool) {
	out, err := json.Marshal(in)
	if err != nil {
		return "", false
	}
	return string(out), true
}
