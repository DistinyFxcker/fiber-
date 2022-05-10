package postgre

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreDriver struct {
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

func New(conf *Config) *postgreDriver {
	return &postgreDriver{
		Config: conf,
	}
}

func (this *postgreDriver) Init() {

}

func (this *postgreDriver) Connect() (err error) {
	dsn := "host=" +
		this.IP + "  user=" +
		this.User + " dbname=" +
		this.DBName + " sslmode=disable  password=" +
		this.Password
	this.db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), this.GormConf)
	if err != nil {
		return err
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

func (this *postgreDriver) Close() error {
	sqlDB, err := this.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (this *postgreDriver) DB() *gorm.DB {
	return this.db
}
