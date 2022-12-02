package log

import (
	"log"
	"strings"

	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/fmts"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
)

var (
	ErrStatus = "ERROR"
)

func InsertLogMelhorEnvio(db *pg.DbConnection, status, service, msg, rawJson string) {
	if !config.DB_INSERT_LOG {
		return
	}

	status = strings.ToUpper(status)

	query := fmts.ConcatStr(`
		INSERT INTO melhorenvio_log_`, strings.ToLower(config.ENV_AMBI), ` (
			log_status,
			log_ecommerce,
			log_msg,
			log_json
		) VALUES (
			$1, 
			$2,
			$3,
			$4
		)
	`)

	go func() {
		_, err := db.Exec(
			*db.Ctx,
			query,
			status,
			service,
			msg,
			rawJson,
		)
		if err != nil {
			log.Println(err)
		}
	}()
}
