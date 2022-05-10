package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type mysqlDriver struct {
	*Config
	db *gorm.DB
}

type Config struct {
	GormConf    *gorm.Config //gorm配置
	IP          string       //ip
	Port        string       //端口
	User        string       //用户名
	Password    string       //密码
	DBName      string       //数据库名称
	MaxIdleConn int          //最大空闲连接数
	MaxOpenConn int          //最大打开的连接数
	MaxLifetime int          //连接可重用的最大时间
}

func New(conf *Config) *mysqlDriver {
	return &mysqlDriver{
		Config: conf,
	}
}
func (this *mysqlDriver) Init() {

}

//创建连接
func (this *mysqlDriver) Connect() (err error) {
	dsn := this.User + ":" +
		this.Password + "@tcp(" +
		this.IP + ":" + this.Port + ")/" +
		this.DBName + "?charset=utf8"

	this.db, err = gorm.Open(mysql.Open(dsn), this.GormConf)
	if err != nil {
		return
	}
	sqlDB, err := this.db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(this.MaxIdleConn)
	sqlDB.SetMaxOpenConns(this.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(this.MaxLifetime) * time.Second)
	return nil
}

//关闭连接
func (this *mysqlDriver) Close() error {
	sqlDB, err := this.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

//获取db对象
func (this *mysqlDriver) DB() *gorm.DB {
	return this.db
}
