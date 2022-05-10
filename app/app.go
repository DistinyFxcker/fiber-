package service

type Service interface {
	Init()
	Start() error
	Stop() error
	Name() string
}
