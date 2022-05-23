package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Id              string                `gorm:"column:id;type:varchar(50);primary_key;not null"`                              //id
	Nick            string                `gorm:"column:nick;type:varchar(50);uniqueIndex:idx_users_nick;not null;comment:昵称"`  //昵称
	Password        string                `gorm:"column:password;type:varchar(50);not null;comment:密码"`                         //密码，md5加密
	Role            int                   `gorm:"column:role;type:int;index:idx_users_role;not null;default:2;comment:角色"`      //角色，1：管理员，2：操作员
	PrimaryDomainId string                `gorm:"column:primary_domain_id;type:varchar(50);comment:主域名ID"`                      //绑定的主域名ID
	Email           string                `gorm:"column:email;type:varchar(50);not null;uniqueIndex:idx_users_nick;comment:邮箱"` //邮箱
	State           int                   `gorm:"column:state;type:int;comment:1正常 2禁用"`                                        //1:正常  2:禁用
	LastTime        int64                 `gorm:"column:last_time;comment:最后登录时间"`                                              //最新登录时间
	UserGroup       string                `gorm:"column:user_group;type:varchar(50);comment:用户分组"`                              //用户分组ID
	LastIP          string                `gorm:"column:last_ip;type:varchar(50);comment:最新登录IP"`                               //最新登录IP
	Region          string                `gorm:"column:region;type:varchar(100);comment:最新登录所属地区"`                             //最新登录所属地区
	CreatedIp       string                `gorm:"column:created_ip;type:varchar(50);comment:创建IP"`                              //创建IP
	CreatedAt       int64                 `gorm:"column:created_at;type:int8;index:idx_users_created_at;not null"`              //创建时间
	UpdatedAt       int64                 `gorm:"column:updated_at;type:int8;index:idx_users_updated_at;not null"`              //修改时间，新记录保持和创建时间一致
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;uniqueIndex:idx_users_nick"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	return nil
}
