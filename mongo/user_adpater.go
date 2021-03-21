package mongo

import (
	"fmt"
	"log"
	"github.com/dineshtbits/golang-mongo-crud/model"
)

func CreateUser(user model.User) {
	client, ctx := MongoClient()
	database := client.Database("quickstart")
	users := database.Collection("users")
	fmt.Println(user)
	insertedUser, err := users.InsertOne(ctx, user)
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted user")
	fmt.Println(insertedUser.InsertedID)
}
