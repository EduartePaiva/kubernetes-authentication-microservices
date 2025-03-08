package types

import (
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
)

type UsersService interface {
	// it returns a common.HttpError
	ValidateCredentials(email, password string) error
	CheckUserExistence(email string) error
	GetHashedPassword(password string) (string, error)
	GetTokenForUser(password, hashedPassword string) (string, error)
	SaveUser(email, hashedPassword string) (models.User, error)
}
