package log

import (
	"context"
	"testing"

	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	cLog := pg.NewConn(&ctx, pg.ConfigLog)

	InsertLogMelhorEnvio(
		cLog,
		"TESTING",
		"testing",
		"",
	)
}
