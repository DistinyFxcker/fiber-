package main

import (
	service "cdnFiber/app"
	"cdnFiber/cmd"
	"cdnFiber/infrastructure/valid"
	"cdnFiber/utility/sql"
	"github.com/spf13/viper"
)

var (
	SqlConnect *sql.SqlDB
)

//var RedisConnect

func main() {
	var (
		_svm service.Service
		err  error
	)
	SetConfig()
	SetLog()
	SetValidator()
	//SqlConnect = sqlInit()
	//gen.New(SqlConnect.DB())
	{
		_svm = cmd.New()
		_svm.Init()
		if err = _svm.Start(); err != nil {
			panic(err)
		}
	}
}

func SetConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

//数据库初始化
func sqlInit() *sql.SqlDB {
	sqlDB := sql.New(sql.DBConfig{
		DBType:      viper.GetString("sql.DBType"),
		User:        viper.GetString("sql.User"),
		Password:    viper.GetString("sql.Password"),
		IP:          viper.GetString("sql.IP"),
		Port:        viper.GetString("sql.Port"),
		DBName:      viper.GetString("sql.DBName"),
		MaxIdleConn: viper.GetInt("sql.MaxIdleConn"),
		MaxOpenConn: viper.GetInt("sql.MaxOpenConn"),
		MaxLifetime: viper.GetInt("sql.MaxLifetime"),
	})
	if err := sqlDB.Connect(); err != nil {
		panic(err.Error())
	}
	return sqlDB
}

func SetLog() {
	//zap.New()
}

func SetValidator() {
	valid.Init()
}

//redis初始化
//func redisInit() redis.RedisClient {
//	redisClient := redis.New(redis.Config{
//		Host:            viper.GetString("redis.Host"),
//		Port:            viper.GetString("redis.Port"),
//		DB:              viper.GetInt("redis.DB"),
//		MaxIdle:         viper.GetInt("redis.MaxIdle"),
//		MaxActive:       viper.GetInt("redis.MaxActive"),
//		IdleTimeoutSecs: viper.GetInt("redis.IdleTimeoutSecs"),
//		Password:        viper.GetString("redis.Password"),
//	})
//	return redisClient
//}
