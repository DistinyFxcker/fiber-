package model

import "gorm.io/gorm"

const (
	Configure_Node_type   = 1 //TODO:节点现在已经不带有公共配置,可取消
	Configure_Region_type = 2
	Configure_Domain_type = 3

	//默认使用区域的缓存配置
	Configure_CacheType_default = 1
	//使用自定义的缓存配置
	Configure_CacheTyoe_custom = 2
)

type Configure struct {
	Id           string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	RelationId   string `gorm:"column:relation_id;type:varchar(50);not null;comment:关联ID"`                       //关联ID(节点ID 或者 区域ID)
	RelationType int    `gorm:"column:relation_type;type:int2;not null;default:1;comment:关联节点类型(1节点 2区域ID 3域名)"` //关联节点类型(1节点 2区域ID 3域名)
	//超时设置
	ConnectTime int `gorm:"column:connect_time;type:int;default:30;comment:连接超时时间"` //连接超时时间(单位:秒)
	//缓存规则
	MinConnNum       int `gorm:"column:min_conn_num;type:int;default:1;comment:最小请求连接数"`              //最小请求连接数
	MaxMemoryCache   int `gorm:"column:max_memory_cache;type:int;comment:最大内存缓存"`                     //最大内存缓存(单位M)
	MaxDiskCache     int `gorm:"column:max_disk_cache;type:int;comment:磁盘缓存大小"`                       //磁盘缓存大小(单位M)
	InteMaxDiskCache int `gorm:"column:inte_max_disk_cache;type:int;comment:只能最大磁盘缓存"`                //智能最大磁盘缓存(单位M)
	DiskCacheType    int `gorm:"column:disk_cache_type;type:int;default:1;comment:最大磁盘缓存类型(1普通 2智能)"` //最大磁盘缓存类型(1普通 2智能)

	//压缩规则
	CompressionGrade     int    `gorm:"column:compression_grade;type:int;default:6;comment:压缩等级"`                   //压缩等级
	CompressionType      string `gorm:"column:compression_type;type:text;comment:压缩类型"`                             //压缩类型
	CompressionSizeUpper int    `gorm:"column:compression_size_upper;type:int;comment:压缩上限"`                        //压缩上限
	CompressionSizeLower int    `gorm:"column:compression_size_lower;type:int;comment:压缩下限"`                        //压缩下限
	CompressionUpperType int    `gorm:"column:compression_upper_type;type:int2;default:1;comment:压缩上限大小类型 1.k 2.m"` //压缩上限大小类型 1.k 2.m
	CompressionLowerType int    `gorm:"column:compression_lower_type;type:int2;default:1;comment:压缩下限大小类型 1.k 2.m"` //压缩下限大小类型 1.k 2.m
	//访问限制
	FrequencyLimit  int    `gorm:"column:frequency_limit;type:int;default:10;comment:频率限制"`   //频率限制
	MinuteSetting   int    `gorm:"column:minute_setting;type:int;default:1;comment:分钟数"`      //分钟数
	MinuteTotalTime int    `gorm:"column:minute_total_time;type:int;default:1;comment:分钟总次数"` //分钟总次数
	ConnectNum      int    `gorm:"column:connect_num;type:int;default:10;comment:连接数限制"`      //连接数限制
	LocationInfo    string `gorm:"column:location_info;type:text;comment:地区信息"`               //地区信息
	ReleaseTime     int    `gorm:"column:release_time;type:int;comment:释放时间"`                 //释放时间
	Remark          string `gorm:"column:remark;type:varchar(50);comment:备注"`                 //备注

	CreatedAt int64 `gorm:"column:created_at;type:int8;index:idx_configure_created_at;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;index:idx_configure_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *Configure) TableName() string {
	return "configure"
}

func (r *Configure) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&Configure{}); err != nil {
		panic(err)
	}
	return nil
}
