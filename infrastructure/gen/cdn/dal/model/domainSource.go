package model

import "gorm.io/gorm"

type DomainSource struct {
	Id           string `gorm:"column:id;type:varchar(50);primary_key;not null"`             //id
	DomainId     string `gorm:"column:domain_id;type:varchar(50);not null;comment:域名ID"`     //域名ID
	CurrenNumber int    `gorm:"column:curren_number;type:int;default:0;comment:是否是主备 1主 2备"` //是否是主备 1主 2备
	//Type     int    `gorm:"column:type;type:int2;not null;comment:"`                 //协议类型 1.IPV4-公网IP 2.IPV4-内网IP  3.域名-http-域名 4.域名-http-端口 5.域名-https-域名 6.域名-https-端口
	Context string `gorm:"column:context;type:varchar(50);not null;comment:协议内容 域名 IP地址"` //协议内容 域名 IP地址

	Port         int `gorm:"column:port;type:int;not null;comment:端口"` //端口
	SerialNumber int `gorm:"column:serial_number;type:int;comment:权重"` //权重

	CreatedAt int64 `gorm:"column:created_at;type:int8;index:idx_domain_source_created_at;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;index:idx_domain_source_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *DomainSource) NewTable() string {
	return "domain_source"
}

func (r *DomainSource) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&DomainSource{}); err != nil {
		panic(err)
	}
	return nil
}
