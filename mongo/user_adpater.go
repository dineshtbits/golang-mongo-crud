package mongo

import (
	"fmt"
	"github.com/dineshtbits/golang-mongo-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CreateUser(user model.User) *mongo.InsertOneResult {
	database := Client.Database("quickstart")
	users := database.Collection("users")
	insertedUser, err := users.InsertOne(CTX, user)
	if err != nil {
		log.Panic(err)
	}
	return insertedUser
}

func FindAllUsers() []model.User {
	database := Client.Database("quickstart")
	cursor, err := database.Collection("users").Find(CTX, bson.M{})
	if err != nil {
		panic(err)
	}
	var users []model.User
	if err := cursor.All(CTX, &users); err != nil {
		panic(err)
	}
	return users
}

func DeleteUser(uid string) *mongo.DeleteResult {
	database := Client.Database("quickstart")
	users := database.Collection("users")
	result, err := users.DeleteOne(CTX, bson.M{"_id": uid})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
	return result
}

func UpdateUser(uid string, user model.User) *mongo.UpdateResult {
	database := Client.Database("quickstart")
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

	updateResult, err := users.UpdateOne(CTX, filter, setMap)

	if err != nil {
		log.Panic(err)
	}
	return updateResult
}
