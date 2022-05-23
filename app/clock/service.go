package clock

import (
	service "go_web_example/app"
	"log"
)

type Initialize struct {
}

func New() service.Service {
	return &Initialize{}
}

func (a *Initialize) Name() string {
	return "manager clock"
}

func (a *Initialize) Init() error {
	log.Println("开始初始化全局服务")
	return nil
}

func (a *Initialize) Start() error {
	return nil
}

func (a *Initialize) Stop() error {
	return nil
}
