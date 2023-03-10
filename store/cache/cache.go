package cache

import (
	"fmt"

	"github.com/andyantrim/qstore"
)

type Cache struct {
	results map[string]interface{}
}

func (c *Cache) Get(key string) (interface{}, error) {
	val, ok := c.results[key]
	if !ok {
		return "", fmt.Errorf(qstore.ErrNotFound, key)
	}
	return val, nil
}

func (c *Cache) Set(key string, value interface{}) error {
	c.results[key] = value
	return nil
}

func NewCache() *Cache {
	return &Cache{
		results: make(map[string]interface{}),
	}
}
