package cache

import "github.com/jeffotoni/gmelhorenvio/config"

func DeleteAll() {
	config.Cache.Flush()
}
