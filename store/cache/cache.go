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

func (c *Cache) Delete(key string) error {
	delete(c.results, key)
	return nil
}

func (c *Cache) List() ([]qstore.Pair, error) {
	var results []qstore.Pair
	for k, v := range c.results {
		results = append(results, qstore.Pair{
			Key:   k,
			Value: v,
		})
	}
	return results, nil
}

func NewCache() *Cache {
	return &Cache{
		results: make(map[string]interface{}),
	}
}
