package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client = DBInstance()

func DBInstance() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongoadmin:password@localhost:27018"))

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected To Mongodb")

	return client

}

func OpenConnection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	userCollection := client.Database(databaseName).Collection(collectionName)
	return userCollection
}
