package main

import (
	"log"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/handlers"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/services"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()
	mongo := db.NewActions()
	authService := services.NewUsersService(mongo)
	httpHandler := handlers.NewUsersHttpHandler(authService)
	httpHandler.RegisterRouter(router)
	log.Println("running http server on port", h.addr)
	return http.ListenAndServe(h.addr, router)
}
