package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/fmts"
	z "github.com/jeffotoni/gmelhorenvio/internal/zerolog/error"
	"github.com/jeffotoni/gmelhorenvio/repo"
)

var (
	ErrRequiredEnvEmpty = errors.New("required env empty")
	ErrStatusNotOk      = errors.New("status not ok")
)

func GetToken() (cred *ApiCredentials, err error) {
	fname := "func GetToken()"
	service := "pkg/auth"
	dfLog := z.ElasticMin(service, fname)
	status := http.StatusInternalServerError

	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.TIMEOUT_HTTP_SEC,
	)
	defer cancel()

	endpoint := fmts.ConcatStr(
		config.MELHORENVIO_PROTOCOL,
		"://",
		config.MELHORENVIO_HOST,
		config.MELHORENVIO_ENDPOINT_GET_TOKEN,
	)
	println("Getting token at:", endpoint)

	if config.MELHORENVIO_CLIENT_ID == "" ||
		config.MELHORENVIO_CLIENT_SECRET == "" ||
		config.MELHORENVIO_REDIRECT_URI == "" ||
		config.MELHORENVIO_CODE == "" {
		err = ErrRequiredEnvEmpty
		return
	}

	formData := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {config.MELHORENVIO_CLIENT_ID},
		"client_secret": {config.MELHORENVIO_CLIENT_SECRET},
		"redirect_uri":  {config.MELHORENVIO_REDIRECT_URI},
		"code":          {config.MELHORENVIO_CODE},
	}
	println("..........................................................")
	fmt.Println("Form data:", formData)
	println("..........................................................")

	payload := strings.NewReader(formData.Encode())
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		payload,
	)
	if err != nil {
		dfLog.Msg(err.Error())
		return
	}

	req.Header.Add("Host", config.MELHORENVIO_HOST)
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	req.Header.Add("Accept", "application/json")

	resp, err := repo.HttpClient.Do(req)
	if err != nil {
		dfLog.Msg(err.Error())
		return
	}
	defer resp.Body.Close()
	status = resp.StatusCode

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		dfLog.Msg(err.Error())
		zlog(dfLog, status, nil).Msg(err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrStatusNotOk
		zlog(dfLog, status, respBody).Msg(err.Error())
		return
	}

	err = json.Unmarshal(respBody, &cred)
	if err != nil {
		zlog(dfLog, status, respBody).Msg(err.Error())
		return
	}

	cred.RefreshedAt = time.Now()
	if !cred.IsValid() {
		err = ErrInvalidCredentialsFields
		zlog(dfLog, status, respBody).Msg(err.Error())
		return
	}

	err = SaveCredentials(cred)
	if err != nil {
		zlog(dfLog, status, respBody).Msg(err.Error())
	}

	return
}
