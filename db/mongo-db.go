package db

import (
	"context"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// todo replace this with local string - or from env
const uri = "mongodb://localhost:27017"

func InitMongo() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("could not connect to mongodb instance")
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal("somehow disconnected")
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal("something else went wrong")
	}

	color.Green("Successfully connected to mongoDB instance")
}
