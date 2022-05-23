package model

import "gorm.io/gorm"

type UserGroup struct {
	Id           string `gorm:"column:id;type:varchar(50);not null;primary_key"`
	Name         string `gorm:"column:name;type:varchar(50);not null;;comment:分组名称"`
	Description  string `gorm:"column:description;type:varchar(255);;comment:分组描述"`
	IsDefault    int    `gorm:"column:is_default;type:int;default:2;comment:是否默认分组 1默认"`   //是否为默认分组
	RegionSwitch int    `gorm:"column:region_switch;type:int2;default:2;comment:是否默认开启区域"` //是否开启默认区域
	CreatedAt    int64  `gorm:"column:created_at;type:int8;not null"`                      //创建时间
	UpdatedAt    int64  `gorm:"column:updated_at;type:int8;not null"`                      //修改时间，新记录保持和创建时间一致

	RegionIds []string `gorm:"-;comment:"` //绑定的用户ID
}

func (This *UserGroup) NewTable() string {
	return "user_group"
}

func (m *UserGroup) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserGroup{}); err != nil {
		panic(err)
	}
	return nil
}
