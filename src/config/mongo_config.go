package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func MongoConnection(mongo_URI string, db_name string) *mongo.Database {
	if mongo_URI == "" || db_name == "" {
		log.Fatal("Mongo config missing!")
		panic("mongo config missing")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(mongo_URI).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connect failed", zap.Error(err))
		panic("PANIC - Mongo connect failed")
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping failed", zap.Error(err))
		panic("PANIC - Mongo ping failed")
	}
	return client.Database(db_name)
}