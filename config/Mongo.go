package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection bla bla bla...
func MongoConnection() (*mongo.Client, context.Context) {
	dbIP := "localhost"
	dbPORT := "27017"
	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {

	}
	clientOptions := options.Client().ApplyURI("mongodb://" + dbIP + ":" + dbPORT)
	var client *mongo.Client
	client, _ = mongo.Connect(ctx, clientOptions)
	return client, ctx
}
