package cdn

import (
	"cdnFiber/infrastructure/gen/cdn/dal/model"
	"gorm.io/gorm"
)

type cdnDao struct {
	Address            model.Address
	Admin              model.Admin
	CacheConfig        model.CacheConfig
	Configure          model.Configure
	ConfigureTemplate  model.ConfigureTemplate
	DocManager         model.DocManage
	Domain             model.Domain
	DomainAgreement    model.DomainAgreement
	DomainForward      model.DomainForward
	DomainSource       model.DomainSource
	DomainSourceConfig model.DomainSourceConfig
	EmailRecord        model.EmailRecord
	ErrorPageConfig    model.ErrorPageConfig
	HeaderConfig       model.HeaderConfig
	LoginRecord        model.LoginRecord
	Node               model.Node
	NodeConfig         model.NodeConfig
	NodeIp             model.NodeIp
	NodeRegion         model.NodeRegion
	PrimaryDomain      model.PrimaryDomain
	Region             model.Region
	ResourceLabelValue model.ResourcesLabelValue
	ResourcesLabBind   model.ResourcesLabelBind
	ResourcesLabel     model.ResourceLabel
	ServiceType        model.ServiceType
	System             model.System
	User               model.User
	UserGroup          model.UserGroup
	UserGroupRegion    model.UserGroupRegion
	Version            model.Version
}

type repo interface {
	Init(db *gorm.DB) error
}

func New(db *gorm.DB) (*cdnDao, error) {
	//自动迁移和初始化数据
	Dao := &cdnDao{
		Address:            model.Address{},
		Admin:              model.Admin{},
		CacheConfig:        model.CacheConfig{},
		Configure:          model.Configure{},
		ConfigureTemplate:  model.ConfigureTemplate{},
		DocManager:         model.DocManage{},
		Domain:             model.Domain{},
		DomainAgreement:    model.DomainAgreement{},
		DomainForward:      model.DomainForward{},
		DomainSource:       model.DomainSource{},
		DomainSourceConfig: model.DomainSourceConfig{},
		EmailRecord:        model.EmailRecord{},
		ErrorPageConfig:    model.ErrorPageConfig{},
		HeaderConfig:       model.HeaderConfig{},
		LoginRecord:        model.LoginRecord{},
		Node:               model.Node{},
		NodeConfig:         model.NodeConfig{},
		NodeIp:             model.NodeIp{},
		NodeRegion:         model.NodeRegion{},
		PrimaryDomain:      model.PrimaryDomain{},
		Region:             model.Region{},
		ResourceLabelValue: model.ResourcesLabelValue{},
		ResourcesLabBind:   model.ResourcesLabelBind{},
		ResourcesLabel:     model.ResourceLabel{},
		ServiceType:        model.ServiceType{},
		System:             model.System{},
		User:               model.User{},
		UserGroup:          model.UserGroup{},
		UserGroupRegion:    model.UserGroupRegion{},
		Version:            model.Version{},
	}
	InitMap := []repo{
		&Dao.User,
		&Dao.Address,
		&Dao.Admin,
		&Dao.CacheConfig,
		&Dao.Configure,
		&Dao.ConfigureTemplate,
		&Dao.DocManager,
		&Dao.Domain,
		&Dao.DomainAgreement,
		&Dao.DomainForward,
		&Dao.DomainSource,
		&Dao.DomainSourceConfig,
		&Dao.EmailRecord,
		&Dao.ErrorPageConfig,
		&Dao.HeaderConfig,
		&Dao.LoginRecord,
		&Dao.Node,
		&Dao.NodeConfig,
		&Dao.NodeIp,
		&Dao.NodeRegion,
		&Dao.PrimaryDomain,
		&Dao.Region,
		&Dao.ResourceLabelValue,
		&Dao.ResourcesLabBind,
		&Dao.ResourcesLabel,
		&Dao.ServiceType,
		&Dao.System,
		&Dao.User,
		&Dao.UserGroup,
		&Dao.UserGroupRegion,
		&Dao.Version,
	}
	for _, repo := range InitMap {
		if err := repo.Init(db); err != nil {
			return nil, err
		}
	}
	return Dao, nil
}
