package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
  createdAt time.Time
  val []byte
}
type Cache struct {
  dict map[string]cacheEntry
  mu *sync.RWMutex
}

func NewCache() Cache {
  return Cache{
    dict: map[string]cacheEntry{},
    mu: &sync.RWMutex{},
  }
}

func (c *Cache) Add() {}
func (c *Cache) Get() {}