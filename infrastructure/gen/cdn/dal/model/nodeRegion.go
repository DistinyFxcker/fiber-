package model

import "gorm.io/gorm"

//节点区域的绑定表
type NodeRegion struct {
	Id        string `gorm:"column:id;type:varchar(50);not null" json:"id"`                                          //id
	RegionId  string `gorm:"column:region_id;type:varchar(50);primary_key;not null" json:"regionId"`                 //区域ID
	NodeId    string `gorm:"column:node_id;type:varchar(50);primary_key;not null" json:"nodeId"`                     //节点ID
	Cname     string `gorm:"column:cname;type:varchar(50);"`                                                         //绑定的CNAME
	CreatedAt int64  `gorm:"column:created_at;type:int8;index:idx_node_region_created_at;not null" json:"createdAt"` //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;index:idx_node_region_updated_at;not null" json:"updatedAt"` //修改时间，新记录保持和创建时间一致
}

func (this *NodeRegion) TableName() string {
	return "node_region"
}

func (r *NodeRegion) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&NodeRegion{}); err != nil {
		panic(err)
	}
	return nil
}
