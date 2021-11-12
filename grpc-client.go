package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "os"
    "workplace/config"
    gen "workplace/proto"
)

func main() {
    configuration := config.GetConfig()
    address := configuration.GrpcHost + ":" + configuration.GrpcPort
    conn, _ := grpc.Dial(address, grpc.WithInsecure())
    client := gen.NewTestServiceClient(conn)
    
    name := os.Args[1]
    resp, err := client.Do(context.Background(), &gen.Request{Name: name})
    if err != nil {
        log.Fatalf("could not get answer: %v", err)
    }
    
    log.Println("Response:", resp.Message)
}
