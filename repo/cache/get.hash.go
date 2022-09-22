package cache

import "encoding/base64"

func GetHash(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}
