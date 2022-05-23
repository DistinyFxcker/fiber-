package model

import "gorm.io/gorm"

//邮件发送
type EmailRecord struct {
	Id         string `gorm:"column:id;type:varchar(50);primary_key;AUTO_INCREMENT"`                                //id
	Uid        string `gorm:"column:uid;type:varchar(50);index:idx_email_records_uid;not null;comment:用户ID"`        //uid
	Type       int    `gorm:"column:type;type:int8;not null;comment:邮箱发送类型 1管理员 2用户"`                               //邮箱发送类型 1管理员 2用户
	OperatorId string `gorm:"column:operatorId;type:varchar(50);not null;comment:操作员ID"`                            //操作员id
	Mail       string `gorm:"column:ip;type:varchar(50);not null;comment:邮箱" `                                      //邮箱
	CreatedAt  int64  `gorm:"column:created_at;type:int8;index:idx_email_records_created_at;not null;comment:创建时间"` //创建时间
}

func (this *EmailRecord) TableName() string {
	return "email_records"
}

func (r *EmailRecord) Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&EmailRecord{}); err != nil {
		panic(err)
	}
	return nil
}
