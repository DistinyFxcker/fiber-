package model

import "gorm.io/gorm"

type HeaderConfig struct {
	Id        string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	DomainId  string `gorm:"column:domain_id;type:varchar(50);not null"`  //绑定域名ID
	Parameter string `gorm:"column:parameter;not null;type:varchar(128)"` //相应头配置参数
	Value     string `gorm:"column:value;type:varchar(50)"`               //响应头取值
	CreatedAt int64  `gorm:"column:created_at;type:int8;not null"`        //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;not null"`        //修改时间，新记录保持和创建时间一致
}

func (This *HeaderConfig) NewTable() string {
	return "header_config"
}

func (r *HeaderConfig) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&HeaderConfig{}); err != nil {
		panic(err)
	}
	return nil
}
