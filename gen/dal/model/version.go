package model

import "gorm.io/gorm"

//版本管理
type Version struct {
	Id         uint64  `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT"`                     //id
	Newest     int     `gorm:"column:newest;type:int;not null;comment:最新文件 1位最新"`                 //最新文件，1为最新
	FileName   string  `gorm:"column:file_name;type:varchar(50);not null;comment:文件名"`            //文件名
	FileSize   float64 `gorm:"column:file_size;type:float;not null;comment:文件体积"`                 //文件体积(单位byte)
	FileDir    string  `gorm:"column:file_dir;type:varchar(100);not null;comment:文件路径"`           //文件路径
	VersionNum string  `gorm:"column:version_num;type:varchar(50);not null;comment:版本号"`          //版本号
	CreatedAt  int64   `gorm:"column:created_at;type:int8;index:idx_version_created_at;not null"` //创建时间
	UpdatedAt  int64   `gorm:"column:updated_at;type:int8;index:idx_version_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (this *Version) TableName() string {
	return "version"
}

func (m *Version) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&Version{}); err != nil {
		panic(err)
	}
	return nil
}
