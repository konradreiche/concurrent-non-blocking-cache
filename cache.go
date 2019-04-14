package main

type Cache struct {
	cache map[string]*entry
}

type Func func() ([]byte, error)

type entry struct {
	value []byte
	err   error
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]*entry)}
}

func (c *Cache) Get(key string, f Func) ([]byte, error) {
	res, ok := c.cache[key]
	if !ok {
		res = &entry{}
		res.value, res.err = f()
		c.cache[key] = res
	}
	return res.value, res.err
}
