package models

import "go.mongodb.org/mongo-driver/v2/bson"

type InsertUserResult struct {
	ID    string
	Email string
}

type User struct {
	ID       bson.ObjectID `bson:"_id"`
	Password string        `bson:"password"`
	Email    string        `bson:"email"`
}
