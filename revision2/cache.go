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

// URLs are fetched twice, can be improved with duplicate surpression.
func (c *Cache) Get(key string, f Func) ([]byte, error) {
	c.Lock()
	res, ok := c.cache[key]
	c.Unlock()
	if !ok {
		res = &entry{}
		res.value, res.err = f()
		c.Lock()
		c.cache[key] = res
		c.Unlock()
	}
	return res.value, res.err
}
