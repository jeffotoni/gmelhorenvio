package frete

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	tests "github.com/jeffotoni/gmelhorenvio/internal/tests"
	"github.com/jeffotoni/gmelhorenvio/models"

	"github.com/gofiber/fiber/v2"
)

// go test -v -run ^TestPostCalc$
func TestPostCalc(t *testing.T) {
	type req struct {
		From     models.PostalCode `json:"from"`
		To       models.PostalCode `json:"to"`
		Services string            `json:"services"`

		Products []models.Product       `json:"products"`
		Package  models.Package         `json:"package"`
		Options  models.ShipmentOptions `json:"options"`
	}

	type args struct {
		method      string
		ctype       string
		header      map[string]string
		url         string
		urlReq      string
		handlerfunc func(c *fiber.Ctx) error
		tmReq       req
	}

	tt := []struct {
		name     string
		args     args
		want     int //status code
		bodyShow bool
	}{
		{name: "test_post_products", args: args{
			method:      http.MethodPost,
			ctype:       "application/json",
			url:         "/v1/frete/calc",
			urlReq:      "/v1/frete/calc",
			handlerfunc: PostCalc,
			tmReq: req{
				From:     models.PostalCode{Code: "90570020"},
				To:       models.PostalCode{Code: "09040302"},
				Services: "1,2,18",
				Products: []models.Product{
					{
						Id:             "x",
						Width:          11,
						Height:         17,
						Length:         11,
						Weight:         0.3,
						InsuranceValue: 10.1,
						Quantity:       1,
					},
				},
				Options: models.ShipmentOptions{},
			},
		}, want: 200, bodyShow: true},

		{name: "test_post_package", args: args{
			method:      http.MethodPost,
			ctype:       "application/json",
			url:         "/v1/frete/calc",
			urlReq:      "/v1/frete/calc",
			handlerfunc: PostCalc,
			tmReq: req{
				From: models.PostalCode{Code: "90570020"},
				To:   models.PostalCode{Code: "09040302"},
				// Services: "1,2,18",
				Package: models.Package{
					Width:  11,
					Height: 17,
					Length: 11,
					Weight: 0.3,
				},
				// Options: models.ShipmentOptions{
				// 	InsuranceValue: 32,
				// },
			},
		}, want: 200, bodyShow: true},
	}

	for _, tt := range tt {
		requestBody, err := json.Marshal(&tt.args.tmReq)
		if err != nil {
			t.Errorf("Error json.Marshal:%s", err.Error())
			return
		}
		tt1 := tt
		t.Run(tt1.name, func(t *testing.T) {
			//t.Parallel()
			if tt1.bodyShow {
				fmt.Println("Json: ", string(requestBody))
			}
			tests.TestNewRequest(t, tt1.args.url, tt1.args.urlReq, tt1.args.method,
				tt1.args.handlerfunc,
				bytes.NewBuffer(requestBody),
				tt1.args.ctype, tt1.args.header, tt1.want, tt1.bodyShow)
		})
	}
}
