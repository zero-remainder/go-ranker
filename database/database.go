package database

import (
	"context"
	"fmt"
	"github.com/zero-remainder/go-ranker/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo(cfg *config.Config) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	DB = client.Database(cfg.MongoDBName)
	fmt.Println("Connected to MongoDB!")
}

func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	return DB.Collection(name)
}
