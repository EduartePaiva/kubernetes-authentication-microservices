package main

import (
	"log"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
)

var (
	SERVER_PORT            = common.EnvString("USERS_SERVER_PORT", "3001")
	COMMUNICATION_PROTOCOL = common.EnvString("COMMUNICATION_PROTOCOL", "REST")
)

func main() {
	switch COMMUNICATION_PROTOCOL {
	case "REST":
		httpServer := NewHttpServer(":" + SERVER_PORT)
		log.Fatal(httpServer.Run())
	case "gRPC":
		log.Fatal("gRPC still in development")
	default:
		log.Println("supported communication protocols are REST and gRPC")
		log.Fatalf("the communication protocol %s is unsupported", COMMUNICATION_PROTOCOL)
	}

}
