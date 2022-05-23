package model

import "gorm.io/gorm"

//文档管理
type DocManage struct {
	Id        uint64 `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT"`                        //id
	Title     string `gorm:"column:title;type:varchar(100);not null;comment:标题"`                   //标题
	Content   string `gorm:"column:content;type:text;not null;comment:内容"`                         //内容
	Reading   uint64 `gorm:"column:reading;type:int;not null;comment:阅读数"`                         //阅读数
	Type      int    `gorm:"column:type;type:int;not null;comment:1常见问题 2使用教程"`                    //1,常见问题，2使用教程
	Label     string `gorm:"column:label;type:varchar(50);not null;comment:标签"`                    //标签
	CreatedAt int64  `gorm:"column:created_at;type:int8;index:idx_doc_manage_created_at;not null"` //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;index:idx_doc_manage_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (this *DocManage) TableName() string {
	return "doc_manage"
}

func (r *DocManage) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&DocManage{}); err != nil {
		panic(err)
	}
	return nil
}
