package main

import (
	"log"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
)

var (
	SERVER_PORT = common.EnvString("USERS_SERVER_PORT", "3001")
)

func main() {
	httpServer := NewHttpServer(":" + SERVER_PORT)
	log.Fatal(httpServer.Run())
}
