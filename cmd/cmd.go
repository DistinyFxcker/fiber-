package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_web_example/app"
	"go_web_example/app/admin"
	"go_web_example/app/clock"
	"go_web_example/app/grpcClient"
	"go_web_example/app/grpcServer"
	"go_web_example/app/user"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var (
	ZapLog *zap.SugaredLogger
)

type serviceManage struct {
	errGroup errgroup.Group
	services []service.Service
}

func New(db *gorm.DB) service.Service {
	_serviceManage := &serviceManage{
		services: make([]service.Service, 0),
		errGroup: errgroup.Group{},
	}
	//
	if err := _serviceManage.addService(db); err != nil {
		panic(err.Error())
	}
	return _serviceManage
}

func (this *serviceManage) Name() string {
	return "service manage"
}

func (this *serviceManage) Init() error {
	for _, service := range this.services {
		err := service.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *serviceManage) Start() error {
	for i := range this.services {
		index := i //临时变量
		this.errGroup.Go(func() error {
			return this.services[index].Start()
		})
	}
	return this.errGroup.Wait()
}

func (this *serviceManage) Stop() error {
	for _, service := range this.services {
		this.errGroup.Go(func() error {
			return service.Stop()
		})
	}
	return this.errGroup.Wait()
}

//go:generate swag init --instanceName=CdnAdmin --dir=../app/admin --generalInfo=adminService.go  --output ../docs/cdnAdmin -struct-tags dc
//go:generate swag init --instanceName=CdnUser --dir=../app/user  --generalInfo=userService.go --output ../docs/cdnUser
//添加服务
func (this *serviceManage) addService(db *gorm.DB) error {
	//cdnAdmin服务
	{
		_apiService := admin.New(db)
		if err := this.register(_apiService); err != nil {
			return err
		}
	}
	//cdnUser服务
	{
		_apiUserService := user.New(viper.GetString("cdnUser.APIPort"))
		if err := this.register(_apiUserService); err != nil {
			return err
		}
	}
	//定时器服务
	{
		_clockService := clock.New()
		if err := this.register(_clockService); err != nil {
			return err
		}
	}
	//grpc服务端
	{
		_grpcServer := grpcServer.New()
		if err := this.register(_grpcServer); err != nil {
			return err
		}
	}
	//grpc
	{
		_grpcClient := grpcClient.New()
		if err := this.register(_grpcClient); err != nil {
			return err
		}
	}

	return nil
}

//注册
func (this *serviceManage) register(_service service.Service) error {
	for _, service := range this.services {
		if service.Name() == _service.Name() {
			return errors.New(fmt.Sprintf("服务%s,已存在", _service.Name()))
		}
	}
	this.services = append(this.services, _service)
	return nil
}
