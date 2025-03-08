package db

import (
	"context"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/mongo"
	sdk "go.mongodb.org/mongo-driver/v2/mongo"
)

type Actions interface {
	CreateUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error)
}

func NewActions(client *sdk.Client) *mongo.MongoDB {
	return mongo.NewMongoAction(client)
}
