package main

import (
	"context"
	"log"
	"os"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/pkg/database/mongodb"
	"github.com/Win-TS/gleam-backend.git/server"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// dbConn, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer dbConn.Close()

    // ctx := context.Background()

    // user, err := db.GetUser(ctx, dbConn, 1)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // log.Printf("User: %+v\n", user)

	ctx := context.Background()
	_ = ctx

	// Initiaize Config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	var db any

	if cfg.DbType.Type == "mongodb" {
		var db *mongo.Client
		db = mongodb.MongoConn(ctx, &cfg)
		defer db.Disconnect(ctx)
	} else if cfg.DbType.Type == "postgresql" {
		
	}
	
	
	server.Start(ctx, &cfg, db)
}