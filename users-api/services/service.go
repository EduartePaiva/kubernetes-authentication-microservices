package services

import (
	"context"
	"net/http"
	"strings"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/types"
)

type usersService struct {
	db        db.Actions
	transport types.TransportsService
}

func NewUsersService(db db.Actions, transport types.TransportsService) *usersService {
	return &usersService{db: db, transport: transport}
}
func (h *usersService) ValidateCredentials(email, password string) error {
	if len(strings.Trim(email, " \n\t")) == 0 ||
		!strings.Contains(email, "@") ||
		len(strings.Trim(password, " ")) < 7 {
		return common.HttpError{Code: http.StatusUnprocessableEntity, Message: "Invalid email or password."}
	}
	return nil
}
func (h *usersService) CheckUserExistence(ctx context.Context, email string) (bool, error) {
	_, err := h.db.GetUserByEmail(ctx, email)
	httpErr, ok := err.(common.HttpError)
	if ok && httpErr.Code == http.StatusNotFound {
		return false, nil
	}
	if err != nil {
		return false, common.HttpError{Message: "Failed to create user.", Code: http.StatusInternalServerError}
	}
	return true, nil
}
func (h *usersService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return h.db.GetUserByEmail(ctx, email)
}

func (h *usersService) GetHashedPassword(ctx context.Context, password string) (string, error) {
	return h.transport.GetHashedPassword(ctx, password)
}
func (h *usersService) GetTokenForUser(ctx context.Context, password, hashedPassword string) (string, error) {
	return h.transport.GetToken(ctx, password, hashedPassword)
}
func (h *usersService) SaveUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error) {
	return h.db.CreateUser(ctx, email, hashedPassword)
}
