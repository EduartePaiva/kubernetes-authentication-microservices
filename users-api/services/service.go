package services

import (
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
func (h *usersService) SaveUser(email, hashedPassword string) (models.User, error) {
	return models.User{}, nil
}
