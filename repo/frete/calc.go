package frete

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
	"github.com/jeffotoni/gmelhorenvio/internal/fmts"
	"github.com/jeffotoni/gmelhorenvio/repo"
)

func Calc(body []byte) (respBody []byte, status int, err error) {
	status = http.StatusInternalServerError

	cred, err := auth.GetCredentialsCache()
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	// if !cred.IsValid() {
	// 	status = http.StatusInternalServerError
	// 	err = auth.ErrInvalidCredentialsFields
	// 	return
	// }

	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.TIMEOUT_HTTP_SEC,
	)
	defer cancel()

	endpoint := fmts.ConcatStr(
		config.MELHORENVIO_PROTOCOL,
		"://",
		config.MELHORENVIO_HOST,
		config.MELHORENVIO_ENDPOINT_FRETE_CALC,
	)

	payload := bytes.NewBuffer(body)
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
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmts.ConcatStr("Bearer ", cred.AccessToken))
	req.Header.Add("User-Agent", config.MELHORENVIO_USER_AGENT)

	//start := time.Now()
	resp, err := repo.HttpClient.Do(req)
	if err != nil {
		return
	}
	//end := time.Now()
	//fmt.Println("time:", start.Sub(end))
	defer resp.Body.Close()

	status = resp.StatusCode
	respBody, err = io.ReadAll(resp.Body)
	return
}
