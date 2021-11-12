package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "os"
    gen "workplace/proto"
)

func main() {
    conn, _ := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
    client := gen.NewTestServiceClient(conn)
    
    name := os.Args[1]
    resp, err := client.Do(context.Background(), &gen.Request{Name: name})
    if err != nil {
        log.Fatalf("could not get answer: %v", err)
    }
    
    log.Println("Response:", resp.Message)
}
