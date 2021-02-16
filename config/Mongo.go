package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection bla bla bla...
func MongoConnection() (*mongo.Client, context.Context) {
	dbIP := "localhost"
	dbPORT := "27017"
	ctx := context.TODO()
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", dbIP, dbPORT))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
	/*
		ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
		if err != nil {

		}
		clientOptions := options.Client().ApplyURI("mongodb://" + dbIP + ":" + dbPORT)
		var client *mongo.Client
		client, _ = mongo.Connect(ctx, clientOptions)
		return client, ctx
	*/
}
