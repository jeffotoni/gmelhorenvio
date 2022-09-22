package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jeffotoni/gmelhorenvio/config"

	"github.com/rs/zerolog"
)

const CREDENTIAL_KEY = "credentials"

var (
	ErrCredentialsNotFound      = errors.New("credentials not found")
	ErrInvalidCredentialsStored = errors.New("invalid credentials stored")
	ErrInvalidCredentialsFields = errors.New("invalid credentials fields")
)

type ApiCredentials struct {
	ExpiresIn    int       `json:"expires_in"`
	RefreshedAt  time.Time `json:"refreshed_at"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type accessToken struct {
	Aud string `json:"aud"`
}

func (cred *ApiCredentials) IsValid() bool {
	if cred.ExpiresIn == 0 || cred.AccessToken == "" || cred.RefreshToken == "" {
		return false
	}

	tkParts := strings.Split(cred.AccessToken, ".")
	if len(tkParts) != 3 {
		return false
	}

	rawAccessTk, err := base64.StdEncoding.DecodeString(tkParts[1])
	if err != nil {
		return false
	}

	var accessTk accessToken
	if err := json.Unmarshal(rawAccessTk, &accessTk); err != nil {
		return false
	}

	if accessTk.Aud != config.MELHORENVIO_CLIENT_ID {
		return false
	}

	return true
}

func (cred *ApiCredentials) SaveCache() {
	config.Cache.Set(CREDENTIAL_KEY, cred, 0)
}

func GetCredentialsCache() (cred *ApiCredentials, err error) {
	rawCred, ok := config.Cache.Get(CREDENTIAL_KEY)
	if !ok {
		err = ErrCredentialsNotFound
		return
	}

	cred, ok = rawCred.(*ApiCredentials)
	if !ok {
		err = ErrInvalidCredentialsStored
	}
	return
}

func CalcExpirationRemaining(cred *ApiCredentials) time.Duration {
	expires := cred.RefreshedAt.Add(
		time.Second * time.Duration(cred.ExpiresIn-1),
	)
	return time.Until(expires)
}

func usingAws() bool {
	return config.AWS_BUCKET_CREDENTIALS != "" && config.AWS_KEY_CREDENTIALS != ""
}

func zlog(dfLog *zerolog.Event, code int, body []byte) *zerolog.Event {
	return dfLog.
		Int("Status-Code", code).
		RawJSON("Response-Body", body)
}
