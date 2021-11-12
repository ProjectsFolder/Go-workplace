package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"workplace/config"
	gen "workplace/proto"
)

func main() {
	server := grpc.NewServer()

	instance := new(gen.TestService)

	gen.RegisterTestServiceServer(server, instance)

	configuration := config.GetConfig()
	address := configuration.GrpcHost + ":" + configuration.GrpcPort
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
