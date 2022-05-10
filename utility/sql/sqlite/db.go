package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDriver struct {
	*Config
	db *gorm.DB
}

type Config struct {
	GormConf *gorm.Config //gorm配置
	DBName   string       //数据库名称
}

func New(conf *Config) *sqliteDriver {
	return &sqliteDriver{
		Config: conf,
	}
}

func (this *sqliteDriver) Init() {

}

func (this *sqliteDriver) Connect() (err error) {
	this.db, err = gorm.Open(sqlite.Open(this.DBName), this.GormConf)
	if err != nil {
		return
	}
	return nil
}

func (this *sqliteDriver) Close() error {
	sqlDB, err := this.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (this *sqliteDriver) DB() *gorm.DB {
	return this.db
}
