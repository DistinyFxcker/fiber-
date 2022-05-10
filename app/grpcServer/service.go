package grpcServer

import (
	service "cdnFiber/app"
	"log"
)

type grpcServer struct {
}

func New() service.Service {
	return &grpcServer{}
}

func (a *grpcServer) Init() {
	log.Println("开始初始化全局服务")
}

func (a *grpcServer) Start() error {
	return nil
}

func (a *grpcServer) Stop() error {
	return nil
}

func (a *grpcServer) Name() string {
	return "manager grpc_server"
}
