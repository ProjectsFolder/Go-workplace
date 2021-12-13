package gen

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
    "gorm.io/gorm"
    "log"
    "time"
    "workplace/internal/entity"
    "workplace/internal/injector"
    "workplace/internal/services"
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
    container.Invoke(func(db *gorm.DB, rc *redis.Client, logger *services.Logger) {
        name := req.GetName()
        db.Create(&entity.GrpcLog{Message: name})
        rc.Set("grpc-redis", name, 60 * time.Second)
        if err := logger.Log(name); err != nil {
            log.Println("Unable to write log:", err)
        }
    })

    time.Sleep(5 * time.Second)

    return response, nil
}
