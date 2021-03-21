package model

type User struct {
	ID string `bson:"_id"`
	FirstName string `bson:"first_name"`
	LastName string `bson:"last_name"`
	Roles []string `bson:"roles"`
}