package mongo

import (
	"fmt"
	"github.com/dineshtbits/golang-mongo-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CreateUser(user model.User) *mongo.InsertOneResult {
	client, ctx := MongoClient()
	database := client.Database("quickstart")
	users := database.Collection("users")
	insertedUser, err := users.InsertOne(ctx, user)
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func FindAllUsers() []model.User {
	client, ctx := MongoClient()
	database := client.Database("quickstart")
	cursor, err := database.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	var users []model.User
	if err := cursor.All(ctx, &users); err != nil {
		panic(err)
	}
	return users
}

func DeleteUser(uid string) *mongo.DeleteResult {
	client, ctx := MongoClient()
	database := client.Database("quickstart")
	users := database.Collection("users")
	result, err := users.DeleteOne(ctx, bson.M{"_id": uid})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return result
}

func UpdateUser(uid string, user model.User) *mongo.UpdateResult {
	client, ctx := MongoClient()
	database := client.Database("quickstart")
	users := database.Collection("users")
	fmt.Println(user)
	filter := bson.M{"_id": uid}

	updates := bson.D{}

	if len(user.FirstName) > 0 {
		updates = append(updates, bson.E{"first_name", user.FirstName})
	}

	if len(user.LastName) > 0 {
		updates = append(updates, bson.E{"last_name", user.LastName})
	}

	if user.Roles != nil {
		updates = append(updates, bson.E{"roles", user.Roles})
	}

	setMap := bson.D{
		{"$set", updates},
	}

	updateResult, err := users.UpdateOne(ctx, filter, setMap)
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	return updateResult
}
