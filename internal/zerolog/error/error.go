package zerror

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	z "github.com/jeffotoni/gmelhorenvio/internal/zerolog"
)

type Log struct {
	FiberCtx *fiber.Ctx
	Handler  string
	Status   int
	Service  string
	MsgUUID  string
}

func NewLog(c *fiber.Ctx, handler, service, msgUUID string, status int) Log {
	return Log{
		FiberCtx: c,
		Handler:  handler,
		Status:   status,
		Service:  service,
		MsgUUID:  msgUUID,
	}
}

func ElasticMin(service, fname string) *zerolog.Event {
	return log.Error().
		Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
		Str("data", time.Now().Format("2006-01-02 15:04:05")).
		Str("version", z.LOG_VERSION).
		Str("service", service).
		Str("host", z.LOG_HOST).
		Str("func", fname)
}

type responseError struct {
	Msg string `json:"msg"`
}

func MsgErr(msg string) responseError {
	return responseError{
		Msg: msg,
	}
}
