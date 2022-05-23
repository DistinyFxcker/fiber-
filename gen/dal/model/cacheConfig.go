package model

import "gorm.io/gorm"

const (
	CacheConfig_CacheType_文件后缀名 = 1
	CacheConfig_CacheType_目录路径  = 2
	CacheConfig_CacheType_全路径   = 3
	CacheConfig_CacheType_缓存首页  = 4

	CacheConfig_CacheTimeUnit_小时 = 1
	CacheConfig_CacheTimeUnit_分钟 = 2
	CacheConfig_CacheTimeUnit_天  = 3

	CacheConfig_RelationType_node   = 1
	CacheConfig_RelationType_region = 2
	CacheConfig_RelationType_domain = 3
	CacheConfig_RelationType_model  = 4
)

type CacheConfig struct {
	Id           string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	ConfigureId  string `gorm:"column:configure_id;type:varchar(50);comment:配置Id"`           //配置ID
	RelationId   string `gorm:"column:relation_id;type:varchar(50);not null;comment:配置所属ID"` //配置所属ID
	RelationType int    `gorm:"column:relation_type;type:int;not null;comment:配置所属类型"`       //配置所属类型 1节点 2区域 3域名 4业务类型
	Name         string `gorm:"column:name;type:varchar(50);comment:规则名称"`                   //规则名称
	Description  string `gorm:"column:description;type:text;comment:描述"`                     //描述
	//类型
	Type        int    `gorm:"column:type;type:int;default:1;comment:缓存类型 1文件后缀名 2目录路径 3全路径 4缓存首页"` //缓存类型 1文件后缀名 2目录路径 3全路径 4缓存首页
	TypeContext string `gorm:"column:type_context;type:text;comment:缓存内容"`                          //缓存内容

	PriorityLevel int `gorm:"column:priority_level;type:int;default:1;comment:优先级"` //优先级

	CacheTimeV2 string `gorm:"column:cache_time_v2;type:varchar(50);default:1d;comment:缓存时间"` //缓存时间

	CreatedAt int64 `gorm:"column:created_at;type:int8;index:idx_cache_config_created_at;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;index:idx_cache_config_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *CacheConfig) NewTable() string {
	return "cache_config"
}

func (r *CacheConfig) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&CacheConfig{}); err != nil {
		panic(err)
	}
	return nil
}
