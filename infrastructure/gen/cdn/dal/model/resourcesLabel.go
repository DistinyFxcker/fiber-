package model

import "gorm.io/gorm"

type ResourceLabel struct {
	Id        string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	Key       string `gorm:"column:key;type:varchar(50);not null;not null;comment:标签名称"`
	CreateId  string `gorm:"column:create_id;type:varchar(50);not null;comment:创建用户ID"`
	CreatedAt int64  `gorm:"column:created_at;type:int8;not null"` //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *ResourceLabel) NewTable() string {
	return "resource_label"
}

func (r *ResourceLabel) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&ResourceLabel{}); err != nil {
		panic(err)
	}
	return nil
}
