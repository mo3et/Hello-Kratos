package service

import (
	"context"

	v1 "real-kratosworld/api/kratosworld/v1"
	"real-kratosworld/internal/biz"
)

// KratosWorldService is a greeter service.
type KratosWorldService struct {
	v1.UnimplementedKratosWorldServiceServer

	uc *biz.GreeterUsecase
}

// NewKratosWorldService new a greeter service.
func NewKratosWorldService(uc *biz.GreeterUsecase) *KratosWorldService {
	return &KratosWorldService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *KratosWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
