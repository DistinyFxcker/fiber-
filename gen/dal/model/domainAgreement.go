package model

import "gorm.io/gorm"

const (
	Domain_Agreement_http       = 1
	Domain_Agreement_https      = 2
	Domain_Agreement_http_https = 3
	Domain_Agreement_Custom     = 4

	Domain_Agreement_ProtocolType_http  = 1
	Domain_Agreement_ProtocolType_https = 2

	Domain_Agreement_skip_type_永久 = 1
	Domain_Agreement_skip_type_暂时 = 2
)

//域名跳转协议
type DomainAgreement struct {
	Id            string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	DomainId      string `gorm:"column:domain_id;type:varchar(50);not null;comment:域名ID"`  //域名ID
	InAgreement   int    `gorm:"column:in_agreement;type:int;comment:来访协议 1http 2https"`   //来访协议 1http 2https
	InPort        int    `gorm:"column:in_port;type:int;comment:来访端口"`                     //来访端口
	SkipType      int    `gorm:"column:skip_type;type:int;comment:跳转类型 1永久 2临时"`           //跳转类型 1永久 2临时
	SkipAgreement int    `gorm:"column:skip_agreement;type:int;comment:跳转协议 1http 2https"` //跳转协议 1http 2https
	SkipPort      int    `gorm:"column:skip_port;type:int;comment:跳转端口"`                   //跳转端口
	CreatedAt     int64  `gorm:"column:created_at;type:int8;not null"`                     //创建时间
	UpdatedAt     int64  `gorm:"column:updated_at;type:int8;not null"`                     //修改时间，新记录保持和创建时间一致
}

func (This *DomainAgreement) NewTable() string {
	return "cache_config"
}

func (r *DomainAgreement) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&DomainAgreement{}); err != nil {
		panic(err)
	}
	return nil
}
