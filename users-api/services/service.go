package services

import (
	"context"
	"net/http"
	"strings"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
)

type usersService struct {
	db db.Actions
}

func NewUsersService(db db.Actions) *usersService {
	return &usersService{db: db}
}
func (h *usersService) ValidateCredentials(email, password string) error {
	if len(strings.Trim(email, " \n\t")) == 0 ||
		!strings.Contains(email, "@") ||
		len(strings.Trim(password, " ")) < 7 {
		return common.HttpError{Code: http.StatusUnprocessableEntity, Message: "Invalid email or password."}
	}
	return nil
}
func (h *usersService) CheckUserExistence(email string) error {
	return nil
}
func (h *usersService) GetHashedPassword(password string) (string, error) {
	return "", nil
}
func (h *usersService) GetTokenForUser(password, hashedPassword string) (string, error) {
	return "", nil
}
func (h *usersService) SaveUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error) {
	return h.db.CreateUser(ctx, email, hashedPassword)
}
