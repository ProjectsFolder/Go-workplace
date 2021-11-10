package gen

import "context"

type TestService struct {
}

func (s *TestService) mustEmbedUnimplementedTestServiceServer() {
    panic("implement me")
}

func (s *TestService) Do(ctx context.Context, req *Request) (*Response, error) {
    response := new(Response)

    response.Message = "Hello, grpc! Request: " + req.Name

    return response, nil
}
