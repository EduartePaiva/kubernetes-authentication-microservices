package mongo

import (
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
)

type MongoDB struct{}

func (m *MongoDB) CreateUser(email, hashedPassword string) (models.User, error) {
	return models.User{}, nil
}
