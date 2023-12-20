package pokecache

import (
	"sync"
	"time"

	"github.com/samvimes01/go-bootdev-pokedexcli/internal/pokeapi"
)

type cacheEntry struct {
	createdAt time.Time
	val       pokeapi.LocationAreaResp
}
type Cache struct {
	dict map[int]cacheEntry
	mu   *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		dict: make(map[int]cacheEntry),
		mu:   &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timePassed := time.Now().UTC().Add(-interval)
	for key, entry := range c.dict {
		if entry.createdAt.Before(timePassed) {
			delete(c.dict, key)
		}
	}
}

func (c *Cache) Add(key int, val pokeapi.LocationAreaResp) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dict[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key int) (pokeapi.LocationAreaResp, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.dict[key]
	return entry.val, ok
}
