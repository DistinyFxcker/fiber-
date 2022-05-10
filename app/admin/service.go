package admin

import (
	service "cdnFiber/app"
	"cdnFiber/app/admin/controller"
	_ "cdnFiber/docs/cdnAdmin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/swaggo/fiber-swagger"
)

type AdminService struct {
	Fiber      *fiber.App
	Mode       string
	ListenPort string
	SecretKey  string
}

func New(_port, mode, secretKey string) service.Service {
	return &AdminService{
		ListenPort: _port,
		Fiber: fiber.New(fiber.Config{
			AppName: "cdn_admin",
		}),
		Mode:      mode,
		SecretKey: secretKey,
	}
}

func Name() string {
	return "manager cdn_service"
}

func (a *AdminService) Init() {
	a.SetMiddlewares()
	a.SetRoute()
}

func (a *AdminService) Name() string {
	return "manager cdn_admin"
}

func (a *AdminService) Start() error {
	if err := a.Fiber.Listen(":" + a.ListenPort); err != nil {
		return err
	}
	return nil
}
func (a *AdminService) Stop() error {
	a.Fiber.Shutdown()
	return nil
}

func (a *AdminService) SetMiddlewares() {
	a.Fiber.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE,UPDATE",
		AllowHeaders:     "Authorization,userID,token, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
		MaxAge:           172800,
	}),
		recover2.New(),
		compress.New(compress.Config{
			Level: compress.LevelDisabled,
		}),
		logger.New(logger.Config{
			Format:     "${pid} ${locals:requestid} - cdnAdmin - ${status} - ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "Asia/Shanghai",
		}),
	)
}

// @title cdnAdmin的swagger文档
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /v2
func (a *AdminService) SetRoute() {
	jwt := jwtware.New(jwtware.Config{
		SigningKey: a.SecretKey,
	})
	//使用跟ginSwagger类似的swagger包
	a.Fiber.Get("/swagger/*", fiberSwagger.FiberWrapHandler(func(c *fiberSwagger.Config) {
		c.InstanceName = "cdnAdmin"
	}))

	version := a.Fiber.Group("/api/v1")
	version.Post("auth", controller.Auth.Auth)
	users := version.Group("users", jwt)
	{
		//登出
		users.Post("logout", controller.Auth.LogOut)
		//刷新token
		users.Post("refresh", controller.Auth.RefreshToken)
	}

	//:= version.
}
