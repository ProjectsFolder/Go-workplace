package main

import (
    "context"
    "google.golang.org/grpc"
    gen "grpc/proto"
    "log"
    "os"
)

func main() {
    conn, _ := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
    client := gen.NewTestServiceClient(conn)

    sample := os.Args[1]
    resp, err := client.Do(context.Background(), &gen.Request{Name: sample})
    if err != nil {
        log.Fatalf("could not get answer: %v", err)
    }

    log.Println("Response:", resp.Message)
}
