package auth

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/aws"
)

func SaveCredentials(cred *ApiCredentials) (err error) {
	cred.SaveCache()

	if usingAws() {
		log.Println("saving credentials on aws")
		err = saveCredentialsAws(cred)
	} else {
		log.Println("saving credentials locally")
		err = saveCredentialsLocal(cred)
	}

	if err != nil {
		log.Println("error saving credentials:", err)
		return
	}

	log.Println("credentials successfully saved")
	return
}

func saveCredentialsLocal(cred *ApiCredentials) error {
	data, err := json.Marshal(cred)
	if err != nil {
		return err
	}
	return os.WriteFile(config.MELHORENVIO_CREDENTIALS_FILE, data, 0644)
}

func saveCredentialsAws(cred *ApiCredentials) error {
	data, err := json.Marshal(cred)
	if err != nil {
		return err
	}

	_, err = aws.Upload(
		config.AWS_BUCKET_CREDENTIALS,
		config.AWS_KEY_CREDENTIALS,
		data,
	)
	return err
}
