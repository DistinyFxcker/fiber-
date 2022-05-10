package model

import "gorm.io/gorm"

//区域列表
type Region struct {
	Id           string `gorm:"column:id;type:varchar(50);not null"`
	Name         string `gorm:"column:name;type:varchar(50);not null;comment:区域名称"`   //区域名称
	MasterDomain string `gorm:"column:master_domain;type:varchar(128);comment:主域名"`   //主域名(手动添加,跟站点所用的主域名区分开)
	Cname        string `gorm:"column:c_name;type:varchar(50);comment:根据主域名创建的Cname"` //根据主域名创建的Cname
	ConnectTime  int    `gorm:"column:connect_time;type:int;default:30;comment:连接时间"` //连接时间
	CreatedAt    int64  `gorm:"column:created_at;type:int8;not null"`
	UpdatedAt    int64  `gorm:"column:updated_at;type:int8;not null"`
	//ConfigureId  string `gorm:"-;comment:"`  //配置ID
	//NodeNumber   int    `gorm:"-;comment:"`   //关联节点数量
	//DomainNumber int    `gorm:"-;comment:"` //域名数量
}

func (this *Region) TableName() string {
	return "region"
}

func (r *Region) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&Region{}); err != nil {
		panic(err)
	}
	return nil
}
