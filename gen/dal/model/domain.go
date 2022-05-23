package model

import "gorm.io/gorm"

const (
	Domain_Certification_normal    = 1 //可用
	Domain_Certification_beOverdue = 2 //过期
	Domain_Certificate_Sonn        = 3 //即将过期
	Domain_Certificate_Applying    = 4 //申请中
	Domain_Certificate_Error       = 5 //错误
	Domain_Certificate_Ns_Error    = 6 //ns错误

	Domain_Certificate_Auto   = 2 //自动申请
	Domain_Certificate_normal = 1 //手动填写

	Domain_TypeSimple               = 1 //二级域名
	Domain_TypePanDomainName        = 2 //泛域名
	Domain_Sub                      = 3 //子域名
	Domain_AgreementType_http       = 1 //http
	Domain_AgreementType_https      = 2 //https
	Domain_AgreementType_http_https = 3 //http+https
	Domain_AgreementType_custom     = 4 //自定义端口

	Domain_Status_open    = 1 //启动
	Domain_Status_close   = 2 //未启动
	Domain_Status_arrears = 3 //欠费停机
)

type Domain struct {
	Id                   string `gorm:"column:id;type:varchar(50);primary_key;not null"`                                                 //id
	Name                 string `gorm:"column:name;type:varchar(50);not null;comment:域名名称"`                                              //域名名称
	Status               int    `gorm:"column:status;type:int2;not null;default:1;comment:启用状态1(触发自动同步)启动 2未启动 3 欠费停机状态"`                //启用状态1(触发自动同步)启动 2未启动 3 欠费停机状态
	NsState              int    `gorm:"column:ns_state;type:int2;default:2;comment:ns状态"`                                                //ns状态
	Domain               string `gorm:"column:domain;type:text;not null;comment:域名"`                                                     //域名
	MasterDomain         string `gorm:"column:master_domain;type:text;comment:主域名"`                                                      //主域名
	DomainType           int    `gorm:"domain_type;type:int;not null;comment:域名类型 1普通域名(二级域名) 2泛域名 3子域名"`                                //域名类型 1普通域名(二级域名) 2泛域名 3子域名
	DnsState             int    `gorm:"column:dns_state;type:int2;default:1;comment:DNS同步状态;comment:DNS的同步状态,1成功 2失败 如果失败了,用定时任务定时重新执行"` //DNS的同步状态,1成功 2失败 如果失败了,用定时任务定时重新执行
	Region               string `gorm:"column:region;type:varchar(50);not null;index:idx_domain_region;comment:所属区域"`                    //所属区域
	BusinessType         string `gorm:"column:business_type;type:varchar(50);index:idx_domain_business_type;comment:业务类型"`               //业务类型
	DomainSourceSortType int    `gorm:"column:domain_source_sort_type;type:int;default:1;comment:源站主备 权重设置 1主备 2权重"`                     //源站主备 权重设置 1主备 2权重
	DomainSourceType     int    `gorm:"column:domain_source_type;type:int;default:1;comment:源站协议类型 1http 2https"`                        //源站协议类型 1http 2https
	BillingMethod        int    `gorm:"column:billing_method;type:int;default:1;comment:计费方式 1流量计费 2带宽计费"`                               //计费方式 1流量计费 2带宽计费
	Cname                string `gorm:"column:cname;type:varchar(128);comment:配置生成的Cname"`                                               //配置生成的Cname
	PrimaryDomainId      string `gorm:"column:primary_domain_id;type:varchar(50);comment:绑定的主域名ID"`                                      //绑定的主域名ID

	//switch
	DomainForward      int `gorm:"column:domain_forward;type:int;default:2;comment:回源配置开关"`           //回源配置开关
	Http2              int `gorm:"column:http2;type:int;default:1;default:2;comment:http2开关 1开启 2关闭"` //http2开关 1开启 2关闭
	GzipCompressSwitch int `gorm:"column:gzip_compress_switch;type:int;default:2;comment:Gzip压缩开关"`   //Gzip压缩开关
	//新增开关
	AccessControl        int `gorm:"column:access_control;type:int2;default:2;comment:访问控制开关"` //访问控制开关
	AntiTheftChainSwitch int `gorm:"column:anti_theft_chain_switch;type:int2;comment:防盗链开关"`   //防盗链开关
	BlackListSwitch      int `gorm:"column:black_list_switch;type:int2;comment:黑名单开关"`         //黑名单开关
	WhiteListSwitch      int `gorm:"column:white_list_switch;type:int2;comment:白名单开关"`         //白名单开关

	IgnoreURL    int `gorm:"column:ignore_url;type:int;default:1;comment:是否忽略Url参数"`   //是否忽略Url参数
	CacheSource  int `gorm:"column:cache_source;type:int;default:1;comment:缓存遵循源站"`    //缓存遵循源站
	GzipCompress int `gorm:"column:gzip_compress;type:int;default:6;comment:Gzip压缩级别"` //Gzip压缩级别
	//status
	DomainAgreementType int `gorm:"column:domain_agreement_type;type:int;not null;default:1;comment:域名网站协议类型 1仅HTTP 2强制HTTPS 3Http+https 4自定义协议端口"` //域名网站协议类型 1仅HTTP 2强制HTTPS 3Http+https 4自定义协议端口

	//手动证书协议:证书,状态,端口,是否过期
	Certificate string `gorm:"column:certificate;type:text;comment:证书"` //证书
	Key         string `gorm:"column:key;type:text;comment:证书密钥"`       //证书密钥

	//自动证书协议:
	CertificateAuto  string `gorm:"column:certificate_auto;type:text;comment:自动创建证书"`                    //自动创建证书
	KeyAuto          string `gorm:"column:key_auto;type:text;comment:自动创建证书秘钥"`                          //自动创建证书秘钥
	CertificateBrand string `gorm:"column:certificate_brand;type:varchar(50);comment:证书品牌;comment:证书品牌"` //证书品牌
	Arithmetic       string `gorm:"column:arithmetic;type:varchar(50);comment:证书算法;comment:证书算法"`        //证书算法
	Validate         string `gorm:"column:validate;type:varchar(50);comment:证书验证方法;comment:证书验证方法"`      //证书验证方法

	//证书配置
	CertificateType int64 `gorm:"column:certificate_type;type:int2;default:0;comment:证书类型;comment:证书类型 1手动添加 2自动生成"` //证书类型 1手动添加 2自动生成
	HSCStatus       int   `gorm:"column:hsc_status;type:int2;default:0;comment:手动证书状态0未添加1正常2过期"`                    //手动证书状态0未添加1正常2过期
	OverTime        int64 `gorm:"column:over_time;type:int8;comment:证书过期时间"`                                         //证书过期时间
	OverTimeAuto    int64 `gorm:"column:over_time_auto;type:int8;comment:自动证书过期时间"`                                  //自动证书过期时间

	//防盗链
	EmptyReferer    int    `gorm:"column:empty_referer;type:int2;default:1;comment:1合法 2不合法"` // 1合法 2不合法
	BlackListDomain string `gorm:"column:black_list_domain;type:text;comment:黑名单域名"`          //黑名单域名
	WhiteListDomain string `gorm:"column:white_list_domain;type:text;comment:白名单域名"`          //白名单域名

	EstablishId   string `gorm:"column:establish_id;type:varchar(50);not null;comment:创建ID"`              //创建ID
	EstablishType int    `gorm:"column:establish_type;type:varchar(50);not null;default:1;comment:创建人类型"` //创建人类型
	EstablishName string `gorm:"column:establish_name;type:varchar(50);not null;comment:创建人姓名"`           //创建人姓名
	//其他设置
	CreatedAt int64 `gorm:"column:created_at;type:int8;index:idx_domain_created_at;not null"` //创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int8;index:idx_domain_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *Domain) TableName() string {
	return "domain"
}

func (r *Domain) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&Domain{}); err != nil {
		panic(err)
	}
	return nil
}
