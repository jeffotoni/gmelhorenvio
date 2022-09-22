package main

import (
	"log"
	"time"

	"github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
)

func main() {
	_, err := auth.GetToken()
	log.Println("GetToken error:", err)

	time.Sleep(time.Second * 5)

	_, err = auth.RefreshToken()
	log.Println("GetToken error:", err)
}
