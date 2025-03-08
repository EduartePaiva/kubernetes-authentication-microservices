package mocks

import (
	"context"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
)

type dbMock struct {
}

func NewActionMock() *dbMock {
	return &dbMock{}
}

func (m *dbMock) CreateUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error) {
	return models.InsertUserResult{}, nil
}
func (m *dbMock) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return models.User{}, nil
}
