package mocks

import (
	"context"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
	"github.com/stretchr/testify/mock"
)

type dbMock struct {
	mock.Mock
}

func NewActionMock() *dbMock {
	return &dbMock{}
}

func (m *dbMock) CreateUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error) {
	return models.InsertUserResult{}, nil
}
func (m *dbMock) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(models.User), args.Error(1)
}
