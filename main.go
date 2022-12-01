package main

import (
	"context"
	"log"

	apifiber "github.com/jeffotoni/gmelhorenvio/cmd/api.fiber"
	apigin "github.com/jeffotoni/gmelhorenvio/cmd/api.gin"
	"github.com/jeffotoni/gmelhorenvio/config"
	auth "github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
)

func main() {
	cred, err := auth.LoadCredentials()
	if err != nil {
		log.Println("Error loadFile:", err)
		return
	}
	log.Println("version 0.0.0.3")
	log.Println("credentials expires in:", auth.CalcExpirationRemaining(cred))

	ctx := context.Background()
	cLog := pg.NewConn(&ctx, pg.ConfigLog)

	switch config.API_FRAMEWORK {
	case "fiber":
		apifiber.Run(cLog)
	case "gin":
		apigin.Run(cLog)
	default:
		apifiber.Run(cLog)
	}
}
