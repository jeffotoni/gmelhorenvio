package config

import (
	"time"

	"github.com/jeffotoni/gmelhorenvio/internal/env"
)

var (
	ENV_AMBI         = env.GetString("ENV_AMBI", "dev") // dev | prod | local
	SERVER_DOMAIN    = env.GetString("SERVER_DOMAIN", "0.0.0.0:8080")
	API_TEST_PORT    = env.GetString("API_TEST_PORT", "8081")
	TLS_SECURE       = env.GetBool("TLS_SECURE", true)
	API_STATIC_TOKEN = env.GetString("API_STATIC_TOKEN", "")
	API_FRAMEWORK    = env.GetString("API_FRAMEWORK", "fiber")

	AWS_REGION             = env.GetString("AWS_REGION", "us-east-1")
	AWS_BUCKET_CREDENTIALS = env.GetString("AWS_BUCKET_CREDENTIALS", "")
	AWS_KEY_CREDENTIALS    = env.GetString("AWS_KEY_CREDENTIALS", "")

	LIMITER_MAX_REQUESTS   = env.GetInt("MAX_REQUESTS", 100000)
	LIMITER_EXPIRATION_SEC = time.Second * time.Duration(env.GetInt("LIMITER_EXPIRATION_SEC", 60))

	TIMEOUT_HTTP_SEC              = time.Second * time.Duration(env.GetInt("TIMEOUT_HTTP_SEC", 30))
	DELAY_ACCESS_TOKEN_RETRY_SEC  = time.Second * time.Duration(env.GetInt("DELAY_ACCESS_TOKEN_RETRY_SEC", 1))
	DELAY_REFRESH_TOKEN_DAY       = DAY * time.Duration(env.GetInt("DELAY_REFRESH_TOKEN_DAY", 1))
	DELAY_REFRESH_TOKEN_RETRY_SEC = time.Second * time.Duration(env.GetInt("DELAY_REFRESH_TOKEN_RETRY_SEC", 1))

	CACHE_EXPIRATION_MIN = time.Minute * time.Duration(env.GetInt("CACHE_EXPIRATION_MIN", 30))
	CACHE_CLEANUP_MIN    = time.Minute * time.Duration(env.GetInt("CACHE_CLEANUP_MIN", 60))

	MELHORENVIO_CODE          = env.GetString("MELHORENVIO_CODE", "")
	MELHORENVIO_CLIENT_ID     = env.GetString("MELHORENVIO_CLIENT_ID", "")
	MELHORENVIO_CLIENT_SECRET = env.GetString("MELHORENVIO_CLIENT_SECRET", "")
	MELHORENVIO_REDIRECT_URI  = env.GetString("MELHORENVIO_REDIRECT_URI", "https://www.mongeembalagensonline.com.br")

	MELHORENVIO_PROTOCOL         = env.GetString("MELHORENVIO_PROTOCOL", "https")
	MELHORENVIO_HOST             = env.GetString("MELHORENVIO_HOST", "sandbox.melhorenvio.com.br")
	MELHORENVIO_USER_AGENT       = env.GetString("MELHORENVIO_USER_AGENT", `Monge Embalagens Online (contato@mongeembalagens.com.br)`)
	MELHORENVIO_CREDENTIALS_FILE = env.GetString("MELHORENVIO_CREDENTIALS_FILE", "./cmd/credentials/credentials.json")

	MELHORENVIO_ENDPOINT_GET_TOKEN = env.GetString(
		"MELHORENVIO_ENDPOINT_AUTH_TOKEN",
		"/oauth/token",
	)

	MELHORENVIO_ENDPOINT_REFRESH_TOKEN = env.GetString(
		"MELHORENVIO_ENDPOINT_AUTH_TOKEN",
		"/oauth/token",
	)

	MELHORENVIO_ENDPOINT_FRETE_CALC = env.GetString(
		"MELHORENVIO_ENDPOINT_FRETE_CALC",
		"/api/v2/me/shipment/calculate",
	)
)
