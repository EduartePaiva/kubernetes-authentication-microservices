package handlers

import (
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/types"
)

type authHttpHandler struct {
	authService types.AuthService
}

func NewAuthHttpHandler(authService types.AuthService) *authHttpHandler {
	return &authHttpHandler{authService}
}

func (h *authHttpHandler) RegisterRouter(router *http.ServeMux) {

}
