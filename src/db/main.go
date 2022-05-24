package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Message struct {
	AuthorId string
	Content  string
	Date     string
}

func GetClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func LogMessage(message *Message, channelName string) {
	client := GetClient()
	collection := client.Database("discord").Collection(channelName)
	_, err := collection.InsertOne(context.TODO(), message)

	if err != nil {
		panic(err)
	}
}
