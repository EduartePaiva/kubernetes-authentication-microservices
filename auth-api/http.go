package main

import (
	"log"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/handlers"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/services"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()
	authService := services.NewAuthService()
	httpHandler := handlers.NewAuthHttpHandler(authService)
	httpHandler.RegisterRouter(router)
	log.Println("running http server on port", h.addr)
	return http.ListenAndServe(h.addr, router)
}
