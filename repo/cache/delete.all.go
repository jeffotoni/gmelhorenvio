package cache

import (
	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
)

func DeleteAll() (err error) {
	cred, err := auth.GetCredentialsCache()
	if err != nil {
		return
	}
	config.Cache.Flush()
	cred.SaveCache()
	return
}
