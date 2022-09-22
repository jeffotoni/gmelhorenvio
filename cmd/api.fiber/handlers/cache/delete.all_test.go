package cache

import (
	"bytes"
	"net/http"
	"testing"

	tests "github.com/jeffotoni/gmelhorenvio/internal/tests"

	"github.com/gofiber/fiber/v2"
)

// go test -v -run ^TestDeleteAll$
func TestDeleteAll(t *testing.T) {
	type args struct {
		method      string
		ctype       string
		header      map[string]string
		url         string
		urlReq      string
		handlerfunc func(c *fiber.Ctx) error
	}

	tt := []struct {
		name     string
		args     args
		want     int //status code
		bodyShow bool
	}{
		{name: "test_delete_all", args: args{
			method:      http.MethodPost,
			ctype:       "application/json",
			url:         "/v1/cache",
			urlReq:      "/v1/cache",
			handlerfunc: DeleteAll,
		}, want: 200, bodyShow: false},
	}

	for _, tt := range tt {
		tt1 := tt
		t.Run(tt1.name, func(t *testing.T) {
			//t.Parallel()
			tests.TestNewRequest(t, tt1.args.url, tt1.args.urlReq, tt1.args.method,
				tt1.args.handlerfunc,
				bytes.NewBuffer([]byte{}),
				tt1.args.ctype, tt1.args.header, tt1.want, tt1.bodyShow)
		})
	}
}
