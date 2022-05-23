package cmd

import (
	"go_web_example/gen/dal/model"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func InitGen() {
	g := gen.NewGenerator(
		gen.Config{
			OutPath:      "gen/dal/query",
			ModelPkgPath: "gen/dal/model",
			WithUnitTest: true,
		})
	//g.UseDB(db)

	g.ApplyBasic(
		model.Address{},
		model.Admin{},
		model.CacheConfig{},
		model.Configure{},
		model.ConfigureTemplate{},
		model.DocManage{},
		model.Domain{},
		model.DomainAgreement{},
		model.DomainForward{},
		model.DomainSource{},
		model.DomainSourceConfig{},
		model.EmailRecord{},
		model.ErrorPageConfig{},
		model.HeaderConfig{},
		model.LoginRecord{},
		model.Node{},
		model.NodeConfig{},
		model.NodeIp{},
		model.NodeRegion{},
		model.PrimaryDomain{},
		model.Region{},
		model.ResourcesLabelValue{},
		model.ResourcesLabelBind{},
		model.ResourceLabel{},
		model.ServiceType{},
		model.System{},
		model.User{},
		model.UserGroup{},
		model.UserGroupRegion{},
		model.Version{})
	g.Execute()

}

type repo interface {
	Init(db *gorm.DB) error
}

func InitMigrate(db *gorm.DB) error {
	InitMap := []repo{
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.CacheConfig{},
		&model.Configure{},
		&model.ConfigureTemplate{},
		&model.DocManage{},
		&model.Domain{},
		&model.DomainAgreement{},
		&model.DomainForward{},
		&model.DomainSource{},
		&model.DomainSourceConfig{},
		&model.EmailRecord{},
		&model.ErrorPageConfig{},
		&model.HeaderConfig{},
		&model.LoginRecord{},
		&model.Node{},
		&model.NodeConfig{},
		&model.NodeIp{},
		&model.NodeRegion{},
		&model.PrimaryDomain{},
		&model.Region{},
		&model.ResourcesLabelValue{},
		&model.ResourcesLabelBind{},
		&model.ResourceLabel{},
		&model.ServiceType{},
		&model.System{},
		&model.User{},
		&model.UserGroup{},
		&model.UserGroupRegion{},
		&model.Version{},
	}
	for _, repo := range InitMap {
		if err := repo.Init(db); err != nil {
			return err
		}
	}
	return nil
}
