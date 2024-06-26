package singleflight

import "sync"

// Processing or ended request
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group is used to handle requests with different key
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do will ensure that for each distinct key, the function fn() will only be called once
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
