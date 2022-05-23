package grpcClient

import (
	service "go_web_example/app"
)

type grpcClient struct {
}

func New() service.Service {
	return &grpcClient{}
}

func (a *grpcClient) Name() string {
	return "manager grpcClient"
}

func (a *grpcClient) Init() error {
	//log.Println("开始初始化全局服务")
	return nil
}

func (a *grpcClient) Start() error {
	return nil
}

func (a *grpcClient) Stop() error {
	return nil
}
