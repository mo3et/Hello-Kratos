// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: logger_job.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoggerJobClient is the client API for LoggerJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerJobClient interface {
}

type loggerJobClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerJobClient(cc grpc.ClientConnInterface) LoggerJobClient {
	return &loggerJobClient{cc}
}

// LoggerJobServer is the server API for LoggerJob service.
// All implementations must embed UnimplementedLoggerJobServer
// for forward compatibility
type LoggerJobServer interface {
	mustEmbedUnimplementedLoggerJobServer()
}

// UnimplementedLoggerJobServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerJobServer struct {
}

func (UnimplementedLoggerJobServer) mustEmbedUnimplementedLoggerJobServer() {}

// UnsafeLoggerJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerJobServer will
// result in compilation errors.
type UnsafeLoggerJobServer interface {
	mustEmbedUnimplementedLoggerJobServer()
}

func RegisterLoggerJobServer(s grpc.ServiceRegistrar, srv LoggerJobServer) {
	s.RegisterService(&LoggerJob_ServiceDesc, srv)
}

// LoggerJob_ServiceDesc is the grpc.ServiceDesc for LoggerJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggerJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.job.v1.LoggerJob",
	HandlerType: (*LoggerJobServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "logger_job.proto",
}
