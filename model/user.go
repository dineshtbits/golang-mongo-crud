package model

type User struct {
	ID string `bson:"_id"`
	FirstName string `bson:"first_name"`
	Roles []string `bson:"roles"`
}