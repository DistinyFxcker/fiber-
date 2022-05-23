package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

const (
	Admin_Default_Password = "123.abc"
	Admin_Default_Email    = "18581560773@163.com"
)

type Admin struct {
	Id       string `gorm:"column:id;type:varchar(50);primary_key;not null"`                              //id
	Nick     string `gorm:"column:nick;type:varchar(50);uniqueIndex:idx_admins_nick;not null;comment:昵称"` //昵称
	Password string `gorm:"column:password;type:varchar(50);not null;comment:密码"`                         //密码，md5加密
	//Role      int    `gorm:"column:role;type:int;index:idx_users_role;not null;default:2;comment:" `    //角色，1：管理员，2：操作员
	Email     string                `gorm:"column:email;type:varchar(50);uniqueIndex:idx_admins_nick;not null;comment:邮箱"` //邮箱
	State     int                   `gorm:"column:state;type:int;comment:状态 1正常 2禁止"`                                      //1:正常  2:禁用
	LastTime  int64                 `gorm:"column:last_time;最新登录时间"`                                                       //最新登录时间
	LastIP    string                `gorm:"column:last_ip;type:varchar(50);comment:最新登录IP"`                                //最新登录IP
	Region    string                `gorm:"column:region;type:varchar(100);comment:最新登录所属地区"`                              //最新登录所属地区
	CreatedIp string                `gorm:"column:created_ip;type:varchar(50);comment:创建IP"`                               //创建IP
	CreatedAt int64                 `gorm:"column:created_at;type:int8;index:idx_users_created_at;not null" `              //创建时间
	UpdatedAt int64                 `gorm:"column:updated_at;type:int8;index:idx_users_updated_at;not null" `              //修改时间，新记录保持和创建时间一致
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;uniqueIndex:idx_admins_nick"`
}

func (m *Admin) tableName() string {
	return "admins"
}

func (m *Admin) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&Admin{}); err != nil {
		panic(err)
	}
	return nil
}
