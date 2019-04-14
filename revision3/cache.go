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
	ready chan struct{}
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]*entry)}
}

func (c *Cache) Get(key string, f Func) ([]byte, error) {
	c.Lock()
	e := c.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		c.cache[key] = e
		c.Unlock()

		e.value, e.err = f()
		close(e.ready)
	} else {
		c.Unlock()
		<-e.ready
	}
	return e.value, e.err
}
