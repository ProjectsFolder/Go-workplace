package gen

import (
    "context"
    "fmt"
    "time"
    "workplace/database"
    "workplace/entity"
)

type TestService struct {
}

func (s *TestService) mustEmbedUnimplementedTestServiceServer() {
    panic("implement me")
}

func (s *TestService) Do(ctx context.Context, req *Request) (*Response, error) {
    beautiful := ""
    if req.GetBeautiful() {
        beautiful = "!"
    }

    response := new(Response)
    response.Message = fmt.Sprintf("Hello, %s%s", req.GetName(), beautiful)

    db, err := database.GetConnection()
    if err != nil {
        panic("failed to connect database")
    }
    log := entity.GrpcLog{Message: req.GetName()}
    db.Create(&log)

    time.Sleep(5 * time.Second)

    return response, nil
}
