package cache

import (
	"errors"

	"github.com/jeffotoni/gmelhorenvio/config"
)

var ErrCacheNotFound = errors.New("cache not found")

func DeleteKey(key string) error {
	if _, found := config.Cache.Get(key); !found {
		return ErrCacheNotFound
	}
	config.Cache.Delete(key)
	return nil
}
