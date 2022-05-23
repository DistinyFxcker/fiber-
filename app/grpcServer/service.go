package grpcServer

import (
	service "go_web_example/app"
	"log"
)

type grpcServer struct {
}

func New() service.Service {
	return &grpcServer{}
}

func (a *grpcServer) Init() error {
	log.Println("开始初始化全局服务")
	return nil
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
