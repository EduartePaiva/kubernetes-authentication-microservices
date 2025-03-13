package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/mocks"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/transports"
	"github.com/stretchr/testify/assert"
)

func createServiceForTest() *usersService {
	dbMock := mocks.NewActionMock()
	transport := transports.NewTransportService("REST")
	return NewUsersService(dbMock, transport)
}

type testCasesValidadeCred struct {
	Email   string
	Pass    string
	Message string
}

func TestValidateCredentials(t *testing.T) {
	s := createServiceForTest()
	tests := []testCasesValidadeCred{
		{
			Email:   "     \n   \n   \t",
			Pass:    "1234567",
			Message: "test if it errors when email is empty with whitespaces or enter values",
		},
		{
			Email:   "something#gmail.com",
			Pass:    "1234567",
			Message: "error if it not includes a @",
		},
		{
			Email:   "something@gmail.com",
			Pass:    "     12345      ",
			Message: "error if password len is bigger than 7 but it have whitespace around",
		},
	}

	for _, test := range tests {
		err := s.ValidateCredentials(test.Email, test.Pass)
		assert.Error(t, err, test.Message)
		_, ok := err.(common.HttpError)
		assert.True(t, ok, "test if error is a httpError")
	}
	err := s.ValidateCredentials("something@gmail.com", "1234567")
	assert.NoError(t, err, "test valid input")
}

func TestCheckUserExistence(t *testing.T) {
	dbMock := mocks.NewActionMock()
	transport := transports.NewTransportService("REST")
	s := NewUsersService(dbMock, transport)

	t.Run("when email exists in database", func(t *testing.T) {
		ctx := context.Background()
		email := "test@gmail.com"
		mk := dbMock.On("GetUserByEmail", ctx, email).Return(models.User{}, nil)
		defer mk.Unset()
		ok, err := s.CheckUserExistence(ctx, email)
		assert.True(t, ok)
		assert.NoError(t, err)
		dbMock.AssertExpectations(t)
	})

	t.Run("when email don't exists in database", func(t *testing.T) {
		ctx := context.Background()
		email := "test@gmail.com"
		mk := dbMock.On("GetUserByEmail", ctx, email).Return(models.User{}, common.HttpError{
			Code:    http.StatusNotFound,
			Message: "Hi test",
		})
		defer mk.Unset()
		ok, err := s.CheckUserExistence(ctx, email)
		assert.False(t, ok)
		assert.NoError(t, err)
		dbMock.AssertExpectations(t)
	})
	t.Run("when database erros with something that isn't Not Found", func(t *testing.T) {
		ctx := context.Background()
		email := "test@gmail.com"
		mk := dbMock.On("GetUserByEmail", ctx, email).Return(models.User{}, errors.New("others errors"))
		defer mk.Unset()
		ok, err := s.CheckUserExistence(ctx, email)
		assert.False(t, ok)
		assert.EqualError(t, err, "Failed to create user.")
		dbMock.AssertExpectations(t)
	})

}
