package model

import "gorm.io/gorm"

const (
	Node_Status_close        = 2
	Node_Status_open         = 1
	Node_Online_status_close = 2
	Node_Online_status_open  = 1
)

type Node struct {
	Id   string `gorm:"column:id;type:varchar(50);primary_key;not null"`                       //id
	Name string `gorm:"column:name;type:varchar(50);index:idx_nodes_name;not null;comment:名称"` //名称
	Os   string `gorm:"column:os;type:varchar(50);comment:操作系统"`                               //操作系统

	Status       int    `gorm:"column:status;type:int2;default:2;not null;comment:状态1开启2关闭"`            //状态1开启2关闭
	OnlineStatus int    `gorm:"column:online_status;type:int2;not null;default:2;comment:在线状态 1在线 2离线"` //在线状态 1在线 2离线
	AtsStatus    int    `gorm:"column:ats_status;type:int8;default:2;comment:宕机状态 1宕机 2未宕机"`            //宕机状态 1宕机 2未宕机
	Version      string `gorm:"column:version;type:varchar(50);comment:版本号"`                            //版本号
	NewVersion   string `gorm:"column:new_version;type:varchar(50);comment:最新版本"`                       //最新版本
	MaxUpload    int    `gorm:"column:max_upload;type:int;default:0;comment:上传峰值"`                      //上传峰值
	MaxDownLoad  int    `gorm:"column:max_download;type:int;default:0;comment:下载峰值"`                    //下载峰值
	MemoryUsed   int    `gorm:"column:memory_used;type:int8;comment:内存使用率"`                             //内存使用率
	CpuUsed      int    `gorm:"column:cpu_used;type:int8;comment:Cpu使用率"`                               //Cpu使用率
	DiskUsed     int    `gorm:"column:disk_used;type:int8;comment:硬盘使用率"`                               //硬盘使用率

	GroupName string `gorm:"-"`
	CnameAddr string `gorm:"column:cname_addr;type:varchar(50);comment:自动分配"` //cname地址，自动分配

	//IP部分配置
	IpV6       string `gorm:"column:ipv6;type:varchar(50);comment:ipv6地址"`             //ipv6地址
	SynTiming  int    `gorm:"column:syn_timing;type:int;comment:同步定时时间(秒)如果为0,说明手动同步"` //同步定时时间(秒)如果为0,说明手动同步
	SyncConfAt int64  `gorm:"column:sync_conf_at;type:int8;comment:最后同步时间"`            //最后同步时间

	//设置内容
	Pid       int `gorm:"column:pid;type:int;default:1;comment:如果使用模板,则填充此字段 1使用默认模板 2使用自定义数据"` //如果使用模板,则填充此字段 1使用默认模板 2使用自定义数据
	CacheType int `gorm:"column:cache_type;type:int;default:1;comment:缓存配置-缓存类型"`               //缓存配置-缓存类型
	//内存缓存占比
	MemCacheProportion int `gorm:"column:mem_cache_proportion;type:int8;comment:内存站比(1-100 百分比)"` //内存站比(1-100 百分比)
	//磁盘缓存设置
	Path string `gorm:"column:path;type:varchar(128);comment:磁盘路径"` //磁盘路径
	//分区
	PartitionPath string `gorm:"column:partition_path;type:varchar(50);comment:分区路径"` //分区路径
	//磁盘缓存占比(如果设置了分区路径默认为100)
	DiskCacheProportion int `gorm:"column:disk_cache_proportion;type:int8;comment:磁盘缓存占比"` //磁盘缓存占比

	UseRegionConfig    int    `gorm:"column:use_region_config;type:int2;default:2;comment:是否使用区域配置项 1开启 2关闭"` //是否使用区域配置项 1开启 2关闭
	RegionConfigTerm   string `gorm:"column:region_config_term;type:varchar(50);comment:制定设置项"`               //制定设置项
	RegionConfigBindId string `gorm:"column:region_config_bind_id;type:varchar(50);comment:使用区域默认设置的区域对象"`    //使用区域默认设置的区域对象

	LastCleanCacheAt int   `gorm:"column:last_clean_cache_at;type:int8;default:0;comment:最后清除缓存时间"`                           //最后清除缓存时间
	CreatedAt        int64 `gorm:"column:created_at;type:int8;index:idx_nodes_created_at;not null;comment:创建时间"`              //创建时间
	UpdatedAt        int64 `gorm:"column:updated_at;type:int8;index:idx_nodes_updated_at;not null;comment:修改时间，新记录保持和创建时间一致"` //修改时间，新记录保持和创建时间一致
	//附加参数
	//RegionIds       []string `gorm:"-;comment:"` //返回用绑定对区域ID
	//RegionNames     []string `gorm:"-;comment:"` //返回绑定区域的名称
	//Personalization []int    `gorm:"-;comment:"` // 绑定配置项 1.压缩管理 2.缓存配置 3.访问限制 4.自定义错误页面 5.黑名单 6.白名单 7.其他设置
	//WithinIpBack    []string `gorm:"-;comment:"` //内网IP
	//OutSideIpBack   []string `gorm:"-;comment:"` //外网IP
	//OutSideIp       string   `gorm:"-;comment:"` //外网主IP
}

func (this *Node) TableName() string {
	return "nodes"
}

func (r *Node) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&Node{}); err != nil {
		panic(err)
	}
	return nil
}
