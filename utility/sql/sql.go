package sql

//关系数据库
import (
	"errors"
	"go_web_example/utility/sql/mysql"
	"go_web_example/utility/sql/postgre"
	"go_web_example/utility/sql/sqlite"

	"gorm.io/gorm"
)

const (
	MysqlDirver   = "mysql"
	PostgreDirver = "postgre"
	SqliteDirver  = "sqlite"
)

func New(config DBConfig) *SqlDB {
	db := &SqlDB{
		DBConfig: config,
	}
	return db
}

//配置
type DBConfig struct {
	GormConf    *gorm.Config //gorm配置
	DBType      string       //数据库类型
	IP          string       //ip
	Port        string       //端口
	User        string       //用户名
	Password    string       //密码
	DBName      string       //数据库名称
	MaxIdleConn int          //最大空闲连接数
	MaxOpenConn int          //最大打开的连接数
	MaxLifetime int          //连接可重用的最大时间
}

//驱动接口
type driver interface {
	Init()
	Connect() error
	DB() *gorm.DB
	Close() error
}

//数据库操作对象
type SqlDB struct {
	DBConfig
	db *gorm.DB
}

//创建数据库连接
func (s *SqlDB) Connect() (err error) {
	var (
		entity   driver
		gormConf *gorm.Config
	)
	//gorm配置
	gormConf = &gorm.Config{}
	//创建数据库对象
	switch s.DBType {
	case MysqlDirver:
		entity = mysql.New(&mysql.Config{IP: s.IP, Port: s.Port, User: s.User, Password: s.Password, DBName: s.DBName, MaxOpenConn: s.MaxOpenConn, MaxIdleConn: s.MaxIdleConn, MaxLifetime: s.MaxLifetime, GormConf: gormConf})
	case PostgreDirver:
		entity = postgre.New(&postgre.Config{IP: s.IP, Port: s.Port, User: s.User, Password: s.Password, DBName: s.DBName, MaxOpenConn: s.MaxOpenConn, MaxIdleConn: s.MaxIdleConn, MaxLifetime: s.MaxLifetime, GormConf: gormConf})
	case SqliteDirver:
		entity = sqlite.New(&sqlite.Config{DBName: s.DBName, GormConf: gormConf})
	default:
		return errors.New("未声明或不受支持的数据库类型")
	}
	//初始化
	entity.Init()
	//创建连接
	if err = entity.Connect(); err != nil {
		return err
	}
	//获取数据库操作对象
	s.db = entity.DB()
	return nil
}

//获取数据库操作对象
func (s *SqlDB) DB() *gorm.DB {
	return s.db
}

//关闭连接
func (s *SqlDB) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
