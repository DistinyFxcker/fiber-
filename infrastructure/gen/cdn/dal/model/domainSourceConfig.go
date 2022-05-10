package model

import "gorm.io/gorm"

//回源请求头
type DomainSourceConfig struct {
	Id        string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	DomainId  string `gorm:"column:domain_id;type:varchar(50);not null;comment:域名ID"`     //域名ID
	HostName  string `gorm:"column:host_name;type:varchar(50);not null;comment:回源HOST域名"` //回源HOST域名
	Header    string `gorm:"column:header;type:text;comment:回源请求头"`                       //回源请求头
	CacheTime int    `gorm:"column:cache_time;type:int;not null;comment:预缓存时间(单位:小时)"`    //预缓存时间(单位:小时)

	HeaderBack []Header `gorm:"-"`
	CreatedAt  int64    `gorm:"column:created_at;type:int8;index:idx_domain_source_config_created_at;not null"` //创建时间
	UpdatedAt  int64    `gorm:"column:updated_at;type:int8;index:idx_domain_source_config_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

type Header struct {
	Header string //header头
	Value  string //header值
	Type   int    //1添加 2删除
}

func (This *DomainSourceConfig) NewTable() string {
	return "cache_config"
}

func (r *DomainSourceConfig) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&DomainSourceConfig{}); err != nil {
		panic(err)
	}
	return nil
}
