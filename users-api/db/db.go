package db

import (
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/mongo"
)

type Actions interface {
	CreateUser(email, hashedPassword string) (models.User, error)
}

func NewActions() *mongo.MongoDB {
	return &mongo.MongoDB{}
}
