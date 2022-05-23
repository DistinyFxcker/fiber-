package model

import "gorm.io/gorm"

const (
	NodeIp_Type_WithIn  = 1 //内网IP
	NodeIp_Type_OutSide = 2 //外部IP
	NodeIp_Is_master    = 1 //主IP
	NodeIp_Not_master   = 2 //非主IP
)

type NodeIp struct {
	Id        string `gorm:"column:id;type:varchar(50);primary_key;not null;comment:"`
	Ip        string `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	NodeId    string `gorm:"column:node_id;type:varchar(50);not null;comment:"`
	Type      int    `gorm:"column:type;type:int2;comment:类型 1内网IP 2外网IP"` //类型
	IsMaster  int    `gorm:"column:is_master;type:int2;comment:是否是主IP"`    //是否是主IP
	CreatedAt int64  `gorm:"column:created_at;type:int8;not null"`         //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;not null"`         //修改时间，新记录保持和创建时间一致
}

func (This *NodeIp) NewTable() string {
	return "node_ip"
}

func (r *NodeIp) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&NodeIp{}); err != nil {
		panic(err)
	}
	return nil
}
