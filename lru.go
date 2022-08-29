// A simple implementation of LRU for caching with old data preempted.
package lru

// --------------------------------------------------------
// Copyright (c) 2022 Constantine Zavezeon <kwynto@mail.ru>
// --------------------------------------------------------

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// The internal structure of the cache store.
type dataCache struct {
	queue  int64
	result any
}

// Contract for the use of the cache.
type Cache interface {
	// Get an entry from the cache.
	Load(key any) (any, error)

	// Write a new value to the cache or update an old value.
	Store(key any, value any) bool
}

// Cache infrastructure.
type cache struct {
	data     map[string]dataCache
	latch    sync.Mutex
	capacity int
	border   int
	cleaning bool
}

// Cache constructor.
func New(size int) Cache {
	if size < 100 {
		size = 100
	}

	return &cache{
		data:     make(map[string]dataCache, 0),
		latch:    sync.Mutex{},
		capacity: size,
		border:   size / 100 * 98,
		cleaning: false,
	}
}

// Get an entry from the cache.
func (c *cache) Load(key any) (any, error) {
	keyStr, err := marshalKey(key)
	if err != nil {
		return nil, errors.New("can't get result from cache")
	}

	newTime := time.Now().UnixNano()

	c.latch.Lock()
	defer c.latch.Unlock()

	value, ok := c.data[keyStr]
	if !ok {
		return nil, errors.New("can't get result from cache")
	}

	value.queue = newTime
	c.data[keyStr] = value

	return value.result, nil
}

// Write a new value to the cache or update an old value.
func (c *cache) Store(key any, value any) bool {
	var valueStore dataCache

	c.balancing()

	keyStr, err := marshalKey(key)
	if err != nil {
		return false
	}

	valueStore.queue = time.Now().UnixNano()
	valueStore.result = value

	c.latch.Lock()
	c.data[keyStr] = valueStore
	c.latch.Unlock()

	return true
}

// Helper functions and metods

// Retrieves the key of the oldest cache item
func (c *cache) extractMinValue() string {
	markerT := time.Now().UnixNano()
	markerK := ""

	// c.latch.Lock()
	for key, value := range c.data {
		if value.queue < markerT {
			markerT = value.queue
			markerK = key
		}
	}
	// c.latch.Unlock()

	return markerK
}

// Removing a cache element by key.
func (c *cache) remove(key string) {
	c.latch.Lock()
	delete(c.data, key)
	c.latch.Unlock()
}

// Clearing the cache of old items
func (c *cache) cleanUp(count int8) {
	if count < 2 {
		minKey := c.extractMinValue()
		c.remove(minKey)
		c.cleanUp(count + 1)
	}
}

// Cache balancing on overflow reducing it by 2%
func (c *cache) balancing() {
	lenData := len(c.data)

	if !c.cleaning && lenData >= c.capacity {
		c.cleaning = true
		c.cleanUp(0)
	} else if c.cleaning && lenData > c.border {
		c.cleanUp(0)
	}

	if c.cleaning && lenData <= c.border {
		c.cleaning = false
	}
}

// Serialize any type to JSON to use as a key.
func marshalKey(in any) (string, error) {
	out, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
