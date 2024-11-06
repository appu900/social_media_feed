package db

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("failed to connect to mongodb", err)
	}
	Client = client
	return Client
}

func GetCollections(client *mongo.Client, collectionName string) *mongo.Collection{
	collections := client.Database("auth-service").Collection(collectionName)
	return collections
}