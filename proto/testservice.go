package gen

import (
	"context"
	"fmt"
	"time"
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

	time.Sleep(5 * time.Second)

	return response, nil
}
