package handlers

import (
	"errors"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/types"
)

type usersHttpHandler struct {
	service types.UsersService
}

func NewUsersHttpHandler(usersService types.UsersService) *usersHttpHandler {
	return &usersHttpHandler{service: usersService}
}

func (h *usersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.HandleFunc("POST /signup", h.createUser)
	router.HandleFunc("POST /login", h.verifyUser)
}

type credentialsBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *usersHttpHandler) createUser(w http.ResponseWriter, r *http.Request) {
	cred := credentialsBody{}
	err := common.ParseJSON(r, &cred)
	if err != nil || cred.Password == "" || cred.Email == "" {
		common.WriteError(w, http.StatusBadRequest, errors.New("missing or invalid fields"))
		return
	}
	err = h.service.ValidateCredentials(cred.Email, cred.Password)
	if err != nil {
		common.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}
	err = h.service.CheckUserExistence(r.Context(), cred.Email)
	if err != nil {
		common.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}
	hashedPassword, err := h.service.GetHashedPassword(r.Context(), cred.Password)
	if err != nil {
		common.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}
	user, err := h.service.SaveUser(r.Context(), cred.Email, hashedPassword)
	if err != nil {
		common.HandleHttpError(err, w, http.StatusInternalServerError)
		return
	}

	// best practice, never return a object directly to the user
	common.WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "User created.",
		"user": map[string]string{
			"email": user.Email,
			"id":    user.ID,
		},
	})
}

type verifyUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *usersHttpHandler) verifyUser(w http.ResponseWriter, r *http.Request) {
	reqBody := verifyUserBody{}
	err := common.ParseJSON(r, &reqBody)
	if err != nil {
		common.HandleHttpError(common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to parse request body"}, w, 500)
		return
	}

	user, err := h.service.GetUserByEmail(r.Context(), reqBody.Email)
	if err != nil {
		common.HandleHttpError(err, w, 500)
		return
	}
	token, err := h.service.GetTokenForUser(r.Context(), reqBody.Password, user.HashedPassword)
	if err != nil {
		common.HandleHttpError(err, w, 500)
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"token": token, "userId": user.ID.String()})
}
