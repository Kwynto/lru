package lru

import (
	"encoding/json"
	"time"
)

// DataTable

type DataTable interface {
	Add(key string, value any) bool
}

type dataTable map[string]any

func NewDataTable(size int) DataTable {
	return make(dataTable, size)
}

func (dt dataTable) Add(key string, value any) bool {
	dt[key] = value
	return true
}

// TimeQueue

type TimeQueue interface {
	Add(key string, qtime time.Time) bool
}

type timeQueue map[string]time.Time

func NewTimeQueue(size int) TimeQueue {
	return make(timeQueue, size)
}

func (tq timeQueue) Add(key string, qtime time.Time) bool {
	tq[key] = qtime
	return true
}

// Cache

type Cache interface {
	WithCache(key string, value any) bool
}

type cache struct {
	data   DataTable
	queue  TimeQueue
	volume int
}

func New(size int) Cache {
	return cache{
		data:   NewDataTable(size),
		queue:  NewTimeQueue(size),
		volume: size,
	}
}

func (c cache) WithCache(key string, value any) bool {
	return true
}

// Helper functions

func MarshalKey(in any) (string, bool) {
	out, err := json.Marshal(in)
	if err != nil {
		return "", false
	}
	return string(out), true
}
