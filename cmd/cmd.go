package cmd

import (
	"cdnFiber/app"
	"cdnFiber/app/admin"
	"cdnFiber/app/clock"
	"cdnFiber/app/grpcClient"
	"cdnFiber/app/grpcServer"
	"cdnFiber/app/initialize"
	"cdnFiber/app/user"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

type serviceManage struct {
	errGroup errgroup.Group
	services []service.Service
}

func New() service.Service {
	_serviceManage := &serviceManage{
		services: make([]service.Service, 0),
		errGroup: errgroup.Group{},
	}
	//
	if err := _serviceManage.addService(); err != nil {
		panic(err.Error())
	}
	return _serviceManage
}

func (this *serviceManage) Name() string {
	return "service manage"
}

func (this *serviceManage) Init() {
	for _, service := range this.services {
		service.Init()
	}
}

func (this *serviceManage) Start() error {
	for i := range this.services {
		index := i //临时变量
		this.errGroup.Go(func() error {
			return this.services[index].Start()
		})
	}

	//会阻塞
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

//go:generate swag init --instanceName=cdnAdmin --dir=../app/admin --generalInfo=service.go  --output ../docs/cdnAdmin
//go:generate swag init --instanceName=cdnUser --dir=../app/user  --generalInfo=service.go --output ../docs/cdnUser
//添加服务
func (this *serviceManage) addService() error {
	//init服务
	{
		_initService := initialize.New()
		if err := this.register(_initService); err != nil {
			return err
		}
	}
	//cdnAdmin服务
	{
		_apiService := admin.New(viper.GetString("cdnAdmin.APIPort"), viper.GetString("cdnAdmin.FiberMode"), viper.GetString("cdnAdmin.SecretKey"))
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
