package main

import (
	"log"
	"net"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/handlers"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/services"
	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *grpcServer {
	return &grpcServer{addr}
}

func (g *grpcServer) Run() error {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", g.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()
	svc := services.NewAuthService()
	handlers.NewGRPCHandler(grpcServer, svc)
	log.Println("Grpc server started at", g.addr)

	return grpcServer.Serve(l)
}
