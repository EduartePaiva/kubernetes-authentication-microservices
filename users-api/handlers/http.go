package handlers

import (
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/types"
)

type usersHttpHandler struct{}

func NewUsersHttpHandler(usersService types.UsersService) *usersHttpHandler {
	return &usersHttpHandler{}
}

func (h *usersHttpHandler) RegisterRouter(router *http.ServeMux) {}
