package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/Win-TS/gleam-backend.git/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoConn(ctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url)) 
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error pinging to database:", err)
	}

	return client
}