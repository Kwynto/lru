package lru

import (
	"time"
)

// DataTable

type DataTable interface {
	Add(key string, value interface{}) bool
}

type dataTable map[string]interface{}

func NewDataTable(size int) DataTable {
	return make(dataTable, size)
}

func (dt dataTable) Add(key string, value interface{}) bool {
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
	WithCache(key string, value interface{}) bool
}

type cache struct {
	data  DataTable
	queue TimeQueue
}

func New(size int) Cache {
	return cache{
		data:  NewDataTable(size),
		queue: NewTimeQueue(size),
	}
}

func (c cache) WithCache(key string, value interface{}) bool {
	return true
}

func init() {
	//
}
