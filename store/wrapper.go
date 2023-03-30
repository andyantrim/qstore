package store

import (
	"github.com/andyantrim/qstore"
	"github.com/andyantrim/qstore/store/cache"
	"github.com/andyantrim/qstore/store/file"
)

// TODO: Implement a wrapper for the store that will allow us to
// Use the cache and have buffered writes to the store file in the background

type Wrapper struct {
	cache        Store
	file         Store
	writeChannel chan qstore.Pair
	tester       bool
}

func NewWrapper(path string) (*Wrapper, error) {
	c := make(chan qstore.Pair, 1)

	file, err := file.NewFileWriter(path, c)
	if err != nil {
		return nil, err
	}
	return &Wrapper{
		cache: cache.NewCache(),
		file:  file,
	}, nil
}

func (w *Wrapper) Get(key string) (interface{}, error) {
	if w.tester {
		return w.file.Get(key)
	}
	return w.cache.Get(key)
}

func (w *Wrapper) Set(key string, value interface{}) error {
	if w.tester {
		return w.file.Set(key, value)
	}
	err := w.cache.Set(key, value)
	if err == nil {
		go w.file.Set(key, value)
	}

	return err
}

func (w *Wrapper) Delete(key string) error {
	if err := w.cache.Delete(key); err != nil {
		return err
	}
	return w.file.Delete(key)
}

func (w *Wrapper) List() ([]qstore.Pair, error) {
	if w.tester {
		return w.file.List()
	}
	return w.cache.List()
}

func LoadAll() error {
	// TODO: Load all the data from the file into the cache

	return nil
}
