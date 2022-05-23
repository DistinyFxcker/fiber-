package admin

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	service "go_web_example/app"
	"go_web_example/app/admin/controller"
	"go_web_example/app/admin/global"
	"go_web_example/app/admin/middleware"
	_ "go_web_example/docs/cdnAdmin"
	"go_web_example/infrastructure/zapLog"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AdminService struct {
	engine     *gin.Engine
	ListenPort string
	SecretKey  string
}

func New(db *gorm.DB) service.Service {
	global.DB = db
	return &AdminService{
		engine: gin.New(),
	}
}

func Name() string {
	return "manager cdn_service"
}

func (a *AdminService) Init() error {
	var (
		err error
	)

	err = a.setConfig()
	if err != nil {
		return err
	}

	a.setMode()
	err = a.setLog()
	if err != nil {
		return err
	}
	a.setMiddlewares()
	a.setRoute()
	return nil
}

func (a *AdminService) Name() string {
	return "manager cdn_admin"
}

//设置模式
func (this *AdminService) setMode() {
	gin.SetMode(global.Viper.GetString("cdnAdmin.Mode"))
}

func (a *AdminService) Start() error {
	log.Println("cdnAdmin服务启动")
	global.ZapLog.Info("cdnAdmin服务启动")
	pprof.Register(a.engine)
	return a.engine.Run(":" + global.Viper.GetString("cdnAdmin.APIPort"))
}
func (a *AdminService) Stop() error {
	global.ZapLog.Info("cdnAdmin api服务停止")
	return nil
}

func (a *AdminService) setMiddlewares() {
	a.engine.Use(gin.Logger())                       //日志
	a.engine.Use(gin.Recovery())                     //错误信息回调
	a.engine.Use(gzip.Gzip(gzip.DefaultCompression)) //压缩设置
	a.engine.Use(middleware.MiddlewareCors())        //跨域设置
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:10001
// @BasePath /api/v1
func (a *AdminService) setRoute() {
	a.engine.LoadHTMLFiles("./dist/index.html")
	a.engine.StaticFS("/static", http.Dir("./dist/static"))
	a.engine.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	jwt := jwtware.New(jwtware.Config{
		SigningKey: a.SecretKey,
	})

	//使用跟ginSwagger类似的swagger包
	a.engine.GET("/swagger/*", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "CdnAdmin"
	}))

	version := a.engine.Group("/api/v1")
	version.POST("auth", controller.Auth.Auth)
	users := version.Group("users", jwt)
	{
		//登出
		users.POST("logout", controller.Auth.LogOut)
		//刷新token
		users.POST("refresh", controller.Auth.RefreshToken)

	}

}

func (a *AdminService) setLog() error {
	var (
		err error
	)

	log := zapLog.ZapLog{
		Path: global.Viper.GetString("log.Path"),
		Day:  global.Viper.GetDuration("log.Day"),
		Hour: global.Viper.GetDuration("log.Hour"),
	}
	global.ZapLog, err = log.InitLogger()
	if err != nil {
		return err
	}
	return nil
}

//初始化全局配置
func (a *AdminService) setConfig() (err error) {
	global.Viper = viper.New()
	global.Viper.SetConfigName("admin")
	global.Viper.SetConfigType("ini")
	global.Viper.AddConfigPath("./config")

	global.Viper.SetDefault("cdnAdmin", map[string]string{
		"Mode":                "debug",
		"APIPort":             "8015",
		"CertificateCallback": "/api/v1/upload/certificate/cdn",
		"SecretKey":           "text@163.com",
	})

	global.Viper.SetDefault("log", map[string]string{
		"Path": "log/cdnAdmin/cdnAdmin.log",
		"Day":  "90",
		"Hour": "24",
	})

	err = global.Viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
