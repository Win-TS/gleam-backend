package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	firebase "firebase.google.com/go"
	//dbuser "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/pkg/database/mongodb"
	"github.com/Win-TS/gleam-backend.git/server"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()

	// Initiaize Config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	var database interface{}

	if cfg.DbType.Type == "mongodb" {
		client := mongodb.MongoConn(ctx, &cfg)
		defer client.Disconnect(ctx)
		database = client
	} else if cfg.DbType.Type == "firebase" {
		opt := option.WithCredentialsFile("config/gleam-firebase-6925b-firebase-adminsdk-qzvvk-11a1d6f129.json")
		config := &firebase.Config{ProjectID: cfg.Firebase.ProjectId}
		app, err := firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
		database = app
	} else if cfg.DbType.Type == "postgres" {
		dbConn, err := sql.Open("postgres", cfg.Db.Url)
		if err != nil {
			log.Fatal(err)
		}
		defer dbConn.Close()
		if err = dbConn.Ping(); err != nil {
			log.Fatalf("Error connecting to the database: %v\n", err)
		}
		database = dbConn
	}

	server.Start(ctx, &cfg, database)
}