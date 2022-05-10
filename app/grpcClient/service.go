package grpcClient

import (
	service "cdnFiber/app"
	"log"
)

type grpcClient struct {
}

func New() service.Service {
	return &grpcClient{}
}

func (a *grpcClient) Name() string {
	return "manager grpcClient"
}

func (a *grpcClient) Init() {
	log.Println("开始初始化全局服务")
}

func (a *grpcClient) Start() error {
	return nil
}

func (a *grpcClient) Stop() error {
	return nil
}
