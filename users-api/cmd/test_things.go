package main

import (
	"context"
	"log"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ps := "something"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(ps).SetServerAPIOptions(serverAPI)
	mongoClient, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected with mongo")
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	db := db.NewActions(mongoClient)

	log.Println("sending the get request")
	user, err := db.GetUserByEmail(context.Background(), "test@gmail.com")

	log.Println(user, err)

}
