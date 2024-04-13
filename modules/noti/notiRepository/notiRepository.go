package notiRepository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type (
	NotiRepositoryService interface{
	}
	notiRepository struct{
		messagingClient *messaging.Client
	}
)

func NewNotiRepository(app *firebase.App) NotiRepositoryService {
	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Error - initializing Messaging client: %v\n", err)
		return nil
	}
	return &notiRepository{messagingClient}
}