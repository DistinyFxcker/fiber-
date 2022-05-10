package model

import "gorm.io/gorm"

//键值
type ResourcesLabelValue struct {
	Id        string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	KeyId     string `gorm:"column:key_id;type:varchar(50);not null;comment:标签键ID"`     //标签键ID
	KeyName   string `gorm:"column:key_name;type:varchar(255);not null;comment:冗余的键名称"` //冗余的键名称
	Value     string `gorm:"column:value;type:varchar(50);not null;comment:标签值"`        //标签值
	CreateId  string `gorm:"column:create_id;type:varchar(50);not null;comment:创建人"`    //创建人
	CreatedAt int64  `gorm:"column:created_at;type:int8;index:;not null;comment:"`      //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;index:;not null;comment:"`      //修改时间，新记录保持和创建时间一致
}

func (This *ResourcesLabelValue) NewTable() string {
	return "cache_config"
}

func (r *ResourcesLabelValue) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&ResourcesLabelValue{}); err != nil {
		panic(err)
	}
	return nil
}
