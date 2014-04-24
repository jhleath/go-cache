package cache

import (
	"sync"
	"time"
)

// This is the basic cache that has a data map,
// a TTL (time.Duration) that specifies how long
// a cache entry can live, and a RWMutex lock.
type Cache struct {
	Data map[string]Record
	TTL  time.Duration
	Lock sync.RWMutex
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		Data: make(map[string]Record),
		TTL:  ttl,
	}
}

// Store a new item in the cache with a key.
func (c *Cache) Store(key string, value interface{}) {
	c.Lock.Lock()
	c.Data[key] = Record{value, time.Now()}
	c.Lock.Unlock()
}

// Get a value from the cache. Stale will be true if the record is older
// than the Cache TTL OR the item doesn't exist in the cache.
func (c *Cache) Get(key string) (value interface{}, stale bool) {
	c.Lock.RLock()
	obj, ok := c.Data[key]
	c.Lock.RUnlock()
	if !ok {
		return nil, true
	}
	return obj.Value, (time.Now().Sub(obj.Added) > c.TTL)
}

// A cache record with a value and an added time.
type Record struct {
	Value interface{}
	Added time.Time
}
