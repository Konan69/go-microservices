package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port     = "80"
	rpcPort  = "5001"
	mongoUrl = "mongodb://mongo:27017"
	gRPCPort = "50001"
)

var client *mongo.Client

type Cofig struct{}

func main() {
	//connect to mongodb
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	//connect to mongodb

	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return c, nil

}
