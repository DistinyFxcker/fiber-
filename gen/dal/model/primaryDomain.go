package model

import "gorm.io/gorm"

//系统-主域名
type PrimaryDomain struct {
	Id        string `gorm:"column:id;type:varchar(50);not null"`
	Domain    string `gorm:"column:domain;type:varchar(50);not null;"`
	CreatedAt int64  `gorm:"column:created_at;type:int8;not null"`
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;not null"`
}

func (this *PrimaryDomain) TableName() string {
	return "primary_domain"
}

func (r *PrimaryDomain) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&PrimaryDomain{}); err != nil {
		panic(err)
	}
	return nil
}
