package main

type Cache struct {
	requests chan request
}

type request struct {
	key      string
	response chan result
}

type Func func(key string) ([]byte, error)

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value []byte
	err   error
}

func NewCache(f Func) *Cache {
	cache := &Cache{requests: make(chan request)}
	go cache.server(f)
	return cache
}

func (c *Cache) Get(key string) ([]byte, error) {
	response := make(chan result)
	c.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (c *Cache) server(f Func) {
	cache := make(map[string]*entry)
	for req := range c.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
