package cache

import (
	"sync"
)

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		store:    map[string]interface{}{},
		capacity: capacity,
	}
}

type LRUCache struct {
	store    map[string]interface{}
	capacity int
	rwLock   sync.RWMutex
}

func (lru *LRUCache) Get(key string) interface{} {
	return lru.store[key]
}

func (lru *LRUCache) Set(key string, value interface{}) {
	lru.store[key] = value
}
