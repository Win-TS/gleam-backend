package firebasepkg

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/Win-TS/gleam-backend.git/config"
	"google.golang.org/api/option"
)

func FirebaseConn(ctx context.Context, cfg *config.Config) *auth.Client {
	opt := option.WithCredentialsFile(cfg.Db.Url)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing Auth client: %v\n", err)
	}

	return authClient
}
