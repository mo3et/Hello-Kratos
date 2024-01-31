package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewKratosWorldService)

// // KratosWorldService is a greeter service.
// type KratosWorldService struct {
// 	v1.UnimplementedKratosWorldServer

// 	uc  *biz.GreeterUsecase
// 	log *log.Helper
// }

// func NewKratosWorldService(uc *biz.GreeterUsecase, logger log.Logger) *KratosWorldService {
// 	return &KratosWorldService{uc: uc, log: log.NewHelper(logger)}
// }
