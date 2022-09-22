package auth

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/aws"
)

func LoadCredentials() (cred *ApiCredentials, err error) {
	if usingAws() {
		log.Println("loading credentials from aws")
		cred, err = loadCredentialsAws()
	} else {
		log.Println("loading credentials locally")
		cred, err = loadCredentialsLocal()
	}

	if err != nil {
		log.Println("error loading credentials:", err)
		return
	}

	log.Println("credentials successfully loaded")
	cred.SaveCache()
	return
}

func loadCredentialsLocal() (cred *ApiCredentials, err error) {
	file, err := os.ReadFile(config.MELHORENVIO_CREDENTIALS_FILE)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &cred)
	return
}

func loadCredentialsAws() (cred *ApiCredentials, err error) {
	file, err := aws.Donwload(
		config.AWS_BUCKET_CREDENTIALS,
		config.AWS_KEY_CREDENTIALS,
	)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &cred)
	return
}
