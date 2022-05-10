package clock

import (
	service "cdnFiber/app"
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

func (a *Initialize) Init() {
	log.Println("开始初始化全局服务")
}

func (a *Initialize) Start() error {
	return nil
}

func (a *Initialize) Stop() error {
	return nil
}
