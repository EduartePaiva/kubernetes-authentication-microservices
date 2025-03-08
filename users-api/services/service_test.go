package services

import (
	"testing"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/mocks"
	"github.com/stretchr/testify/assert"
)

func createServiceForTest() *usersService {
	dbMock := mocks.NewActionMock()
	return NewUsersService(dbMock)
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
