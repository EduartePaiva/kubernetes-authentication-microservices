package handlers

import (
	"errors"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/types"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/utils"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
)

type authHttpHandler struct {
	authService types.AuthService
}

func NewAuthHttpHandler(authService types.AuthService) *authHttpHandler {
	return &authHttpHandler{authService}
}

func (h *authHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		// health check
		w.WriteHeader(http.StatusOK)
	})
	router.HandleFunc("GET /hashed-pw/{password}", h.getHashedPassword)
	router.HandleFunc("GET /token", h.getToken)
	router.HandleFunc("GET /verify-token", h.getTokenConfirmation)
}

func (h *authHttpHandler) getHashedPassword(w http.ResponseWriter, r *http.Request) {
	password := r.PathValue("password")
	if password == "" {
		common.WriteError(w, http.StatusBadRequest, errors.New("the query param password is missing"))
		return
	}
	hashedPass, err := h.authService.CreatePasswordHash(password)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"hashed": hashedPass})

}
func (h *authHttpHandler) getToken(w http.ResponseWriter, r *http.Request) {

}
func (h *authHttpHandler) getTokenConfirmation(w http.ResponseWriter, r *http.Request) {

}
