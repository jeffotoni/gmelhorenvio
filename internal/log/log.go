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

func InsertLogMelhorEnvio(db *pg.DbConnection, status, msg, rawJson string) {
	if strings.ToLower(config.ENV_AMBI) == "local" {
		return
	}

	status = strings.ToUpper(status)

	query := fmts.ConcatStr(`
		INSERT INTO melhorenvio_log_`, strings.ToLower(config.ENV_AMBI), ` (
			log_status,
			log_msg,
			log_json
		) VALUES (
			$1, 
			$2,
			$3
		)
	`)

	go func() {
		_, err := db.Exec(
			*db.Ctx,
			query,
			status,
			msg,
			rawJson,
		)
		if err != nil {
			log.Println(err)
		}
	}()
}
