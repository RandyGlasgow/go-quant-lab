package auth_service

import (
	"log"
	"os"

	"github.com/authorizerdev/authorizer-go"
	"github.com/joho/godotenv"
)

var defaultHeaders = map[string]string{}

var AuthClient = createAuthClient()

func createAuthClient() *authorizer.AuthorizerClient {
	godotenv.Load()
	clientId := os.Getenv("AUTHORIZER_CLIENT_ID")
	authorizerUrl := os.Getenv("AUTHORIZER_URL")
	redirectUrl := ""

	log.Println("AUTHORIZER_CLIENT_ID: ", clientId)

	authorizerClient, err := authorizer.NewAuthorizerClient(clientId, authorizerUrl, redirectUrl, defaultHeaders)
	if err != nil {
		log.Fatal(err)
	}

	return authorizerClient
}
