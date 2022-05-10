package user

import (
	service "cdnFiber/app"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	Fiber      *fiber.App
	ListenPort string
}

func New(_listenPort string) service.Service {
	return &UserService{
		ListenPort: _listenPort,
		Fiber: fiber.New(fiber.Config{
			AppName: "cdn_user",
		}),
	}
}

func (a *UserService) Init() {
}

func (a *UserService) Start() error {
	a.Fiber.Listen(":" + a.ListenPort)
	return nil
}
func (a *UserService) Stop() error {
	a.Fiber.Shutdown()
	return nil
}

func (a *UserService) Name() string {
	return "manager cdn_user"
}
