package store

import "github.com/andyantrim/qstore"

type Store interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	List() ([]qstore.Pair, error)
}
