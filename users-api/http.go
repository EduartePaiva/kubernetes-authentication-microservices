package main

import (
	"context"
	"log"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/handlers"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/services"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/transports"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

var (
	COMMUNICATION_PROTOCOL = common.EnvString("COMMUNICATION_PROTOCOL", "REST")
)

func (h *httpServer) Run() error {
	// create mongo mongoClient
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(common.EnvString("MONGODB_CONNECTION_URI", "")))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	router := http.NewServeMux()
	mongo := db.NewActions(mongoClient)
	transport := transports.NewTransportService(COMMUNICATION_PROTOCOL)
	usersService := services.NewUsersService(mongo, transport)
	httpHandler := handlers.NewUsersHttpHandler(usersService)
	httpHandler.RegisterRouter(router)
	log.Println("running http server on port", h.addr)
	return http.ListenAndServe(h.addr, router)
}
