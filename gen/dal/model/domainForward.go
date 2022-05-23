package model

import "gorm.io/gorm"

type DomainForward struct {
	Id            string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	DomainId      string `gorm:"column:domain_id;type:varchar(50);not null;comment:所属域名ID"` //所属域名ID
	Address       string `gorm:"column:address;type:varchar(128);comment:域名"`               //域名
	TargetAddress string `gorm:"column:target_address;type:varchar(156);comment:目标地址"`      //目标地址

	CreatedAt int64 `gorm:"column:created_at;type:int8;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *DomainForward) NewTable() string {
	return "domain_forward"
}

func (r *DomainForward) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&DomainForward{}); err != nil {
		panic(err)
	}
	return nil
}
