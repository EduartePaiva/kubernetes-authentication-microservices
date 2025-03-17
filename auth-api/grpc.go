package main

import (
	"log"
	"net"
	"time"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/handlers"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type grpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *grpcServer {
	return &grpcServer{addr}
}

func (g *grpcServer) Run() error {
	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAge:      time.Second * 30,
		MaxConnectionAgeGrace: time.Second * 10,
	}))
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
