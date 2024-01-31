package service

import (
	"context"

	v1 "real-kratosworld/api/kratosworld/v1"
)

// // UserService is a user service.
// type UserService struct {
// 	v1.UnimplementedKratosWorldServer

// 	uc  *biz.GreeterUsecase
// 	log *log.Helper
// }

// func NewUserService()

// func (s *KratosWorldService) Login(ctx context.Context) *LoginService {
// }

func (s *KratosWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}
