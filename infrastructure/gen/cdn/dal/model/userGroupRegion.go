package model

import "gorm.io/gorm"

type UserGroupRegion struct {
	Id          string `gorm:"column:id;type:varchar(50);not null"`
	UserGroupId string `gorm:"column:user_group_id;type:varchar(50);not null;comment:用户分组ID"`
	RegionId    string `gorm:"column:region_id;type:varchar(50);not null;comment:区域ID"`
	CreatedAt   int64  `gorm:"column:created_at;type:int8;not null"`
	UpdatedAt   int64  `gorm:"column:updated_at;type:int8;not null"`
}

func (this *UserGroupRegion) TableName() string {
	return "user_group_region"
}

func (m *UserGroupRegion) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserGroupRegion{}); err != nil {
		panic(err)
	}
	return nil
}
