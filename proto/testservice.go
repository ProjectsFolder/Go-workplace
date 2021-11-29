package gen

import (
    "context"
    "fmt"
    "gorm.io/gorm"
    "time"
    "workplace/internal/entity"
    "workplace/internal/injector"
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

    container := injector.GetContainer()
    container.Invoke(func(db *gorm.DB) {
        log := entity.GrpcLog{Message: req.GetName()}
        db.Create(&log)
    })

    time.Sleep(5 * time.Second)

    return response, nil
}
