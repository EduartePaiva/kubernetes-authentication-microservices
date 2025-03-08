package handlers

import (
	"net/http"

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
	router.HandleFunc("POST /signup", h.signUp)
	router.HandleFunc("POST /login", h.login)
}

func (h *usersHttpHandler) signUp(w http.ResponseWriter, r *http.Request) {

}
func (h *usersHttpHandler) login(w http.ResponseWriter, r *http.Request) {

}
