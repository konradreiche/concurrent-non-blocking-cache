package main

import "sync"

type Cache struct {
	cache map[string]*entry
	sync.Mutex
}

type Func func() ([]byte, error)

type entry struct {
	value []byte
	err   error
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]*entry)}
}

// Coare grained locking results in a bad performance.
func (c *Cache) Get(key string, f Func) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	res, ok := c.cache[key]
	if !ok {
		res = &entry{}
		res.value, res.err = f()
		c.cache[key] = res
	}
	return res.value, res.err
}
