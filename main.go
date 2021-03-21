package main

import (
	"github.com/dineshtbits/golang-mongo-crud/mongo"
	"github.com/dineshtbits/golang-mongo-crud/model"
	"github.com/google/uuid"
)

func main() {
	user := model.User{
		ID: uuid.New().String(),
		FirstName: "Dinesh", 
		Roles: []string{"admin", "user"},
	}
	mongo.CreateUser(user)
}
