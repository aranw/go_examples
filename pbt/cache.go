package pbt

import (
	"fmt"
	"sync"
)

func NewCache() *Cache {
	return &Cache{
		values: make(map[string]string),
	}
}

type Cache struct {
	values map[string]string
	mu     sync.Mutex
}

func (c *Cache) Add(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[k] = v
}

func (c *Cache) Get(k string) (string, error) {
	var v string
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.values[k]; ok {
		return v, nil
	}

	return v, fmt.Errorf("value not found")
}
