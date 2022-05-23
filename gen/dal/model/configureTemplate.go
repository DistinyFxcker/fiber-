package model

import "gorm.io/gorm"

type ConfigureTemplate struct {
	Id           string `gorm:"column:id;type:varchar(50);not null;primary_key;not null"`
	Name         string `gorm:"column:name;type:varchar(50);comment:模板名称"`                   //模板名称
	RelationName string `gorm:"column:relation_name;type:varchar(50);comment:关联对象名称"`        //关联对象名称
	RelationId   string `gorm:"column:relation_id;type:varchar(50);comment:关联ID"`            //关联ID(节点ID 或者 区域ID)
	RelationType int    `gorm:"column:relation_type;type:int;comment:关联节点类型(1节点 2区域ID 3域名)"` //关联节点类型(1节点 2区域ID 3域名)
	CreatedId    string `gorm:"column:created_id;type:varchar(50);not null;comment:创建ID"`    //创建ID
	CreatedRole  int    `gorm:"column:created_role;type:int;comment:创建人角色"`                  //创建人角色
	Type         int    `gorm:"column:type;type:int;default:1;comment:模板类型 1用户专属 2公共模板"`     //模板类型 1用户专属 2公共模板
	//公共配置
	//超时配置
	ConnectTime int `gorm:"column:connect_time;type:int;default:30;comment:连接超时时间(单位:秒)"` //连接超时时间(单位:秒)
	//缓存规则
	MinConnNum     int `gorm:"column:min_conn_num;type:int;default:1;comment:最小请求连接数"`              //最小请求连接数
	MaxMemoryCache int `gorm:"column:max_memory_cache;type:int;comment:最大内存缓存(单位M)"`                //最大内存缓存(单位M)
	MaxDiskCache   int `gorm:"column:max_disk_cache;type:int;comment:磁盘缓存大小(单位M)"`                  //磁盘缓存大小(单位M)
	DiskCacheType  int `gorm:"column:disk_cache_type;type:int;default:1;comment:最大磁盘缓存类型(1普通 2智能)"` //最大磁盘缓存类型(1普通 2智能)
	//压缩配置
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
	Remark          string `gorm:"column:remark;type:varchar(50);comment:备注"`                 //备注

	//*******域名配置
	BusinessType string `gorm:"column:business_type;type:varchar(50);comment:业务类型 1网站加速 2文件下载 3视频点播 4全站加速"` //业务类型 1网站加速 2文件下载 3视频点播 4全站加速
	//源站信息
	DomainSourceSortType int    `gorm:"column:domain_sourceSortType;type:int;comment:源站主备 权重设置 1主备 2权重"`          //源站主备 权重设置 1主备 2权重
	DomainSourceType     int    `gorm:"column:domain_source_type;type:int;default:1;comment:源站协议类型 1http 2https"` //源站协议类型 1http 2https
	DomainSourceList     string `gorm:"column:domain_source_list;type:text;comment:源站列表"`                         //源站列表
	//回源配置
	DomainForward       int    `gorm:"column:domain_forward;type:int;default:2;comment:回源配置开关"`     //回源配置开关
	DomainForwardHost   string `gorm:"column:domain_forward_host;type:varchar(128);comment:回源Host"` //回源Host
	DomainForwardHeader string `gorm:"column:domain_forward_header;type:text;comment:回源请求头"`        //回源请求头
	//缓存配置
	CacheConfig string `gorm:"column:cache_config;type:text;comment:缓存配置"` //缓存配置
	//响应头配置
	CacheSource int    `gorm:"column:cache_source;type:int;default:1;comment:缓存遵循源站"` //缓存遵循源站
	HeaderList  string `gorm:"column:header_list;type:text;comment:HTTP"`             //HTTP
	//自定义错误页面
	ErrorPageList string `gorm:"column:error_page_list;type:text;comment:自定义错误页面"` //自定义错误页面
	//自定义站点协议
	DomainAgreementType int    `gorm:"column:domain_agreement_type;type:int;not null;default:1;comment:域名网站协议类型 1仅HTTP 2强制HTTPS 3Http+https 4自定义协议端口"` //域名网站协议类型 1仅HTTP 2强制HTTPS 3Http+https 4自定义协议端口
	Agreement           string `gorm:"column:agreement;type:text;comment:自定义站点协议"`                                                                     //自定义站点协议
	//访问控制
	EmptyReferer    int    `gorm:"column:empty_referer;type:int2;default:1;comment:1合法 2不合法"` // 1合法 2不合法
	BlackListDomain string `gorm:"column:black_list_domain;type:text;comment:黑名单域名"`          //黑名单域名
	WhiteListDomain string `gorm:"column:white_list_domain;type:text;comment:白名单域名"`          //白名单域名
	//访问控制-开关
	AccessControl        int `gorm:"column:access_control;type:int2;default:2;comment:访问控制开关"` //访问控制开关
	AntiTheftChainSwitch int `gorm:"column:anti_theft_chain_switch;type:int2;comment:防盗链开关"`   //防盗链开关
	BlackListSwitch      int `gorm:"column:black_list_switch;type:int2;comment:黑名单开关"`         //黑名单开关
	WhiteListSwitch      int `gorm:"column:white_list_switch;type:int2;comment:白名单开关"`         //白名单开关
	//访问控制-黑白名单
	BlackList string `gorm:"column:black_list;type:text;comment:黑名单"` //黑名单
	WhiteList string `gorm:"column:white_list;type:text;comment:白名单"` //白名单
	//域名转发
	DomainForwardList string `gorm:"column:domain_forward_list;type:text;comment:域名转发列表"` //域名转发列表
	//其他设置
	IgnoreURL          int `gorm:"column:ignore_url;type:int;default:1;comment:是否忽略Url参数"`          //是否忽略Url参数
	GzipCompress       int `gorm:"column:gzip_compress;type:int;default:6;comment:Gzip压缩级别"`        //Gzip压缩级别
	GzipCompressSwitch int `gorm:"column:gzip_compress_switch;type:int;default:2;comment:Gzip压缩开关"` //Gzip压缩开关

	CreatedAt int64 `gorm:"column:created_at;type:int8;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *ConfigureTemplate) NewTable() string {
	return "configure_template"
}

func (r *ConfigureTemplate) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&ConfigureTemplate{}); err != nil {
		panic(err)
	}
	return nil
}
