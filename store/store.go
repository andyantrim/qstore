package store

type Store interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
