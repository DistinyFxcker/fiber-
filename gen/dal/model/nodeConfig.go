package model

import "gorm.io/gorm"

type NodeConfig struct {
	Id           string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	TemplateName string `gorm:"column:template_name;type:varchar(50);comment:模板名称"` //模板名称
	//内存缓存占比
	MemCacheProportion int `gorm:"column:mem_cache_proportion;type:int8;comment://内存站比(1-100 百分比)"` //内存站比(1-100 百分比)

	//磁盘缓存设置
	CacheType int    `gorm:"column:cache_type;type:int;default:1;comment:磁盘储存类型 1分区 2目录"` //磁盘储存类型 1分区 2目录
	Path      string `gorm:"column:path;type:varchar(128);comment:磁盘目录路径"`                //磁盘目录路径
	//分区
	PartitionPath string `gorm:"column:partition_path;type:varchar(50);comment:分区路径"` //分区路径
	//磁盘缓存占比(如果设置了分区路径默认为100)
	DiskCacheProportion int `gorm:"column:disk_cache_proportion;type:int8;comment:磁盘缓存占比"` //磁盘缓存占比

	CreatedAt int64 `gorm:"column:created_at;type:int8;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *NodeConfig) NewTable() string {
	return "node_config"
}

func (r *NodeConfig) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&NodeConfig{}); err != nil {
		panic(err)
	}
	return nil
}
