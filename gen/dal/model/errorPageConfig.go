package model

import "gorm.io/gorm"

const (
	ErrorPage_Type_301 = 1
	ErrorPage_Type_302 = 2
	ErrorPage_Type_自定义 = 3
)

//错误页面
type ErrorPageConfig struct {
	Id           string `gorm:"column:id;type:varchar(50);primary_key;not null"`                          //id
	RelationId   string `gorm:"column:relation_id;type:varchar(50);not null;comment:关联主题ID"`              //关联主体ID
	RelationType int    `gorm:"column:relation_type;type:int;not null;comment:关联主体类型1节点 2区域 3域名"`         //关联主体类型1节点 2区域 3域名
	ConfigureId  string `gorm:"column:configure_id;type:varchar(50);not null;comment:配置ID"`               //配置ID
	Name         string `gorm:"column:name;type:varchar(50);not null;comment:名称"`                         //名称
	StatusCode   int    `gorm:"column:status_code;type:int2;not null;comment:状态码 类型定死 1 401-417 500-505"` //状态码 类型定死 1 401-417 500-505
	Status       int    `gorm:"column:status;type:int2;not null;default:1;comment:启用状态 1启用2关闭"`           //启用状态 1启用2关闭
	//类型
	Type int `gorm:"column:type;type:int2;comment:页面类型 1.301 2.302 3 自定义页面"` //页面类型 1.301 2.302 3 自定义页面
	//内容
	DestinationAddress string `gorm:"destination_address;type:varchar(255);comment:目标地址 当type = 1 OR 2"` //目标地址 当type = 1 OR 2
	CustomContext      string `gorm:"column:custom_context;type:text;comment:自定义页面"`                     //自定义页面

	CreatedAt int64 `gorm:"column:created_at;type:int8;index:idx_domain_error_page_config_created_at;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;index:idx_domain_error_page_config_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *ErrorPageConfig) NewTable() string {
	return "error_page_config"
}

func (r *ErrorPageConfig) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&ErrorPageConfig{}); err != nil {
		panic(err)
	}
	return nil
}
