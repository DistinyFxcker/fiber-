package model

import "gorm.io/gorm"

//服务类型
type ServiceType struct {
	Id        string `gorm:"column:id;type:varchar(50);primary_key;not null"`
	Name      string `gorm:"column:name;type:varchar(50); not null"`
	CreatedAt int64  `gorm:"column:created_at;type:int8;index:;not null"` //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;index:;not null"` //修改时间，新记录保持和创建时间一致

	CacheNum int64 `gorm:"-"`
}

func (This *ServiceType) NewTable() string {
	return "service_type"
}

func (r *ServiceType) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&ServiceType{}); err != nil {
		panic(err)
	}
	return nil
}
