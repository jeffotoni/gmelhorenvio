package psql

import (
	"github.com/jeffotoni/gmelhorenvio/internal/env"
)

var (
	// ssl    = "verify-ca"
	DB_SSL          = env.GetString("DB_SSL", "require")
	DB_NAME_LOG     = env.GetString("DB_NAME_LOG", "s3wf.log")
	DB_HOST_LOG     = env.GetString("DB_HOST_LOG", "localhost")
	DB_USER_LOG     = env.GetString("DB_USER_LOG", "postgres")
	DB_PASSWORD_LOG = env.GetString("DB_PASSWORD_LOG", "")
	DB_PORT_LOG     = env.GetString("DB_PORT_LOG", "5432")
	//ssl_log         = "verify-ca"
	source_log = "postgres"

	DB_RECONNECT  = env.GetInt("DB_RECONNECT", 10)
	DB_WAIT_START = env.GetBool("DB_WAIT_START", false)

	//DB_CERT_PATH = ""
	dbCertServerCaPath string //="rds-certs/server-ca.pem"
	dbCertClientPath   string //= "rds-certs/client-cert.pem"
	dbCertKeyPath      string //= "rds-certs/client-key.pem"
	//DB_CERT_CLIENT_PATH = "./rds-certs/client-cert.pem"
)
