package model

import "gorm.io/gorm"

const (
	LoginRecord_Display_type_all   = 1 //所有人可见
	LoginRecord_Display_type_admin = 2 //仅限管理员可见
)

type LoginRecord struct {
	Id            string `gorm:"column:id;type:varchar(50);primary_key"`                                         //id
	Uid           string `gorm:"column:uid;type:varchar(50);index:idx_login_records_uid;not null;comment:uid登录"` //uid登录ID
	Type          int64  `gorm:"column:type;type:int8;not null;comment:登录人员类型 1管理员 2用户"`                         //1管理员 2用户
	IP            string `gorm:"column:ip;type:varchar(50);not null;comment:登录IP"`                               //ip
	Region        string `gorm:"column:region;type:varchar(50);not null;comment:登录地区"`                           //地区
	DisplayType   int64  `gorm:"column:display_type;type:int2;default:1;comment:显示逻辑 1管理员用户可见 2管理员可见"`           //显示逻辑 1管理员用户可见 2管理员可见
	Device        string `gorm:"column:device;type:varchar(50);comment:设备"`                                      //设备
	Os            string `gorm:"column:os;type:varchar(50);comment:系统"`                                          //系统
	OsVersion     string `gorm:"column:os_version;type:varchar(50);comment:系统版本"`                                //系统版本
	Browser       string `gorm:"column:browser;type:varchar(50);comment:浏览器"`                                    //浏览器
	BrowserVision string `gorm:"column:browser_vision;type:varchar(50);comment:浏览器版本"`                           //浏览器版本
	CreatedAt     int64  `gorm:"column:created_at;type:int8;index:idx_login_records_created_at;not null"`        //创建时间
}

func (this *LoginRecord) TableName() string {
	return "login_records"
}

func (r *LoginRecord) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&LoginRecord{}); err != nil {
		panic(err)
	}
	return nil
}
