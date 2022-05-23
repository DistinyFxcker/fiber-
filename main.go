package main

import (
	"github.com/spf13/viper"
	service "go_web_example/app"
	"go_web_example/cmd"
	dbCmd "go_web_example/gen/cmd"
	"go_web_example/infrastructure/valid"
	"go_web_example/utility/sql"
)

//var RedisConnect
func init() {
	dbCmd.InitGen()
	SetValidator()
}
func main() {
	var (
		_svm       service.Service
		err        error
		sqlConnect *sql.SqlDB
	)
	//连接数据库并且迁移
	sqlConnect = sqlInit()
	//根据数据库创建CURL模型
	err = dbCmd.InitMigrate(sqlConnect.DB())
	if err != nil {
		panic(err)
	}
	{
		_svm = cmd.New(sqlConnect.DB())
		if err = _svm.Init(); err != nil {
			panic(err)
		}
		if err = _svm.Start(); err != nil {
			panic(err)
		}
	}
}

//数据库初始化
func sqlInit() *sql.SqlDB {
	viper := viper.New()
	viper.SetConfigName("sql")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./gen/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
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

//初始化全局验证
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
