package model

import "gorm.io/gorm"

//系统设置
type System struct {
	Id    int    `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT"` //id
	Key   string `gorm:"column:key;type:varchar(100);not null"`         //key
	Value string `gorm:"column:value;type:varchar(512);not null"`       //value

}

func (this *System) TableName() string {
	return "system"
}

func (r *System) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&System{}); err != nil {
		panic(err)
	}
	return nil
}
