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
		utils.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"hashed": hashedPass})

}

type getTokenBody struct {
	Password       string `json:"password"`
	HashedPassword string `json:"hashedPassword"`
}

func (h *authHttpHandler) getToken(w http.ResponseWriter, r *http.Request) {
	requestBody := &getTokenBody{}
	err := common.ParseJSON(r, requestBody)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.authService.VerifyPasswordHash(requestBody.Password, requestBody.HashedPassword)
	if err != nil {
		utils.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}
	token := h.authService.CreateToken()
	common.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

type getTokenConBody struct {
	Token string `json:"token"`
}

func (h *authHttpHandler) getTokenConfirmation(w http.ResponseWriter, r *http.Request) {
	requestBody := &getTokenConBody{}
	err := common.ParseJSON(r, requestBody)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.authService.VerifyToken(requestBody.Token)
	if err != nil {
		utils.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}
	common.WriteJSON(w, http.StatusOK, make(map[string]string))
}
