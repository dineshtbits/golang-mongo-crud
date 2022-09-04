package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client, ctx = MongoClient()

var (
	Client = client
	CTX    = ctx
)

func MongoClient() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI_FOR_GO_CRUD")))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to local mongo....here are current databases")
	fmt.Println(databases)

	return client, ctx
}
