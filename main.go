package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/pkg/cronjob"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/server"
	_ "github.com/lib/pq"
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

	opt := option.WithCredentialsFile("config/gleam-firebase-6925b-firebase-adminsdk-qzvvk-11a1d6f129.json")
	config := &firebase.Config{ProjectID: cfg.Firebase.ProjectId, StorageBucket: cfg.Firebase.StorageBucket}
	fbApp, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := fbApp.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.DbType.Type == "firebase" {
		database = fbApp
	} else if cfg.DbType.Type == "postgres" {
		dbConn, err := sql.Open("postgres", cfg.Db.Url)
		if err != nil {
			log.Fatal(err)
		}
		defer dbConn.Close()
		if err = dbConn.Ping(); err != nil {
			log.Fatalf("Error connecting to the database: %v\n", err)
		}
		switch cfg.App.Name {
		case "user":
			database = userdb.NewStore(dbConn)
		case "group":
			database = groupdb.NewStore(dbConn)
			cron := cronjob.NewCronjobService(groupdb.NewStore(dbConn))
			if err := cron.Start(); err != nil {
				log.Fatalf("Error starting cronjob: %v", err)
			}
		}
	}

	server.Start(ctx, &cfg, database, client)
}
