package repo

import (
	"crypto/tls"
	"net/http"

	"github.com/jeffotoni/gmelhorenvio/config"
)

var HttpClient = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives:   true,
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.TLS_SECURE,
		},
	},
}
