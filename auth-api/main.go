package main

import (
	"log"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
)

var (
	SERVER_PORT            = common.EnvString("AUTH_SERVER_PORT", "3000")
	COMMUNICATION_PROTOCOL = common.EnvString("COMMUNICATION_PROTOCOL", "REST")
)

func main() {
	switch COMMUNICATION_PROTOCOL {
	case "REST":
		httpServer := NewHttpServer(":" + SERVER_PORT)
		log.Fatal(httpServer.Run())
	case "gRPC":
		gRPCServer := NewGRPCServer("localhost:" + SERVER_PORT)
		log.Fatal(gRPCServer.Run())
	default:
		log.Println("supported communication protocols are REST and gRPC")
		log.Fatalf("the communication protocol %s is unsupported", COMMUNICATION_PROTOCOL)
	}

}
