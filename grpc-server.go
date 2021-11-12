package main

import (
    "google.golang.org/grpc"
    "log"
    "net"
    gen "workplace/proto"
)

func main() {
    server := grpc.NewServer()
    
    instance := new(gen.TestService)
    
    gen.RegisterTestServiceServer(server, instance)
    
    listener, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatal("Unable to create grpc listener:", err)
    }
    
    if err = server.Serve(listener); err != nil {
        log.Fatal("Unable to start server:", err)
    }
}
