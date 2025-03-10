package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
)

type usersService struct {
	db db.Actions
}

var (
	authApiAddress = common.EnvString("AUTH_API_ADDRESS", "http://localhost:3000")
)

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
func (h *usersService) CheckUserExistence(ctx context.Context, email string) error {
	_, err := h.db.GetUserByEmail(ctx, email)
	_, ok := err.(common.HttpError)
	if ok {
		return common.HttpError{Message: "Failed to create user.", Code: http.StatusUnprocessableEntity}
	}
	if err != nil {
		return common.HttpError{Message: "Failed to create user.", Code: http.StatusInternalServerError}
	}
	return nil
}
func (h *usersService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return h.db.GetUserByEmail(ctx, email)
}

func (h *usersService) GetHashedPassword(ctx context.Context, password string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", authApiAddress+"/hashed-pw/"+password, nil)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to create user."}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to create user."}
	}
	body := make(map[string]string)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to create user."}
	}
	errBody, ok := body["error"]
	if ok {
		return "", common.HttpError{Code: res.StatusCode, Message: errBody}
	}
	return body["hashed"], nil
}
func (h *usersService) GetTokenForUser(ctx context.Context, password, hashedPassword string) (string, error) {
	data := map[string]string{
		"password":       password,
		"hashedPassword": hashedPassword,
	}
	jBytes, _ := json.Marshal(&data)
	req, err := http.NewRequestWithContext(ctx, "POST", authApiAddress+"/token", bytes.NewReader(jBytes))
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to verify user."}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to verify user."}
	}
	body := make(map[string]string)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to verify user."}
	}
	errBody, ok := body["error"]
	if ok {
		return "", common.HttpError{Code: res.StatusCode, Message: errBody}
	}
	return body["token"], nil
}
func (h *usersService) SaveUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error) {
	return h.db.CreateUser(ctx, email, hashedPassword)
}
