package types

import (
	"context"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
)

type UsersService interface {
	// it returns a common.HttpError
	ValidateCredentials(email, password string) error
	CheckUserExistence(ctx context.Context, email string) (bool, error)
	GetHashedPassword(ctx context.Context, password string) (string, error)
	GetTokenForUser(ctx context.Context, password, hashedPassword string) (string, error)
	SaveUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type TransportsService interface {
	GetHashedPassword(ctx context.Context, password string) (string, error)
	GetToken(ctx context.Context, password, hashedPassword string) (string, error)
	GetTokenConfirmation(ctx context.Context, token string) (bool, error)
}
