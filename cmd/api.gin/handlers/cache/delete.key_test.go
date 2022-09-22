package cache

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/jeffotoni/gmelhorenvio/config"
	tests "github.com/jeffotoni/gmelhorenvio/internal/tests"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"

	"github.com/gofiber/fiber/v2"
)

// go test -v -run ^TestDeleteKey$
func TestDeleteKey(t *testing.T) {
	key := cache.GetHash([]byte(`
	{
		"from": {
			"postal_code": "96020360"
		},
		"to": {
			"postal_code": "01018020"
		},
		"products": [
			{
				"id": "x",
				"width": 11,
				"height": 17,
				"length": 11,
				"weight": 0.3,
				"insurance_value": 10.1,
				"quantity": 1
			}
		],
		"options": {
			"receipt": false,
			"own_hand": false
		},
		"services": "1,2,18"
	}
	`))

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
		{name: "test_delete_key", args: args{
			method:      http.MethodPost,
			ctype:       "application/json",
			url:         "/v1/cache/:key",
			urlReq:      "/v1/cache/" + key,
			handlerfunc: DeleteKey,
		}, want: 200, bodyShow: true},
	}

	for _, tt := range tt {
		tt1 := tt
		t.Run(tt1.name, func(t *testing.T) {
			//t.Parallel()
			config.Cache.Set(key, "testing", 0)
			tests.TestNewRequest(t, tt1.args.url, tt1.args.urlReq, tt1.args.method,
				tt1.args.handlerfunc,
				bytes.NewBuffer([]byte{}),
				tt1.args.ctype, tt1.args.header, tt1.want, tt1.bodyShow)
		})
	}
}
