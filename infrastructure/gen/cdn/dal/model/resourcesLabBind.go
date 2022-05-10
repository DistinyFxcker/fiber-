package model

import "gorm.io/gorm"

//标签绑定关系
type ResourcesLabelBind struct {
	Id           string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	RelationId   string `gorm:"column:relation_id;type:varchar(50);not null;comment:绑定资源ID"`      //绑定资源ID
	RelationType int8   `gorm:"column:relation_type;type:int;not null;default:1;comment:绑定资源类型"`  //绑定资源类型 1域名 后面可能会加上资源类型
	LabelValueId string `gorm:"column:label_value_id;type:varchar(50);not null;comment:绑定的标签纸ID"` //绑定的标签值ID
	KeyId        string `gorm:"column:key_id;type:varchar(50);not null;comment:绑定的标签键ID"`         //绑定的标签键ID
	Value        string `gorm:"column:value;type:varchar(50);not null;comment:标签值冗余"`             //标签值冗余
	Key          string `gorm:"column:key;type:varchar(50);not null;comment:标签键冗余"`               //标签键冗余
	CreateId     string `gorm:"column:create_id;type:varchar(50);not null;comment:绑定归属用户"`        //绑定归属用户
	CreatedAt    int64  `gorm:"column:created_at;type:int8;index:;not null"`                      //创建时间
	UpdatedAt    int64  `gorm:"column:updated_at;type:int8;index:;not null"`                      //修改时间，新记录保持和创建时间一致
}

func (This *ResourcesLabelBind) NewTable() string {
	return "resources_label_bind"
}

func (r *ResourcesLabelBind) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&ResourcesLabelBind{}); err != nil {
		panic(err)
	}
	return nil
}
