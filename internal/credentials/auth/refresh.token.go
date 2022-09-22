package auth

import (
	"context"
	"encoding/json"
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

func RefreshToken() (cred *ApiCredentials, err error) {
	fname := "func RefreshToken()"
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
		config.MELHORENVIO_ENDPOINT_REFRESH_TOKEN,
	)

	cred, err = GetCredentialsCache()
	if err != nil {
		return
	}

	formData := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {config.MELHORENVIO_CLIENT_ID},
		"client_secret": {config.MELHORENVIO_CLIENT_SECRET},
		"refresh_token": {cred.RefreshToken},
	}
	payload := strings.NewReader(formData.Encode())
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		payload,
	)
	if err != nil {
		return
	}

	req.Header.Add("Host", config.MELHORENVIO_HOST)
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	req.Header.Add("Accept", "application/json")

	resp, err := repo.HttpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = ErrStatusNotOk
		zlog(dfLog, status, nil).Msg(err.Error())
		return
	}
	status = resp.StatusCode

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		dfLog.Msg(err.Error())
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
