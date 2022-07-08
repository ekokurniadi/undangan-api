package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
)

type googleDriveConfig struct {
	secretFile string
}

func NewGoogleDriveConfig(secretFile string) *googleDriveConfig {
	return &googleDriveConfig{secretFile}
}

func (g *googleDriveConfig) ServiceAccount() *http.Client {
	secret, err := ioutil.ReadFile(g.secretFile)
	if err != nil {
		log.Fatal("Error when reading the secret file credentials", err)
	}

	var clientData = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}

	json.Unmarshal(secret, &clientData)

	configuration := &jwt.Config{
		Email:      clientData.Email,
		PrivateKey: []byte(clientData.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}

	httpClient := configuration.Client(context.Background())

	return httpClient
}
